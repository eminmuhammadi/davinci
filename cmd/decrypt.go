package cmd

import (
	"bytes"

	aes "github.com/eminmuhammadi/davinci/aes"
	fs "github.com/eminmuhammadi/davinci/fs"
	rsa "github.com/eminmuhammadi/davinci/rsa"
	cli "github.com/urfave/cli/v2"
)

func Decrypt() *cli.Command {
	return &cli.Command{
		Name:    "decrypt",
		Aliases: []string{"dec"},
		Usage:   "Decrypt the data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "Output file",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "private-key",
				Aliases:     []string{"privK"},
				Usage:       "Private key file",
				DefaultText: "./privateKey.pem",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "passphrase",
				Aliases:     []string{"p"},
				Usage:       "Passphrase to decrypt the keys",
				DefaultText: "./passphrase.txt",
				Required:    true,
			},
		},
		Action: decryptAction,
	}
}

func decryptAction(ctx *cli.Context) error {
	privateKey := fs.ReadFile(ctx.String("private-key"))
	passphrase := fs.ReadFile(ctx.String("passphrase"))
	input := fs.ReadFile(ctx.String("input"))
	output := ctx.String("output")

	// Split the input to get the encrypted key and the encrypted data
	splittedInput := bytes.Split(input, []byte("\n"))
	encKey, ciphertext := splittedInput[0], splittedInput[1]

	// Decrypt the key
	key := rsa.Decrypt(string(privateKey), string(encKey), string(passphrase))

	// Decrypt the data
	plaintext := aes.DecryptGCM256(string(ciphertext), string(key))

	fs.WriteFile(
		output,
		[]byte(plaintext),
	)

	return nil
}
