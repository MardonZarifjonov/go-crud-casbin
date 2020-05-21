package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/auth"
	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/models"
	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/responses"
	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

// Login func to  perform login
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formatedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formatedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

// SignIn fucn
func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID, server.DB)
}
