package middleware

import (
	"github.com/gin-gonic/gin"
)

type contextKey string

const TenantIDKey contextKey = "tenantID"

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//tenantIDStr := c.GetHeader("X-Tenant-ID")
		//tenantID, err := strconv.Atoi(tenantIDStr)
		//if err != nil || tenantID <= 0 {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		//	c.Abort()
		//	return
		//}
		//
		//ctx := context.WithValue(c.Request.Context(), TenantIDKey, uint(tenantID))
		//c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
