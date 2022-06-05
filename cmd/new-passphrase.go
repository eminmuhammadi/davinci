package cmd

import (
	fs "github.com/eminmuhammadi/davinci/fs"
	generator "github.com/eminmuhammadi/davinci/generator"
	cli "github.com/urfave/cli/v2"
)

func NewPassphrase() *cli.Command {
	return &cli.Command{
		Name:    "new-passphrase",
		Aliases: []string{"passphrase"},
		Usage:   "Generate new passphrase file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "folder",
				Aliases:     []string{"f"},
				Usage:       "The folder to store the passphrase file",
				DefaultText: "./",
				Required:    true,
			},
		},
		Action: newPassphraseAction,
	}
}

func newPassphraseAction(ctx *cli.Context) error {
	folder := ctx.String("folder")

	// Generate a random passphrase
	passphrase := generator.RandomBytesBase64(32)

	fs.WriteFile(
		fs.JoinPaths(folder, "passphrase.txt"),
		[]byte(passphrase),
	)

	return nil
}
