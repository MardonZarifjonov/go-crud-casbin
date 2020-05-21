package controllers

import (
	"fmt"
	"net/http"

	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/responses"
)

var (
	casbins = CasbinModel{}
)

// CreateRoleCasbin func to create a role
func CreateRoleCasbin(w http.ResponseWriter, r *http.Request) {
	rolename := r.FormValue("rolename")
	path := r.FormValue("path")
	method := r.FormValue("method")
	ptype := "p"

	casbinObj := CasbinModel{
		Ptype:    ptype,
		RoleName: rolename,
		Path:     path,
		Method:   method,
	}
	isok := casbins.AddCasbin(casbinObj)
	if isok {
		fmt.Println("Role created")
		responses.JSON(w, http.StatusCreated, nil)
	} else {
		fmt.Println("Error during creation of role")
		responses.ERROR(w, http.StatusInternalServerError, nil)
	}
}
