package command

import (
	"context"

	"github.com/PullRequestInc/go-gpt3"
)

type onceCommand struct {
	client  gpt3.Client
	model   string
	message string
}

func (c onceCommand) Run(ctx context.Context) error {
	return doCompletionRequest(ctx, c.client, c.model, c.message)
}

func NewOnceCommand(client gpt3.Client, model, message string) onceCommand {
	return onceCommand{
		client:  client,
		model:   model,
		message: message,
	}
}
