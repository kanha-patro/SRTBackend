package rbac

type Role string

const (
    AdminRole        Role = "admin"
    OrganisationRole Role = "organisation"
    DriverRole       Role = "driver"
    UserRole         Role = "user"
)

func GetAllRoles() []Role {
    return []Role{
        AdminRole,
        OrganisationRole,
        DriverRole,
        UserRole,
    }
}