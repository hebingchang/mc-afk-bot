package main

import (
	"fmt"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
	"mc-afk-bot/bot"
	"mc-afk-bot/bot/basic"
	"mc-afk-bot/yggdrasil"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func main() {
	yggdrasil.AuthURL = viper.GetString("yggdrasil.endpoint")
	resp, err := yggdrasil.Authenticate(viper.GetString("yggdrasil.email"), viper.GetString("yggdrasil.password"))
	if err != nil {
		log.Fatal(err)
	}
	c := bot.NewClient(resp)
	err = c.JoinServer(viper.GetString("server.address"))
	if err != nil {
		log.Fatal(err)
	}

	_ = basic.NewPlayer(c, basic.DefaultSettings)
	basic.EventsListener{
		GameStart: func() error {
			log.Println("game start")
			return nil
		},
		ChatMsg: func(c chat.Message, pos byte, uuid uuid.UUID) error {
			log.Println(c)
			return nil
		},
		Disconnect: func(reason chat.Message) error {
			log.Printf("disconnected: %s", reason)
			return nil
		},
		HealthChange: func(health float32) error {
			log.Printf("health updated: %.2f", health)
			return nil
		},
		Death: func() error {
			log.Println("you died")
			return nil
		},
	}.Attach(c)

	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}
