package scope

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 如果 tenant=0 就不需要增加tenantId筛选

func TenantScope(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		i := 1
		if i != 0 {
			return db.Where("tenant_id = ?", 0)
		} else {
			return db
		}
	}
}
