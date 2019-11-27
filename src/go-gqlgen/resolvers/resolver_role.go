package resolvers

import (
	"context"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
)

func (q queryResolver) Roles(ctx context.Context) ([]*entities.Role, error) {
	return nil, nil
}

func (q queryResolver) Role(ctx context.Context, id string) (*entities.Role, error) {
	return nil, nil
}

func (m mutationResolver) CreateRole(ctx context.Context, input entities.CreateRoleInput) (*entities.Role, error) {
	panic("implement me")
}

func (m mutationResolver) AssignRole(ctx context.Context, input *entities.AssignRoleInput) ([]*entities.AssignedRole, error) {
	panic("implement me")
}
