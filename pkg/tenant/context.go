package tenant

import (
	"context"
	"github.com/gin-gonic/gin"
)

const contextKey = "tenant_context"

// FromContext extracts tenant context from Gin context
func FromContext(c *gin.Context) *Context {
	if ctx, exists := c.Get(contextKey); exists {
		return ctx.(*Context)
	}
	return nil
}

// FromGoContext extracts tenant context from Go context
func FromGoContext(ctx context.Context) *Context {
	if tenantCtx, ok := ctx.Value(contextKey).(*Context); ok {
		return tenantCtx
	}
	return nil
}

// SetContext sets tenant context in Gin context
func SetContext(c *gin.Context, tenantCtx *Context) {
	c.Set(contextKey, tenantCtx)
}

// WithContext adds tenant context to Go context
func WithContext(ctx context.Context, tenantCtx *Context) context.Context {
	return context.WithValue(ctx, contextKey, tenantCtx)
}

// GetTenantID returns tenant ID from context
func GetTenantID(c *gin.Context) string {
	if ctx := FromContext(c); ctx != nil {
		return ctx.ID
	}
	return ""
}

// GetUserID returns user ID from context
func GetUserID(c *gin.Context) string {
	if ctx := FromContext(c); ctx != nil {
		return ctx.UserID
	}
	return ""
}