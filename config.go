package tenantable

import (
  "gorm.io/gorm"
)

type Config struct {
  TenantKey        string
  DB               *gorm.DB
  AutoMigrate      bool
}
