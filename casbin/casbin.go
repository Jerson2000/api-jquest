package casbin

import (
	_ "embed"
)

//go:embed rbac_model.conf
var RbacModel string

//go:embed rbac_policy.csv
var RbacPolicy string
