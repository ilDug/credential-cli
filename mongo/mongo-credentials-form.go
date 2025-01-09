package mongo

import (
	"cre/styles"
	"errors"

	"github.com/charmbracelet/huh"
)

func RunCredentialsForm(credentials *MongoCredentials, path *string) huh.Form {

	// // Define the form and its fields and return it
	var form = huh.NewForm(

		// initial note about the program
		huh.NewGroup(huh.NewNote().
			Title("MongoDB Utilities").
			Description("Creating mongoDB credentials\n\n").
			Next(true).
			NextLabel("Start"),
		),

		huh.NewGroup(
			// select the out directory
			huh.NewFilePicker().
				Picking(true).
				Title("Credentials output folder").
				Description("Select a folder [default './secrets']").
				// AllowedTypes([]string{".yaml"}).
				// Key("credentialsFile").
				DirAllowed(true).
				FileAllowed(false).
				CurrentDirectory(*path).
				ShowHidden(false).
				Value(path),
		),

		huh.NewGroup(
			// select whether the user is acting as a ROOT or not
			huh.NewConfirm().
				Title("ROOT USER").
				Description("If the user is acting as a ROOT?, the user will have full access to the database.").
				Value(&credentials.Root).
				Affirmative("root / admin").
				Negative("regular user"),
		),

		huh.NewGroup(
			// input forn the userame
			huh.NewInput().
				Title("Username").
				Placeholder("dag").
				Description("enter the database user").
				Validate(huh.ValidateMinLength(3)).
				Value(&credentials.Username),

			// input forn the database name
			huh.NewInput().
				Title("Database").
				Placeholder("mydb").
				Description("enter the database name").
				Validate(huh.ValidateMinLength(3)).
				Value(&credentials.Database),
		).
			WithHideFunc(func() bool {
				return credentials.Root
			}),

		huh.NewGroup(
			// input forn the hostnames
			huh.NewInput().
				Title("Host").
				Placeholder("mongo1.dag.lan:27017,mongo2.dag.lan:27017,mongo3.dag.lan:27017").
				Description("enter the database host or the replicaset string. Using th pattern the form HOST:PORT").
				Suggestions([]string{
					"mongo1.dag.lan:27017",
					"mongo1.dag.lan:27017,mongo2.dag.lan:27017,mongo3.dag.lan:27017",
				}).
				Validate(func(s string) error {
					if valid, err := validateHostname(s); !valid || err != nil {
						return errors.New("host must be in the format HOST:PORT")
					}
					return nil
				}).
				Value(&credentials.Host),
		),

		huh.NewGroup(
			// input the replicaSet name
			huh.NewInput().
				Title("Replica Set").
				Placeholder("rs0").
				Description("You entered a hostname that seems a replica set. Please enter the replica set name").
				Suggestions([]string{"rs0", "rs1", "rs2"}).
				Validate(func(s string) error {
					if valid, err := validateHostname(s); !valid && err != nil {
						return errors.New("replica set must be in the format HOST:PORT sequence")
					}
					return nil
				}).
				Value(&credentials.ReplicaSet),
		).
			WithHideFunc(func() bool {
				replica, _ := isReplicaSet(credentials.Host)
				return !replica
			}),
	).
		WithHeight(20).
		WithTheme(styles.ThemeDag()).
		WithShowHelp(true)
	// .WithProgramOptions(tea.WithAltScreen())

	return *form
}
