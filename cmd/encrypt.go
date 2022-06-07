package cmd

import (
	aes "github.com/eminmuhammadi/davinci/aes"
	fs "github.com/eminmuhammadi/davinci/fs"
	rsa "github.com/eminmuhammadi/davinci/rsa"
	cli "github.com/urfave/cli/v2"
)

func Encrypt() *cli.Command {
	return &cli.Command{
		Name:    "encrypt",
		Aliases: []string{"enc"},
		Usage:   "Encrypt the data",
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
				Name:        "public-key",
				Aliases:     []string{"pubK"},
				Usage:       "Public key file",
				DefaultText: "./publicKey.pem",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "key",
				Aliases:     []string{"k"},
				Usage:       "Symmetric key file",
				DefaultText: "./key.txt",
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
		Action: func(ctx *cli.Context) error {
			pubK := fs.ReadFile(ctx.String("public-key"))
			passphrase := fs.ReadFile(ctx.String("passphrase"))
			key := fs.ReadFile(ctx.String("key"))
			input := fs.ReadFile(ctx.String("input"))
			output := fs.JoinPaths(ctx.String("output"))

			ciphertext := EncryptAction(pubK, passphrase, key, input)

			err := fs.WriteFile(output, []byte(ciphertext))

			return err
		},
	}
}

func EncryptAction(pubK []byte, passphrase []byte, key []byte, input []byte) string {
	// Encrypt the key
	encKey := rsa.Encrypt(string(pubK), string(key), string(passphrase))

	// Encrypt the data
	ciphertext := aes.EncryptGCM256(string(input), string(key))

	// Structure
	// -------------
	// Encrypted Key
	//
	// Encrypted Data
	return encKey + "\n" + ciphertext
}
