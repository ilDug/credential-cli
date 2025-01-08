package mongo

import (
	"cre/styles"

	"github.com/charmbracelet/huh"
)

// type mongoTool struct {
// 	Description string
// 	Command     func()
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

func MongoInquireForm(tool *string) *huh.Form {
	// read the file contents of yaml file
	// yamlFileContent, err := os.ReadFile(*credentialsFile)
	// if err != nil {
	// 	e := errors.New("error reading file")
	// 	errors.Join(e, err)
	// 	panic(e)
	// }

	// // Parse the YAML file
	// var credentials MongoCredentials
	// err = yaml.Unmarshal(yamlFileContent, &credentials)
	// if err != nil {
	// 	e := errors.New("error parsing file")
	// 	errors.Join(e, err)
	// 	panic(e)
	// }

	// cmd := MongoCmd{credentials: &credentials}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Tool selection").
				Description("select the utility to parse credentials data").
				Options(
					huh.NewOption("Create Mongo User", "CreateUser"),
					huh.NewOption("Create Mongo Root User", "CreateRootUser"),
					huh.NewOption("Mongo Connection String", "ConnectionString"),
					huh.NewOption("Drop Mongo User", "DropUser"),
					huh.NewOption("Authenticate Mongo User", "Authenticate"),
					huh.NewOption("Change Mongo User Password", "ChangePassword"),
					huh.NewOption("Grant Role to User", "GrantRolesToUser"),
					huh.NewOption("MongoShell Cmd", "MongoShellCmd"),
				).
				// Key("tool").
                Value(tool).
				Height(20),
		),
	).
		WithShowHelp(true).
		WithTheme(styles.ThemeDag()).
		WithHeight(20)

	return form
}
