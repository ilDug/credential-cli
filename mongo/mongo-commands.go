package mongo

import (
	"bytes"
	"text/template"
)

type MongoCmd struct {
	credentials *MongoCredentials
}

func (m *MongoCmd) renderTemplate(tmpl string) string {
	var result bytes.Buffer
	t, _ := template.New("createUser").Parse(tmpl)
	_ = t.Execute(&result, m.credentials)
	return result.String()
}

func (m *MongoCmd) CreateUser() string {
	const tmpl = `
use admin

db.createUser({
	user: "{{.Username}}",
	pwd: "{{.Password}}",
	roles: [
		{ db: "{{.Database}}", role: "readWrite" },
		{ db: "{{.Database}}", role: "dbAdmin" },
	]
})
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) CreateRootUser() string {
	const tmpl = `
use admin

db.createUser({
	user: "{{.Username}}",
	pwd: "{{.Password}}",
	roles: [
		{ db: "{{.Database}}", role: "userAdminAnyDatabase" },
		{ db: "{{.Database}}", role: "dbAdminAnyDatabase" },
		{ db: "{{.Database}}", role: "clusterAdmin" },
		{ db: "{{.Database}}", role: "root" },
	]
})
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) ConnectionString() string {
	const tmpl = `
username: 
    {{.Username}}

password: 
    {{.Password}}

database: 
    {{.Database}}

host: 
    {{.Host}}

connection string: 
    {{.ConnectionString}}
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) DropUser() string {
	const tmpl = `
use admin
db.dropUser("{{.Username}}")
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) Authenticate() string {
	const tmpl = `
use admin
db.auth("{{.Username}}", "{{.Password}}")
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) ChangePassword() string {
	const tmpl = `
use admin
db.changeUserPassword("{{.Username}}", "{{.Password}}")
db.changeUserPassword("{{.Username}}", passwordPrompt())
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) GrantRolesToUser() string {
	const tmpl = `
db.grantRolesToUser(
	"{{.Username}}",
	[
		{ db: "{{.Database}}", role: "readWrite" }
		{ db: "{{.Database}}", role: "userAdmin" },
		{ db: "{{.Database}}", role: "dbAdmin" }
		{ db: "{{.Database}}", role: "dbOwner" }
	]
)
`
	return m.renderTemplate(tmpl)
}

func (m *MongoCmd) MongoShellCmd() string {
	const tmpl = `
mongosh "mongodb://authuser:nd0iYRc52XiPEDVh5h4FwTHXJwRMaRW8Szqkc6hmYdAZiyjsLJ41XeGJOkI3kQ9z@mongo1.dag.lan:27017,mongo2.dag.lan:27017,mongo3.dag.lan:27017/auth?authSource=admin&replicaSet=rs0"
`
	return m.renderTemplate(tmpl)
}
