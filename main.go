package main

import (
	"cre/core"
	"cre/mongo"
	"cre/styles"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	var command string

	// Define command-line flags with short alternatives
	var out string
	flag.StringVar(&out, "out", "./secrets", "Output directory")
	flag.StringVar(&out, "o", "./secrets", "Output directory (shorthand)")

	var mongoCredentialsFile string
	flag.StringVar(&mongoCredentialsFile, "credentials", "", "file of credentials")
	flag.StringVar(&mongoCredentialsFile, "c", "", "file of credentials (shorthand)")

	flag.Parse()

	// Get the command argument
	args := flag.Args()
	if len(args) == 0 {
		command = core.SelectCmd()
	} else if len(args) > 1 {
		log.Fatal("Too many arguments.", "\nSOLUTION", "Please put OPTIONS before ARGUMENTS, or be sure to provide one or no argument.")
	} else {
		command = strings.ToUpper(args[0])
		if !slices.Contains(core.AvailableCommands, command) {
			command = core.SelectCmd()
		}
	}

	////////////////////////////////////////////////////////////////////////////////////////
	// Print the title banner
	fmt.Print(styles.TitleStyle.Render("Credentials Manager"))

	// Select the command to run
	switch command {
	case "HELP":
		fmt.Println(styles.BoxStyle.Render(core.Help))

	case "MONGO":

		if mongoCredentialsFile != "" {
			mongo.MongoCommandSelect(mongoCredentialsFile)
		} else {
			mongo.MongoRun(out)
		}

	case "CERTIFICATE":
		fmt.Println(styles.CommandStyle.Render("Certificate Manager... generating certificate/key pair"))

		// // Define paths for keys and certificates
		// keysPath := fmt.Sprintf("%s/auth.key", out)
		// certsPath := fmt.Sprintf("%s/auth.crt", out)

		// // Generate RSA key
		// privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
		// if err != nil {
		// 	log.Fatalf("Failed to generate private key: %v", err)
		// }

		// // Create a template for the certificate
		// template := x509.Certificate{
		// 	SerialNumber: big.NewInt(1),
		// 	Subject: pkix.Name{
		// 		Organization: []string{"Your Organization"},
		// 	},
		// 	NotBefore:             time.Now(),
		// 	NotAfter:              time.Now().AddDate(10, 0, 0), // 10 years
		// 	KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		// 	ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		// 	BasicConstraintsValid: true,
		// }

		// // Create the certificate
		// certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
		// if err != nil {
		// 	log.Fatalf("Failed to create certificate: %v", err)
		// }

		// // Save the certificate to a file
		// certOut, err := os.Create(certsPath)
		// if err != nil {
		// 	log.Fatalf("Failed to open cert file for writing: %v", err)
		// }
		// pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
		// certOut.Close()

		// // Save the private key to a file
		// keyOut, err := os.Create(keysPath)
		// if err != nil {
		// 	log.Fatalf("Failed to open key file for writing: %v", err)
		// }
		// pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
		// keyOut.Close()

		// fmt.Println("Certificate and key generated successfully.")
	}
}
