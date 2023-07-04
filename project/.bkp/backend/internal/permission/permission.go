package permission

// Permission is the type of permissions.
type Permission byte

// Permission constants.
const (
	Ghost Permission = iota
	User
	Employee
	Admin
)

var arrayPermissions = [...]Permission{
	Ghost,
	User,
	Employee,
	Admin,
}
