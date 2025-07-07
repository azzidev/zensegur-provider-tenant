# Zensegur Provider Tenant

Provider para gerenciamento de multi-tenancy nos microservices da Zensegur.

## 🏗️ Estrutura

- **pkg/tenant** - Core tenant management
- **pkg/jwt** - JWT with tenant claims
- **pkg/firestore** - Firestore repository implementation
- **cmd** - Tenant API service

## 🚀 Como usar nos microservices

### 1. Import do framework
```go
import (
    "github.com/azzidev/zensegur-provider-tenant/pkg/tenant"
    "github.com/azzidev/zensegur-provider-tenant/pkg/firestore"
)
```

### 2. Setup do middleware
```go
// Initialize
repo := firestore.NewRepository(firestoreClient)
config := &tenant.Config{
    JWTSecret: os.Getenv("JWT_SECRET"),
}
middleware := tenant.NewMiddleware(config, repo)

// Use in routes
r.Use(middleware.AuthMiddleware())
```

### 3. Usar context nos handlers
```go
func getUsers(c *gin.Context) {
    tenantID := tenant.GetTenantID(c)
    userID := tenant.GetUserID(c)
    
    // Filter by tenant
    users := db.Where("tenant_id = ?", tenantID).Find(&users)
    c.JSON(200, users)
}
```

## 🔐 JWT Claims

```json
{
  "sub": "user_123",
  "tenant_id": "corretora_xyz", 
  "tenant_name": "Corretora XYZ",
  "username": "admin@corretora",
  "roles": ["admin"],
  "exp": 1234567890
}
```

## 🌐 API Endpoints

- `GET /health` - Health check
- `GET /api/tenant/:id` - Get tenant by ID (for auth service)
- `GET /api/tenant/alias/:alias` - Get tenant by alias (for auth service)

## 🚀 Deploy

Service roda em `api-tenant.zensegur.com.br`

## 📦 Environment Variables

- `GOOGLE_CLOUD_PROJECT` - Firebase project ID
- `JWT_SECRET` - JWT signing secret
- `PORT` - Server port (default: 8080)