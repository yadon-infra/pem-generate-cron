package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := flag.String("dir", ".", "path of directory")
	flag.Parse()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generating key: %v", err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		err = os.MkdirAll(*dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	privateKeyPath := filepath.Join(*dir, "private-key.pem")

	err = os.Remove("private-key.pem")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed remove old pemfile: %v", err)
	}

	err = os.WriteFile(privateKeyPath, privateKeyPEM, 0600)
	if err != nil {
		log.Fatalf("Failed to save privateKey to pemfile: %v", err)
	}

	log.Printf("Private key saved to %s", privateKeyPath)
}
