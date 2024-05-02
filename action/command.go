package action

import (
	"context"

	"github.com/angrifel/gogen/action/internal/handler"
	"github.com/urfave/cli/v3"
)

func HandlerCommand(ctx context.Context, command *cli.Command) error {
	return handler.Action(ctx, command)
}
