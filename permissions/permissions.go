package permissions

//permsTable Is a mapping of the permissions held by users. Jids are keys to a list of strings that represent permissions.
var permsTable map[string][]string

func init() {
	permsTable = make(map[string][]string)
}

func HasPermission(jid string, permission string) bool {
	return true
}

func GrantPermission(jid string, permission string) {

}

func RevokePermission(jid string, permission string) {

}
