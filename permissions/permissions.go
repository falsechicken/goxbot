package permissions

var permsTable map[string][]string

func init() {
	permsTable = make(map[string][]string)
}

func HasPermission(userName string, permission string) bool {
	return true
}
