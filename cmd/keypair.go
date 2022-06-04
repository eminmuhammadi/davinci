package cmd

import (
	"strconv"

	fs "github.com/eminmuhammadi/davinci/fs"
	rsa "github.com/eminmuhammadi/davinci/rsa"
	cli "github.com/urfave/cli/v2"
)

func NewKeyPair() *cli.Command {
	return &cli.Command{
		Name:    "new-keypair",
		Aliases: []string{"keypair"},
		Usage:   "Generate new public and private key pair",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "size",
				Aliases:     []string{"s"},
				Usage:       "Key size",
				DefaultText: "2048",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "folder",
				Aliases:     []string{"f"},
				Usage:       "Folder to save the key files",
				DefaultText: "./",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "passphrase",
				Aliases:     []string{"p"},
				Usage:       "Passphrase to decrypt the keys",
				DefaultText: "./passphrase.key",
				Required:    true,
			},
		},
		Action: newKeysAction,
	}
}

func newKeysAction(ctx *cli.Context) error {
	folder := ctx.String("file")
	passphrase := fs.ReadFile(ctx.String("passphrase"))

	keySize := ctx.String("size")

	size, err := strconv.Atoi(keySize)
	if err != nil {
		return err
	}

	keys := rsa.GenerateKeyPair(size, string(passphrase))

	fs.WriteFile(
		fs.JoinPaths(folder, "publicKey.pem"),
		[]byte(keys.PublicKey),
	)
	fs.WriteFile(
		fs.JoinPaths(folder, "privateKey.pem"),
		[]byte(keys.PrivateKey),
	)

	return nil
}
