package rbac

import (
	"errors"
)

type Role string

const (
	AdminRole        Role = "admin"
	OrganisationRole Role = "organisation"
	DriverRole       Role = "driver"
	UserRole         Role = "user"
)

type Permission string

const (
	ApproveOrgPermission       Permission = "approve_org"
	SuspendOrgPermission       Permission = "suspend_org"
	MonitorTripsPermission     Permission = "monitor_trips"
	ForceStopTripPermission    Permission = "force_stop_trip"
	RevokeOTPSessionPermission Permission = "revoke_otp_session"
	CreateRoutePermission      Permission = "create_route"
	EditRoutePermission        Permission = "edit_route"
	AssignDriverPermission      Permission = "assign_driver"
	UnassignDriverPermission    Permission = "unassign_driver"
	TrackActiveShuttlesPermission Permission = "track_active_shuttles"
)

var rolePermissions = map[Role][]Permission{
	AdminRole: {
		ApproveOrgPermission,
		SuspendOrgPermission,
		MonitorTripsPermission,
		ForceStopTripPermission,
		RevokeOTPSessionPermission,
	},
	OrganisationRole: {
		CreateRoutePermission,
		EditRoutePermission,
		AssignDriverPermission,
		UnassignDriverPermission,
	},
	DriverRole: {
		AssignDriverPermission,
		UnassignDriverPermission,
	},
	UserRole: {
		TrackActiveShuttlesPermission,
	},
}

type Enforcer struct {
	role Role
}

func NewEnforcer(role Role) *Enforcer {
	return &Enforcer{role: role}
}

func (e *Enforcer) HasPermission(permission Permission) (bool, error) {
	permissions, ok := rolePermissions[e.role]
	if !ok {
		return false, errors.New("role does not exist")
	}

	for _, p := range permissions {
		if p == permission {
			return true, nil
		}
	}
	return false, nil
}