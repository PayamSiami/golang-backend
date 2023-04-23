
import (
	"fmt"
	"log"
	"sync"

	"github.com/gorgilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read(){
	defer func() {
	c.Pool.Unregister() <- c
	c.Conn.Close()
	}()
}