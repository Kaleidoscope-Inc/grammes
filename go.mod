module github.com/Kaleidoscope-Inc/grammes

go 1.13

require (
	github.com/google/uuid v1.1.0
	github.com/gorilla/websocket v1.4.0
	github.com/northwesternmutual/grammes v1.2.0
	github.com/smartystreets/goconvey v0.0.0-20190306220146-200a235640ff
	go.uber.org/zap v1.9.1
)

replace github.com/northwesternmutual/grammes v1.2.0 => github.com/Kaleidoscope-Inc/grammes v1.2.1-0.20230122141751-f792e4382f6f
