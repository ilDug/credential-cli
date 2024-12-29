package main

import (
	"cre/core"
	"cre/mongo"
	"cre/styles"
	"fmt"
)


func main() {
	fmt.Print(styles.TitleStyle.Render("Credentials Manager"))

    var command string = core.SelectCmd()

	switch command {
		case "MONGO":
			mongo.MongoRun()

		case "CERTIFICATE":
			fmt.Println(styles.CommandStyle.Render("Certificate Manager... generating certificate/key pair"))
	}
}