package curl

import (
	"context"
	"os/exec"
)

func Command(ctx context.Context, url string, options ...Option) *exec.Cmd {
	args := make([]string, 0, len(options))
	args = append(args, url)

	for _, opt := range options {
		if arg := opt.String(); arg != "" {
			args = append(args, arg)
		}
	}

	return exec.CommandContext(ctx, "curl", args...)
}
