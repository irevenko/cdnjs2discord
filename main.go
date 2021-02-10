package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	c "./commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("DS_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(c.HelpCommand)
	dg.AddHandler(c.LibCommand)
	dg.AddHandler(c.StatsCommand)
	dg.AddHandler(c.AssetsCommand)
	dg.AddHandler(c.WhiteListCommand)
	dg.AddHandler(c.SearchNameCommand)



	dg.AddHandler(func(dg *discordgo.Session, ready *discordgo.Ready) {
		err = dg.UpdateGameStatus(0, "!cdn")
		if err != nil {
			log.Fatal("Status error")
		}
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
