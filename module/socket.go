package module

import (
    "net/http"
    "log"
	socketio "github.com/googollee/go-socket.io"
)

func Socket() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}	
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("test", func(msg string) {
			log.Println("emit:", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	http.Handle("/socket.io/", server)
	http.ListenAndServe(":5000", nil)
	log.Println("Serving at localhost:5000...")
}