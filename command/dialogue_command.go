package command

import (
	"context"
	"flag"
	"os"

	"github.com/Lambels/go-dialogue"
	"github.com/PullRequestInc/go-gpt3"
)

type dialogueCommand struct {
    client gpt3.Client
    model string
}

func (c dialogueCommand) Run(ctx context.Context) error {
    d := dialogue.NewDialogue(
        ctx,
        "(prompt) ",
        os.Stdin,
        os.Stdout,
        nil,
    )

    fs := flag.NewFlagSet("send", flag.ContinueOnError)
    msg := fs.String("m", "", "message you want to send")

    d.RegisterCommands(
        &dialogue.Command{
            Name: "send",
            Structure: "send -m <message>",
            HelpLong: "send a message to gpt3 model and waits for the response",
            FlagSet: fs,
            Exec: func(chain *dialogue.CallChain, args []string) error {
                return doCompletionRequest(chain.GetCurrent().Context(), c.client, c.model, *msg)
            },
        },
    )

    return d.Start()
}

func NewDialogueCommand(client gpt3.Client, model string) dialogueCommand {
    return dialogueCommand{
        client: client,
        model: model,
    }
}
