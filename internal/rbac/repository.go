package rbac

import "gorm.io/gorm"


type Repository interface {
    GetPermissionsByRoleNames(roleNames []string) ([]string, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) GetPermissionsByRoleNames(roleNames []string) ([]string, error) {
    var permissions []Permission

    err := r.db.
        Model(&Permission{}).
        Distinct().
        Joins("JOIN role_permissions rp ON rp.permission_id = permissions.id").
        Joins("JOIN roles r ON r.id = rp.role_id").
        Where("r.name IN ?", roleNames).
        Find(&permissions).Error

    if err != nil {
        return nil, err
    }

    var result []string
    for _, p := range permissions {
        result = append(result, p.Code)
    }

    return result, nil
}