package tenantable

import (
	"gorm.io/gorm"
)

const (
	TenantByDomain = iota
	TenantByKey
)

type Config struct {
	TenantKeyMethod  uint
	TenantDomainBase string
	TenantKey        string
	DB               *gorm.DB
	AutoMigrate      bool
}
