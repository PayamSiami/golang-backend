package websocket

type Pool struct{
	Register chan *Client
	Unregister chan *Client
	Clients map[*Client]bool
	Broadcasts chan Message
}

func NewPool() *Pool{
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcasts: make(map Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("size of connection pool:", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type :1, Body:"New User Joined..."})
			}
			break;
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("size of connection pool:", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type :1, Body:"User Disconnected..."})
			}
			break;
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients")
			for client := range pool.Clients{
				if err := client.Conn.WriteJSON(Message); 
					err != nil {			
						fmt.Println(err)
						return
				}
			break;
	}
}