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

	var defaultPerms = new(PermsTable)
	defaultPerms.User = make([]User, 2, 2)

	defaultPerms.User = append(defaultPerms.User, *new(User), *new(User))
	defaultPerms.User[0].JID = "user@example.net"
	defaultPerms.User[0].Groups = make([]string, 2)
	defaultPerms.User[0].Groups[0] = "admin"
	defaultPerms.User[0].Groups[1] = "default"
	defaultPerms.User[0].Permissions = make([]string, 2)
	defaultPerms.User[0].Permissions[0] = "status"
	defaultPerms.User[0].Permissions[1] = "test"

	defaultPerms.User[1].JID = "user2@example2.net"
	defaultPerms.User[1].Groups = make([]string, 2)
	defaultPerms.User[1].Groups[0] = "user"
	defaultPerms.User[1].Groups[1] = "default"
	defaultPerms.User[1].Permissions = make([]string, 2)
	defaultPerms.User[1].Permissions[0] = "exampleCmd"
	defaultPerms.User[1].Permissions[1] = "superAbility"

	perms, err := os.Create(path)
	if err != nil {
		glogger.LogMessage(glogger.Error, "Cannot create permissions file!: "+path)
		panic(err)
	}
	defer perms.Close()

	encoder := toml.NewEncoder(perms)

	encoder.Encode(defaultPerms)

	return true
}
