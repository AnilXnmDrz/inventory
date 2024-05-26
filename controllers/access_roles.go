package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jacky-htg/inventory/libraries/api"
	"github.com/jacky-htg/inventory/models"
	"github.com/jacky-htg/inventory/payloads/response"
)

// RoleBasedAccess : struct for setting AccessRole Dependency Injection
type RoleBasedAccess struct {
	Db  *sql.DB
	Log *log.Logger
}

// List : http handler for returning list of roles with their access
func (ar *RoleBasedAccess) List(w http.ResponseWriter, r *http.Request) {
	var accessRole models.AccessRole
	tx, err := ar.Db.Begin()
	if err != nil {
		ar.Log.Printf("Begin tx : %+v", err)
		api.ResponseError(w, fmt.Errorf("getting roles with access list: %v", err))
		return
	}

	list, err := accessRole.ListRolesWithAccess(r.Context(), tx)
	if err != nil {
		tx.Rollback()
		ar.Log.Printf("ERROR : %+v", err)
		api.ResponseError(w, fmt.Errorf("getting roles with access list: %v", err))
		return
	}

	tx.Commit()

	var listResponse []*response.RoleWithAccessResponse
	for _, role := range list {
		var roleResponse response.RoleWithAccessResponse
		roleResponse.Transform(&role)
		listResponse = append(listResponse, &roleResponse)
	}

	api.ResponseOK(w, response.APIResponse{
		StatusCode:    "REBEL-200",
		StatusMessage: "OK",
		Data:          listResponse,
	}, http.StatusOK)
}
