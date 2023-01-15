package command

import (
	"context"
	"fmt"

	"github.com/PullRequestInc/go-gpt3"
)

type Runner interface {
    Run(context.Context) error
}

func doCompletionRequest(ctx context.Context, client gpt3.Client, model, message string) error {
    return client.CompletionStreamWithEngine(
       ctx,
       model,
       gpt3.CompletionRequest{
            Prompt: []string{ message },        
            Temperature: gpt3.Float32Ptr(0),
            MaxTokens: gpt3.IntPtr(3000),
       },
       func(resp *gpt3.CompletionResponse) {
           fmt.Print(resp.Choices[0].Text)
       },
    )
}
