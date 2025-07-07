package tenant

import "time"

type Tenant struct {
	ID        string                 `json:"tenant_id" firestore:"tenant_id"`
	Name      string                 `json:"name" firestore:"name"`
	Alias     string                 `json:"alias" firestore:"alias"`
	Status    string                 `json:"status" firestore:"status"`
	Settings  map[string]interface{} `json:"settings" firestore:"settings"`
	CreatedAt time.Time              `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time              `json:"updated_at" firestore:"updated_at"`
}

type Context struct {
	ID       string `json:"tenant_id"`
	Name     string `json:"tenant_name"`
	Alias    string `json:"tenant_alias"`
	Status   string `json:"status"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type Config struct {
	FirestoreProjectID string
	JWTSecret         string
	DefaultTenant     string
}