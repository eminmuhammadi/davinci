package cmd

import (
	fs "github.com/eminmuhammadi/davinci/fs"
	generator "github.com/eminmuhammadi/davinci/generator"
	cli "github.com/urfave/cli/v2"
)

func NewKey() *cli.Command {
	return &cli.Command{
		Name:    "new-key",
		Aliases: []string{"key"},
		Usage:   "Generate new symmetric key",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "folder",
				Aliases:     []string{"f"},
				Usage:       "Folder to save the key file",
				DefaultText: "./",
				Required:    true,
			},
		},
		Action: newKeyAction,
	}
}

func newKeyAction(ctx *cli.Context) error {
	folder := ctx.String("folder")

	// Generate new key (size 32 bytes)
	key := generator.RandomBytesBase64(32)

	fs.WriteFile(
		fs.JoinPaths(folder, "key.txt"),
		[]byte(key),
	)

	return nil
}
