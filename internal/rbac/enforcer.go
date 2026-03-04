package rbac

import "errors"

var rolePermissions = map[Role][]Permission{
	AdminRole: {
		PermissionOrgCreate,
		PermissionOrgRead,
		PermissionOrgUpdate,
		PermissionOrgDelete,
		PermissionRouteRead,
		PermissionTripRead,
	},

	OrganisationRole: {
		PermissionRouteCreate,
		PermissionRouteRead,
		PermissionRouteUpdate,
		PermissionRouteDelete,
		PermissionDriverAssign,
		PermissionDriverUnassign,
		PermissionTripRead,
	},

	DriverRole: {
		PermissionTripStart,
		PermissionTripEnd,
		PermissionTripRead,
	},

	UserRole: {
		PermissionUserTrack,
		PermissionTripRead,
	},
}

type Enforcer struct {
	role Role
}

func NewEnforcer(role Role) (*Enforcer, error) {
	if !isValidRole(role) {
		return nil, errors.New("invalid role")
	}
	return &Enforcer{role: role}, nil
}

func (e *Enforcer) HasPermission(permission Permission) bool {
	permissions := rolePermissions[e.role]

	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func isValidRole(role Role) bool {
	for _, r := range GetAllRoles() {
		if r == role {
			return true
		}
	}
	return false
}