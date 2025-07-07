package firestore

import (
	"context"
	"cloud.google.com/go/firestore"
	"github.com/azzidev/zensegur-provider-tenant/pkg/tenant"
)

type Repository struct {
	client *firestore.Client
}

func NewRepository(client *firestore.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) GetByID(ctx context.Context, tenantID string) (*tenant.Tenant, error) {
	doc, err := r.client.Collection("tenants").Doc(tenantID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var t tenant.Tenant
	if err := doc.DataTo(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *Repository) GetByAlias(ctx context.Context, alias string) (*tenant.Tenant, error) {
	iter := r.client.Collection("tenants").Where("alias", "==", alias).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}

	var t tenant.Tenant
	if err := doc.DataTo(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *Repository) Create(ctx context.Context, t *tenant.Tenant) error {
	_, err := r.client.Collection("tenants").Doc(t.ID).Set(ctx, t)
	return err
}

func (r *Repository) Update(ctx context.Context, t *tenant.Tenant) error {
	_, err := r.client.Collection("tenants").Doc(t.ID).Set(ctx, t)
	return err
}

func (r *Repository) List(ctx context.Context) ([]*tenant.Tenant, error) {
	iter := r.client.Collection("tenants").Documents(ctx)
	var tenants []*tenant.Tenant

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var t tenant.Tenant
		if err := doc.DataTo(&t); err != nil {
			continue
		}
		tenants = append(tenants, &t)
	}

	return tenants, nil
}