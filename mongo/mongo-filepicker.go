package mongo

import (
	"cre/styles"

	"github.com/charmbracelet/huh"
)

// type MongoTool struct {
// 	name string
// 	Fn   func() string
// }



func MongoFilePicker(path *string) *huh.Form {

	var form = huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Picking(true).
				Title("Credentials File").
				Description("Select a .yaml file").
				AllowedTypes([]string{".yaml"}).
				// Key("credentialsFile").
				Height(20).
				CurrentDirectory(*path).
				ShowHidden(false).
				Value(path),
		),
	).
		WithShowHelp(true).
		WithTheme(styles.ThemeDag()).
		WithHeight(20)

	return form
}

// func MongoInquireForm(mongoCmd *MongoCmd) *huh.Form {

// 	mongoTools := []MongoTool{
// 		{"Create Mongo User", mongoCmd.CreateUser},
// 		{"Create Mongo Root User", mongoCmd.CreateRootUser},
// 		{"Mongo Connection String", mongoCmd.ConnectionString},
// 		{"Drop Mongo User", mongoCmd.DropUser},
// 		{"Authenticate Mongo User", mongoCmd.Authenticate},
// 		{"Change Mongo User Password", mongoCmd.ChangePassword},
// 		{"Grant Role to User", mongoCmd.GrantRolesToUser},
// 		{"MongoShell Cmd", mongoCmd.MongoShellCmd},
// 	}


// 	form := huh.NewForm(
// 		huh.NewGroup(
// 			huh.NewSelect[*func() string]().
// 				Title("Tool selection").
// 				Description("select the utility to parse credentials data").
// 				// Options(
// 				// 	huh.NewOption("Create Mongo User", "CreateUser"),
// 				// 	huh.NewOption("Create Mongo Root User", "CreateRootUser"),
// 				// 	huh.NewOption("Mongo Connection String", "ConnectionString"),
// 				// 	huh.NewOption("Drop Mongo User", "DropUser"),
// 				// 	huh.NewOption("Authenticate Mongo User", "Authenticate"),
// 				// 	huh.NewOption("Change Mongo User Password", "ChangePassword"),
// 				// 	huh.NewOption("Grant Role to User", "GrantRolesToUser"),
// 				// 	huh.NewOption("MongoShell Cmd", "MongoShellCmd"),
// 				// ).
// 				OptionsFunc(func() []huh.Option {
// 					var opts []huh.Option
// 					for _, tool := range mongoTools {
// 						opts = append(opts, huh.NewOption(tool.name, tool.Fn))
// 					}
// 					return opts
// 				}).
// 				// Key("tool").
// 				Value(tool).
// 				Height(20),
// 		),
// 	).
// 		WithShowHelp(true).
// 		WithTheme(styles.ThemeDag()).
// 		WithHeight(20)

// 	return form
// }
