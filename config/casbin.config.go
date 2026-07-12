package config

import (
	"log"

	"github.com/casbin/casbin/v3"
)

var CasbinEnforcer *casbin.Enforcer

func configCasbinEnforcer() {
	var err error
	CasbinEnforcer, err = casbin.NewEnforcer("casbin/rbac_model.conf", "casbin/rbac_policy.csv")
	if err != nil {
		log.Printf("Failed to initialize Casbin enforcer: %v", err)
		return
	}

	if err = CasbinEnforcer.LoadPolicy(); err != nil {
		log.Printf("Failed to load Casbin policy: %v", err)
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
