package tenant

import "context"

// Repository interface for tenant data access
type Repository interface {
	GetByID(ctx context.Context, tenantID string) (*Tenant, error)
	GetByAlias(ctx context.Context, alias string) (*Tenant, error)
	Create(ctx context.Context, tenant *Tenant) error
	Update(ctx context.Context, tenant *Tenant) error
	List(ctx context.Context) ([]*Tenant, error)
}