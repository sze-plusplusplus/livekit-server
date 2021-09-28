package service

import (
	"context"
	"errors"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/twitchtv/twirp"
	"google.golang.org/grpc"

	"github.com/livekit/protocol/auth"
)

type grantKeyType string

const (
	authorizationHeader              = "Authorization"
	bearerPrefix                     = "Bearer "
	grantsKey           grantKeyType = "grants"
	accessTokenParam                 = "access_token"
)

var (
	ErrPermissionDenied = errors.New("permissions denied")
)

// authentication middleware
type APIKeyAuthMiddleware struct {
	provider auth.KeyProvider
}

func NewAPIKeyAuthMiddleware(provider auth.KeyProvider) *APIKeyAuthMiddleware {
	return &APIKeyAuthMiddleware{
		provider: provider,
	}
}

func (m *APIKeyAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL != nil && r.URL.Path == "/rtc/validate" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	grants, err := checkToken(m.provider, r.Header.Get(authorizationHeader), r.FormValue(accessTokenParam))
	if err != nil {
		handleError(w, http.StatusUnauthorized, err.Error())
		return
	}

	r = r.WithContext(context.WithValue(r.Context(), grantsKey, grants))

	next.ServeHTTP(w, r)
}

func GetGrants(ctx context.Context) *auth.ClaimGrants {
	claims, ok := ctx.Value(grantsKey).(*auth.ClaimGrants)
	if !ok {
		return nil
	}
	return claims
}

func SetAuthorizationToken(r *http.Request, token string) {
	r.Header.Set(authorizationHeader, bearerPrefix+token)
}

func EnsureJoinPermission(ctx context.Context) (name string, err error) {
	claims := GetGrants(ctx)
	if claims == nil || claims.Video == nil {
		err = ErrPermissionDenied
		return
	}

	if claims.Video.RoomJoin {
		name = claims.Video.Room
	} else {
		err = ErrPermissionDenied
	}
	return
}

func EnsureAdminPermission(ctx context.Context, room string) error {
	claims := GetGrants(ctx)
	if claims == nil || claims.Video == nil {
		return ErrPermissionDenied
	}

	if !claims.Video.RoomAdmin || room != claims.Video.Room {
		return ErrPermissionDenied
	}

	return nil
}

func EnsureCreatePermission(ctx context.Context) error {
	claims := GetGrants(ctx)
	if claims == nil {
		return ErrPermissionDenied
	}

	if claims.Video.RoomCreate {
		return nil
	}
	return ErrPermissionDenied
}

func EnsureListPermission(ctx context.Context) error {
	claims := GetGrants(ctx)
	if claims == nil {
		return ErrPermissionDenied
	}

	if claims.Video.RoomList {
		return nil
	}
	return ErrPermissionDenied
}

func EnsureRecordPermission(ctx context.Context) error {
	claims := GetGrants(ctx)
	if claims == nil || !claims.Video.RoomRecord {
		return ErrPermissionDenied
	}
	return nil
}

// wraps authentication errors around Twirp
func twirpAuthError(err error) error {
	return twirp.NewError(twirp.Unauthenticated, err.Error())
}

type grpcAuth struct {
	provider auth.KeyProvider
}

func NewGrpcAuth(p auth.KeyProvider) grpcAuth {
	return grpcAuth{
		provider: p,
	}
}

func (g *grpcAuth) authGrpc(ctx context.Context) (context.Context, error) {

	grants, err := checkToken(g.provider, metautils.ExtractIncoming(ctx).Get(authorizationHeader), "")
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, grantsKey, grants), nil
}

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request auth.
func (g *grpcAuth) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := g.authGrpc(ctx)
		if err != nil {
			return nil, err
		}

		return handler(newCtx, req)
	}
}

// StreamServerInterceptor returns a new unary server interceptors that performs per-request auth.
func (g *grpcAuth) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newCtx, err := g.authGrpc(stream.Context())
		if err != nil {
			return err
		}

		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}

func checkToken(provider auth.KeyProvider, header, form string) (*auth.ClaimGrants, error) {
	if form == "" && (header == "" || !strings.HasPrefix(header, bearerPrefix) || len(bearerPrefix) >= len(header)) {
		return nil, errors.New("invalid authorization header")
	}

	authToken := form
	if header != "" {
		authToken = header[len(bearerPrefix):]
	}

	v, err := auth.ParseAPIToken(authToken)
	if err != nil {
		return nil, errors.New("invalid authorization token")
	}

	secret := provider.GetSecret(v.APIKey())
	if secret == "" {
		return nil, errors.New("invalid API key")
	}

	grants, err := v.Verify(secret)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return grants, nil
}
