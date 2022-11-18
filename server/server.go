package server

var HttpServer *Http
var WebsocketServer *Websocket

func Init(httpConfig *HttpConfig) {

	HttpServer = &Http{
		Conf: httpConfig,
	}
	HttpServer.Setup()

	WebsocketServer = &Websocket{
		App:  HttpServer.App,
		Conf: &WebsocketConfig{},
	}
	WebsocketServer.Setup()
}
