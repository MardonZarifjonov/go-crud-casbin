package controllers

import (
	"fmt"
	"net/http"

	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/auth"
	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/responses"
	"github.com/pkg/errors"
)

// SetMiddlewareJSON func
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("MiddleJSON")
		next(w, r)
	}
}

// SetMiddlewareAuthentificaton func
func SetMiddlewareAuthentificaton(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

// AuthCheckRole func to check the role
func AuthCheckRole(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := auth.ExtractTokenRole(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		enforcer := Casbin()
		fmt.Printf("%s, %s, %s\n", role, r.URL.Path, r.Method)
		res, err := enforcer.EnforceSafe(role, r.URL.Path, r.Method)

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		if res {
			next(w, r)
		} else {
			fmt.Println("No error")
			responses.ERROR(w, http.StatusForbidden, errors.New("Permission denied"))
			return
		}
	}
}
