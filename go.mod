module github.com/livekit/livekit-server

go 1.17

require (
	github.com/bep/debounce v1.2.0
	github.com/c9s/goprocinfo v0.0.0-20210130143923-c95fcf8c64a8
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/elliotchance/orderedmap v1.4.0
	github.com/gammazero/deque v0.1.0
	github.com/gammazero/workerpool v1.1.2
	github.com/go-logr/logr v1.2.0
	github.com/go-logr/zapr v1.1.0
	github.com/go-redis/redis/v8 v8.11.3
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/wire v0.5.0
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jxskiss/base62 v0.0.0-20191017122030-4f11678b909b
	github.com/livekit/protocol v0.10.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/magefile/mage v1.11.0
	github.com/maxbrunsfeld/counterfeiter/v6 v6.4.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pion/ice/v2 v2.1.14
	github.com/pion/interceptor v0.1.0
	github.com/pion/logging v0.2.2
	github.com/pion/rtcp v1.2.9
	github.com/pion/rtp v1.7.4
	github.com/pion/sdp/v3 v3.0.4
	github.com/pion/stun v0.3.5
	github.com/pion/transport v0.12.3
	github.com/pion/turn/v2 v2.0.5
	github.com/pion/webrtc/v3 v3.1.10
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.26.0
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.9.1
	github.com/twitchtv/twirp v8.1.0+incompatible
	github.com/urfave/cli/v2 v2.3.0
	github.com/urfave/negroni v1.0.0
	go.uber.org/zap v1.19.1
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/tools v0.1.6 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/channels v1.1.0 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/gammazero/deque v0.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/lucsky/cuid v1.2.1 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/pion/datachannel v1.4.21 // indirect
	github.com/pion/dtls/v2 v2.0.9 // indirect
	github.com/pion/mdns v0.0.5 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/sctp v1.7.12 // indirect
	github.com/pion/srtp/v2 v2.0.5 // indirect
	github.com/pion/udp v0.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/zerolog v1.25.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20210928044308-7d9f5e0b762b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20210928142010-c7af6a1a74c9 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

// Use local, grpc generated protocol
replace github.com/livekit/protocol => github.com/sze-plusplusplus/protocol v0.9.4-grpc
