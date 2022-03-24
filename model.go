package tenantable

import (
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name      string
	UUID      string
	SubDomain string
}
