package tenant

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
	"github.com/yehlo/storagegrid-sdk-go/services/tenantusage"
)

const (
	tenantEndpoint = "/grid/accounts"
)

// ServiceInterface defines the contract for tenant service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]Tenant, error)
	GetByID(ctx context.Context, id string) (*Tenant, error)
	Create(ctx context.Context, tenant *Tenant) (*Tenant, error)
	Update(ctx context.Context, tenant *Tenant) (*Tenant, error)
	Delete(ctx context.Context, id string) error
	GetUsage(ctx context.Context, id string) (*tenantusage.TenantUsage, error)
}

type Service struct {
	services.HTTPClient
}

// NewService returns a new tenant service using the provided client
func NewService(client services.HTTPClient) *Service {
	return &Service{client}
}

func (s *Service) List(ctx context.Context) ([]Tenant, error) {
	var response models.Response[[]Tenant]
	err := s.DoParsed(ctx, "GET", tenantEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Tenant, error) {
	var response models.Response[*Tenant]
	err := s.DoParsed(ctx, "GET", tenantEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Create(ctx context.Context, tenant *Tenant) (*Tenant, error) {
	var response models.Response[*Tenant]
	err := s.DoParsed(ctx, "POST", tenantEndpoint, tenant, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Update(ctx context.Context, tenant *Tenant) (*Tenant, error) {
	var response models.Response[*Tenant]
	err := s.DoParsed(ctx, "PUT", tenantEndpoint+"/"+tenant.ID, tenant, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.DoParsed(ctx, "DELETE", tenantEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUsage(ctx context.Context, id string) (*tenantusage.TenantUsage, error) {
	var response models.Response[*tenantusage.TenantUsage]

	err := s.DoParsed(ctx, "GET", tenantEndpoint+"/"+id+"/usage", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
