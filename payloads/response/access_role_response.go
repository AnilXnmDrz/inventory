package response

import "github.com/jacky-htg/inventory/models"

// AccessRoleResponse is the response payload for an access entity
type AccessRoleResponse struct {
	AccessID   uint32 `json:"access_id"`
	AccessName string `json:"access_name"`
}

// RoleWithAccessResponse is the response payload for a role with its access entities
type RoleWithAccessResponse struct {
	ID       uint32          `json:"id"`
	RoleID   uint32          `json:"role_id"`
	RoleName string          `json:"role_name"`
	Access   []AccessRoleResponse `json:"access"`
}

// Transform transforms a RoleWithAccess model to a RoleWithAccessResponse
func (r *RoleWithAccessResponse) Transform(model *models.RoleWithAccess) {
	r.ID = model.RoleID
	r.RoleID = model.RoleID
	r.RoleName = model.RoleName
	for _, access := range model.Accesses {
		var accessResponse AccessRoleResponse
		accessResponse.AccessID = access.ID
		accessResponse.AccessName = access.Name
		r.Access = append(r.Access, accessResponse)
	}
}

// APIResponse is the standard response format
type APIResponse struct {
	StatusCode    string          `json:"status_code"`
	StatusMessage string          `json:"status_message"`
	Data          interface{}     `json:"data"`
}
