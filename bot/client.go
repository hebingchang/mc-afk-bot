package bot

import (
	"github.com/Tnze/go-mc/net"
	"github.com/google/uuid"
	"mc-afk-bot/yggdrasil"
)

// Client is used to access Minecraft server
type Client struct {
	Conn *net.Conn
	Auth Auth

	Name string
	UUID uuid.UUID

	Events      Events
	LoginPlugin map[string]func(data []byte) ([]byte, error)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient(access *yggdrasil.Access) *Client {
	id, name := access.SelectedProfile()
	return &Client{
		Auth: Auth{
			Name: name,
			UUID: id,
			AsTk: access.AccessToken(),
		},
		Events: Events{handlers: make(map[int32]*handlerHeap)},
	}
}

//Position is a 3D vector.
type Position struct {
	X, Y, Z int
}
