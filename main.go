package main

import (
	"log"
	"os"

	cmd "github.com/eminmuhammadi/davinci/cmd"
	cli "github.com/urfave/cli/v2"
)

const (
	VERSION = "0.1.0"
)

var Commands = []*cli.Command{
	cmd.NewPassphrase(),
	cmd.NewKeyPair(),
	cmd.Decrypt(),
	cmd.Encrypt(),
	cmd.NewKey(),
}

func main() {
	app := &cli.App{
		Name:    "davinci",
		Usage:   "CLI tool for encrypting and decrypting large files",
		Version: VERSION,
		Authors: []*cli.Author{{
			Name:  "Emin Muhammadi",
			Email: "muemin.info@gmail.com",
		}},
		Copyright: "(c) 2022 Emin Muhammadi",
		ExtraInfo: func() map[string]string {
			return map[string]string{
				"LICENSE": "MIT License",
			}
		},
		Commands: Commands,
		Suggest:  true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
