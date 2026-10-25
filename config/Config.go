package config

var Config = struct {
	AllowedUsersList []string
}{}

func SetAllowedUsers(allowedUsers []string) {
	Config.AllowedUsersList = allowedUsers
}
