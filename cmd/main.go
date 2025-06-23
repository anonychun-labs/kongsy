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
				Name:     "to",
				Usage:    "remote host to proxy requests to",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "port",
				Usage:    "port to listen on",
				Value:    8080,
				Required: false,
			},
			&cli.IntFlag{
				Name:     "limit",
				Usage:    "limit requests per-ip",
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
				cmd.String("to"),
				cmd.Int("port"),
				cmd.Int("limit"),
				cmd.Int("interval"),
			)
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
