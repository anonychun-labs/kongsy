package main

import (
	"context"
	"kongsy"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "start",
		Usage: "start the server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "host",
				Usage:    "server host",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "limit",
				Usage:    "limit requests per minute",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "interval",
				Usage:    "interval in seconds",
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return kongsy.Start(
				cmd.String("host"),
				int(cmd.Int("limit")),
				int(cmd.Int("interval")),
			)
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
