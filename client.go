package liquidtap

import (
	"github.com/okb48/pusher-websocket-go"
)

type liquidtap struct {

	c *pusher.Client
}

const (
	pusherKey = "LIQUID"
	pusherTapHost = "tap.liquid.com"
)

func NewClient() *liquidtap  {

	c := pusher.New(pusherKey,pusherTapHost)
	//c.Debug = true

	return &liquidtap{c:c}
}