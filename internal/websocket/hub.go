package websocket

import (
	"sync"
)

// Client представляет WebSocket клиента
type Client struct {
	Hub  *Hub
	Conn *Conn
	Send chan []byte
	ID   string
}

// Hub управляет WebSocket соединениями
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

// NewHub создает новый Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run запускает Hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// Broadcast отправляет сообщение всем клиентам
func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}

// GetClients возвращает количество подключенных клиентов
func (h *Hub) GetClients() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}
