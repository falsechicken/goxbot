package permissions

import (
	"github.com/BurntSushi/toml"
	"github.com/falsechicken/glogger"
	"os"
)

var permsTable PermsTable

type User struct {
	JID         string
	Groups      []string
	Permissions []string
}

type PermGroup struct {
	Name        string
	Permissions []string
}

type PermsTable struct {
	User []User
}

func init() {

	permsTable := new(PermsTable)
	permsTable.User = make([]User, 2, 10)

}

func HasPermission(jid string, permission string) bool {
	return true
}

func GrantPermission(jid string, permission string) {

}

func RevokePermission(jid string, permission string) {

}

func Load(path string) bool {
	generateDefaultConfig(path)
	return true
}

func Save(path string) bool {
	return true
}

func generateDefaultConfig(path string) bool {
	glogger.LogMessage(glogger.Warning, "Generating default permissions file...")

	permsTable.User = append(permsTable.User, *new(User), *new(User))
	permsTable.User[0].JID = "user@example.net"
	permsTable.User[0].Groups = make([]string, 2)
	permsTable.User[0].Groups[0] = "admin"
	permsTable.User[0].Groups[1] = "default"
	permsTable.User[0].Permissions = make([]string, 2)
	permsTable.User[0].Permissions[0] = "status"
	permsTable.User[0].Permissions[1] = "test"

	permsTable.User[1].JID = "user2@example2.net"
	permsTable.User[1].Groups = make([]string, 2)
	permsTable.User[1].Groups[0] = "user"
	permsTable.User[1].Groups[1] = "default"
	permsTable.User[1].Permissions = make([]string, 2)
	permsTable.User[1].Permissions[0] = "exampleCmd"
	permsTable.User[1].Permissions[1] = "superAbility"

	perms, err := os.Create(path)
	if err != nil {
		glogger.LogMessage(glogger.Error, "Cannot create permissions file!: "+path)
		panic(err)
	}
	defer perms.Close()

	encoder := toml.NewEncoder(perms)

	encoder.Encode(permsTable)

	return true
}
