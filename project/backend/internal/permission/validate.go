package permission

// IsValid valida if the permission is valid.
func IsValid(permission byte) bool {
	for _, value := range arrayPermissions {
		if value == Permission(permission) {
			return true
		}
	}
	return false
}
