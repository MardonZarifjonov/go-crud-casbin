package controllers

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

type CasbinModel struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Ptype    string `gorm:"size:100;not null" json:"ptype"`
	RoleName string `gorm:"size:100;not null" json:"rolename"`
	Path     string `gorm:"size:100;not null" json:"path"`
	Method   string `gorm:"size:100;not null" json:"method"`
}

func (c *CasbinModel) AddCasbin(cm CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

func Casbin() *casbin.Enforcer {
	adapter := gormadapter.NewAdapterByDB(GetGormDbPointer())
	enforcer := casbin.NewEnforcer("conf/rbac_model.conf", adapter)
	enforcer.LoadPolicy()
	return enforcer
}

// var (
// 	enforcer *casbin.Enforcer
// )

// // CasInit func to initialize casbin
// func CasInit(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
// 	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
// 	adapterConnect := gormadapter.NewAdapterByDB(GetGormDbPointer())
// 	enforcer = casbin.NewEnforcer("conf/rbac_model.conf", adapterConnect)
// 	fmt.Println("Initialized casbin")
// 	enforcer.LoadPolicy()
// }

// // GetEnforcerPointer func to get casbin enforcer
// func GetEnforcerPointer() *casbin.Enforcer {
// 	return enforcer
// }
