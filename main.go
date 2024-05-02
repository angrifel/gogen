package main

import (
	"context"
	"fmt"
	"os"

	"github.com/angrifel/gogen/action"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name: "gogen",
		Commands: []*cli.Command{
			{
				Name:   "handler",
				Action: action.HandlerCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "path", Value: "", Required: true},
					&cli.BoolFlag{Name: "force", Value: false},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
}
