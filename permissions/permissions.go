package permissions

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/falsechicken/glogger"
)

var permsTable PermsTable
var groupsTable GroupsTable

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
	Users []User
}

type GroupsTable struct {
	Groups []PermGroup
}

func init() {

	permsTable := new(PermsTable)
	permsTable.Users = make([]User, 2, 10)

}

// HasPermission returns true if a jid has the provided permission.
func HasPermission(jid string, permission string) bool {

	jid = strings.Split(jid, "/")[0]

	glogger.LogMessage(glogger.Debug, "Checking permission "+permission+" for user "+jid)

	var userFound = false
	var userPerms *[]string
	for _, user := range permsTable.Users {
		if user.JID == jid {
			userFound = true
			userPerms = &user.Permissions
		}
		if userFound == true {
			break
		}
	}

	if userFound {
		for _, perm := range *userPerms {
			if perm == permission {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

func GrantPermission(jid string, permission string) {

}

func RevokePermission(jid string, permission string) {

}

func Load(path string) bool {
	if _, err := os.Stat(path + "/perms.toml"); err == nil {
		if _, err := toml.DecodeFile(path+"/perms.toml", &permsTable); err != nil {
			glogger.LogMessage(glogger.Error, err.Error())
			os.Exit(2)
		}
	} else {
		generateDefaultConfig(path)
		Load(path)
	}

	if _, err := os.Stat(path + "/groups.toml"); err == nil {
		if _, err := toml.DecodeFile(path+"/groups.toml", &groupsTable); err != nil {
			glogger.LogMessage(glogger.Error, err.Error())
			os.Exit(2)
		}
	} else {
		generateDefaultGroups(path)
		Load(path)
	}
	return true
}

func Save(path string) bool {

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

func generateDefaultConfig(path string) bool {
	glogger.LogMessage(glogger.Info, "Generating default permissions file...")

	var defaultPerms = new(PermsTable)
	defaultPerms.Users = make([]User, 2, 2)

	defaultPerms.Users = append(defaultPerms.Users, *new(User), *new(User))
	defaultPerms.Users[0].JID = "user@example.net"
	defaultPerms.Users[0].Groups = make([]string, 2)
	defaultPerms.Users[0].Groups[0] = "admin"
	defaultPerms.Users[0].Groups[1] = "default"
	defaultPerms.Users[0].Permissions = make([]string, 2)
	defaultPerms.Users[0].Permissions[0] = "status"
	defaultPerms.Users[0].Permissions[1] = "test"

	defaultPerms.Users[1].JID = "user2@example2.net"
	defaultPerms.Users[1].Groups = make([]string, 2)
	defaultPerms.Users[1].Groups[0] = "user"
	defaultPerms.Users[1].Groups[1] = "default"
	defaultPerms.Users[1].Permissions = make([]string, 2)
	defaultPerms.Users[1].Permissions[0] = "exampleCmd"
	defaultPerms.Users[1].Permissions[1] = "superAbility"

	perms, err := os.Create(path + "/perms.toml")
	if err != nil {
		glogger.LogMessage(glogger.Error, "Cannot create permissions file!: "+path+"/perms.toml")
		panic(err)
	}
	defer perms.Close()

	encoder := toml.NewEncoder(perms)

	encoder.Encode(defaultPerms)

	return true
}

func generateDefaultGroups(path string) bool {
	glogger.LogMessage(glogger.Info, "Generating default groups file...")

	var defaultGroups = new(GroupsTable)
	defaultGroups.Groups = make([]PermGroup, 1, 1)

	defaultGroups.Groups = append(defaultGroups.Groups, *new(PermGroup), *new(PermGroup))

	defaultGroups.Groups[0].Name = "default"
	defaultGroups.Groups[0].Permissions = make([]string, 2, 2)
	defaultGroups.Groups[0].Permissions[0] = "talk"
	defaultGroups.Groups[0].Permissions[1] = "status"

	defaultGroups.Groups[1].Name = "user"
	defaultGroups.Groups[1].Permissions = make([]string, 3, 3)
	defaultGroups.Groups[1].Permissions[0] = "shout"
	defaultGroups.Groups[1].Permissions[1] = "move"
	defaultGroups.Groups[1].Permissions[2] = "poke"

	defGrps, err := os.Create(path + "/groups.toml")
	if err != nil {
		glogger.LogMessage(glogger.Error, "Cannot create groups file!: "+path+"/groups.toml")
		panic(err)
	}

	encoder := toml.NewEncoder(defGrps)

	encoder.Encode(defaultGroups)

	return true
}
