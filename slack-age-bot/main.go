package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command: ", event.Command)
		fmt.Println("Timestamp: ", event.Timestamp)
		fmt.Println("Parameters: ", event.Parameters)
		fmt.Println("Event: ", event.Event)

	}
}

func main() {
	loadEnvFile()
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	bot := slacker.NewClient(botToken, appToken)

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
}
