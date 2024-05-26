package models

import (
	"context"
	"database/sql"
)

// AccessRole represents the access_role entity
type AccessRole struct {
	ID       uint32
	AccessID uint32
	RoleID   uint32
}

// RoleWithAccess represents a role with its associated accesses
type RoleWithAccess struct {
	RoleID   uint32
	RoleName string
	Accesses []Access
}

const qRoleWithAccess = `
	SELECT r.id as role_id, r.name as role_name, a.id as access_id, a.alias as access_name
	FROM access_roles ar
	JOIN roles r ON ar.role_id = r.id
	JOIN access a ON ar.access_id = a.id
`

// ListRolesWithAccess fetches the list of roles and their associated accesses from the database
func (ar *AccessRole) ListRolesWithAccess(ctx context.Context, tx *sql.Tx) ([]RoleWithAccess, error) {
	rows, err := tx.QueryContext(ctx, qRoleWithAccess)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roleMap := make(map[uint32]*RoleWithAccess)
	for rows.Next() {
		var roleID uint32
		var roleName string
		var accessID uint32
		var accessName string

		if err := rows.Scan(&roleID, &roleName, &accessID, &accessName); err != nil {
			return nil, err
		}

		if role, exists := roleMap[roleID]; exists {
			role.Accesses = append(role.Accesses, Access{ID: accessID, Name: accessName})
		} else {
			roleMap[roleID] = &RoleWithAccess{
				RoleID:   roleID,
				RoleName: roleName,
				Accesses: []Access{{ID: accessID, Name: accessName}},
			}
		}
	}

	var roles []RoleWithAccess
	for _, role := range roleMap {
		roles = append(roles, *role)
	}

	return roles, rows.Err()
}
