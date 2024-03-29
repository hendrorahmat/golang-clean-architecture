package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rsaGenerator = &cobra.Command{
	Use:   "generate:credentials {name}",
	Args:  cobra.MinimumNArgs(1),
	Short: "Generate public secret key",
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate key
		privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			fmt.Printf("Cannot generate RSA key\n")
			os.Exit(1)
		}
		publickey := &privatekey.PublicKey

		// dump private key to file
		var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
		privateKeyBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		}
		privatePem, err := os.Create("storage/" + args[0] + "-" + "private.key")
		if err != nil {
			fmt.Printf("error when create private.pem: %s \n", err)
			os.Exit(1)
		}
		err = pem.Encode(privatePem, privateKeyBlock)
		if err != nil {
			fmt.Printf("error when encode private pem: %s \n", err)
			os.Exit(1)
		}

		// dump public key to file
		publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
		if err != nil {
			fmt.Printf("error when dumping publickey: %s \n", err)
			os.Exit(1)
		}
		publicKeyBlock := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		}
		publicPem, err := os.Create("storage/" + args[0] + "-" + "public.key")
		if err != nil {
			fmt.Printf("error when create public.pem: %s \n", err)
			os.Exit(1)
		}
		err = pem.Encode(publicPem, publicKeyBlock)
		if err != nil {
			fmt.Printf("error when encode public pem: %s \n", err)
			os.Exit(1)
		}
		return nil
	},
}
