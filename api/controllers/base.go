package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server struct to store data about server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// GlobalGormDB var to store gorm.DB pointer
var (
	GlobalGormDB *gorm.DB
)

// Initialize func to get database initialized
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", Dbdriver)
		GlobalGormDB = server.DB
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// GetGormDbPointer func to get gorm db
func GetGormDbPointer() *gorm.DB {
	return GlobalGormDB
}

// Run func to run app
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
