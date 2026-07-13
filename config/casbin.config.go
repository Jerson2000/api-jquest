package config

import (
	"log"

	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	stringadapter "github.com/casbin/casbin/v3/persist/string-adapter"
	casbindata "github.com/jerson2000/jquest/casbin"
)

var CasbinEnforcer *casbin.Enforcer

func configCasbinEnforcer() {
	var err error
	m, err := model.NewModelFromString(casbindata.RbacModel)
	if err != nil {
		log.Printf("Failed to initialize Casbin enforcer: %v", err)
		return
	}

	sa := stringadapter.NewAdapter(casbindata.RbacPolicy)

	CasbinEnforcer, err = casbin.NewEnforcer(m, sa)
	if err != nil {
		log.Printf("Failed to initialize Casbin enforcer: %v", err)
		return
	}

	log.Println("Casbin enforcer initialized and policy loaded successfully")
}

// package config

// import (
// 	"log"

// 	"github.com/casbin/casbin/v3"
// 	gormadapter "github.com/casbin/gorm-adapter/v3"
// )

// var CasbinEnforcer *casbin.Enforcer

// func configCasbinEnforcer() {
// 	var err error
// 	adapter, err := gormadapter.NewAdapterByDB(Database)
// 	if err != nil {
// 		log.Printf("Failed to initialize Casbin enforcer: %v", err)
// 		return
// 	}
// 	CasbinEnforcer, err = casbin.NewEnforcer("casbin/rbac_model.conf", adapter)
// 	if err != nil {
// 		log.Printf("Failed to initialize Casbin enforcer: %v", err)
// 		return
// 	}

// 	if err = CasbinEnforcer.LoadPolicy(); err != nil {
// 		log.Printf("Failed to load Casbin policy: %v", err)
// 		return
// 	}

// 	CasbinEnforcer.AddPolicy("admin", "/api/users", "GET")
// 	CasbinEnforcer.AddPolicy("admin", "/api/users", "POST")
// 	CasbinEnforcer.AddPolicy("admin", "/api/companies", "POST")
// 	CasbinEnforcer.AddPolicy("admin", "/api/companies", "GET")
// 	CasbinEnforcer.AddPolicy("candidate", "/api/current", "GET")

// 	log.Println("Casbin enforcer initialized and policy loaded successfully")
// }
