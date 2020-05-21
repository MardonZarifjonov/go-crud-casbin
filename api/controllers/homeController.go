package controllers

import (
	"net/http"

	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/responses"
)

// Home func
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome API")
}
