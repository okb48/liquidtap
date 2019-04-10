package liquidtap

import (
	"log"
	"testing"
)

//https://developers.quoine.com/#liquid-tap

func TestClient(t *testing.T) {

	client := NewClient()

	pusherChannel := "executions_cash_btcjpy"
	pusherEvent := "created"

	channel := client.c.Subscribe(pusherChannel)

	channel.Bind(pusherEvent, func(data interface{}){
		log.Printf("%v",data)
	})

	select {}
}
