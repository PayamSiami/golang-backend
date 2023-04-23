package main 

import (
	"net/http"
	"fmt"
	"websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println("Serving")

	conn, err := websocket.Upgrade(w,r)

	if err != nil {
		fmt.Println("w," "%+\n",err)
	}

	client := &websocket.Client{
		Conn: conn
		Pool, pool,
	}
}

func setupRoutes () {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Starting webserver...")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}