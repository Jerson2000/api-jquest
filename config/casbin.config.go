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
