package certificates

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"
)

// OPENSSL COFIGURATION
// [ req ]
// default_bits        = 4096
// distinguished_name  = req_distinguished_name
// string_mask         = utf8only
// prompt              = no
// default_md          = sha256

// [ req_distinguished_name ]
// commonName=api.auth

func GenerateCertificate(certPath *string, keyPath *string) error {
	// Create folders if not exists
	if err := createFolderIfNotExists(*certPath); err != nil {
		return errors.New("error creating certificate folder: " + err.Error())
	}
	if err := createFolderIfNotExists(*keyPath); err != nil {
		return errors.New("error creating key folder: " + err.Error())
	}

	// Generate RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return errors.New("error generating RSA key: " + err.Error())
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "api.auth",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 1 years
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return errors.New("error creating certificate: " + err.Error())
	}

	// Save private key to file
	keyFile, err := os.Create(*keyPath + "/auth.key")
	if err != nil {
		return errors.New("error creating key file: " + err.Error())
	}
	defer keyFile.Close()

	err = pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		return errors.New("error encoding private key: " + err.Error())
	}

	// Save certificate to file
	certFile, err := os.Create(*certPath + "/auth.crt")
	if err != nil {
		return errors.New("error creating certificate file: " + err.Error())
	}
	defer certFile.Close()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	if err != nil {
		return errors.New("error encoding certificate: " + err.Error())
	}

	fmt.Println("Certificate generated successfully")
	return nil
}

func createFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return errors.New("error creating folder: " + err.Error())
		}
	}
	return nil
}
