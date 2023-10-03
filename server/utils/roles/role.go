package roles

const (
	accountService   string = "/pb.Account/"
	messagingService string = "/pb.Messaging/"
)

var accessRoles = map[string][]Roles{}

// HasPermission return true if user has permission in that method.
func HasPermission(fullMethod string, userRoles Roles) bool {
	if roles, ok := accessRoles[fullMethod]; ok {
		for _, role := range roles {
			if role == userRoles {
				return true
			}
		}
	} else {
		return true // This method will be skipped because it does not require to check
	}
	return false
}

// IsNeedToCheckMethodAccess is check whether the method need to check or not
// if yes return true otherwise, return false.
func IsNeedToCheckMethodAccess(fullMethod string) bool {
	_, ok := accessRoles[fullMethod]
	return ok
}

func init() {
	// account services
	{
		accessRoles[accountService+"ReNewToken"] = []Roles{Roles_USER, Roles_ADMIN, Roles_LEADER}
	}

	// messaging service
	{
		accessRoles[messagingService+"GetVerifyToken"] = []Roles{Roles_UNSPECIFIED_USER}
	}
}
