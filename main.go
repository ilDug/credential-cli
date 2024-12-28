package main

import (
	"cre/core"
	"cre/styles"
	"fmt"
)


func main() {
	fmt.Print(styles.TitleStyle.Render("Credentials Manager"))

    var command string = core.SelectCommand()

	switch command {
		case "MONGO":
			fmt.Println(styles.CommandStyle.Render("MongoDB Utilities... creating mongo credentials"))
		case "CERTIFICATE":
			fmt.Println(styles.CommandStyle.Render("Certificate Manager... generating certificate/key pair"))
	}
}