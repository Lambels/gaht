package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

    "github.com/Lambels/gaht/command"
	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("missing api key")
	}

	client := gpt3.NewClient(apiKey)

	var message string
    var engine string
	fs := flag.NewFlagSet("gaht", flag.PanicOnError)
	fs.StringVar(&message, "m", "", "sends one message to open ai without starting a dialogue")
    fs.StringVar(&engine, "e", gpt3.TextDavinci003Engine, "choose the open ai engine to work on your messages")
	fs.Parse(os.Args[1:])

    var cmd command.Runner
    switch message {
    case "":
        cmd = command.NewDialogueCommand(client, engine)
    default:
        cmd = command.NewOnceCommand(client, engine, message)
    }

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := cmd.Run(ctx); err != nil {
		panic(err)
	}
}
