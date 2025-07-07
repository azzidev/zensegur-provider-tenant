package tenant

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/azzidev/zensegur-tenant/pkg/jwt"
)

type Middleware struct {
	config     *Config
	repository Repository
}

func NewMiddleware(config *Config, repo Repository) *Middleware {
	return &Middleware{
		config:     config,
		repository: repo,
	}
}

// AuthMiddleware validates JWT and sets tenant context
func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try cookie first
		token, err := c.Cookie("auth-token")
		if err != nil {
			// Fallback to Authorization header
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication provided"})
				c.Abort()
				return
			}
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// Validate JWT
		claims, err := jwt.ValidateToken(token, []byte(m.config.JWTSecret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Get tenant info
		tenantInfo, err := m.repository.GetByID(c.Request.Context(), claims.TenantID)
		if err != nil || tenantInfo.Status != "active" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid or inactive tenant"})
			c.Abort()
			return
		}

		// Set tenant context
		tenantCtx := &Context{
			ID:       tenantInfo.ID,
			Name:     tenantInfo.Name,
			Alias:    tenantInfo.Alias,
			Status:   tenantInfo.Status,
			UserID:   claims.UserID,
			Username: claims.Username,
		}

		SetContext(c, tenantCtx)
		c.Next()
	}
}

// RequireRole middleware for role-based access
func (m *Middleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// This would need to be implemented based on your role system
		// For now, just pass through
		c.Next()
	}
}