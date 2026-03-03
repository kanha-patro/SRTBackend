package rbac

type Permission string

const (
    PermissionOrgCreate       Permission = "org:create"
    PermissionOrgRead         Permission = "org:read"
    PermissionOrgUpdate       Permission = "org:update"
    PermissionOrgDelete       Permission = "org:delete"
    
    PermissionRouteCreate     Permission = "route:create"
    PermissionRouteRead       Permission = "route:read"
    PermissionRouteUpdate     Permission = "route:update"
    PermissionRouteDelete     Permission = "route:delete"
    
    PermissionDriverAssign     Permission = "driver:assign"
    PermissionDriverUnassign   Permission = "driver:unassign"
    
    PermissionTripStart        Permission = "trip:start"
    PermissionTripEnd          Permission = "trip:end"
    PermissionTripRead         Permission = "trip:read"
    
    PermissionUserTrack        Permission = "user:track"
)

var AllPermissions = []Permission{
    PermissionOrgCreate,
    PermissionOrgRead,
    PermissionOrgUpdate,
    PermissionOrgDelete,
    PermissionRouteCreate,
    PermissionRouteRead,
    PermissionRouteUpdate,
    PermissionRouteDelete,
    PermissionDriverAssign,
    PermissionDriverUnassign,
    PermissionTripStart,
    PermissionTripEnd,
    PermissionTripRead,
    PermissionUserTrack,
}