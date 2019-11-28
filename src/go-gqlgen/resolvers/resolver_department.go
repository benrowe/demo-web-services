package resolvers

import (
	"context"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
)

func (q queryResolver) Departments(ctx context.Context) ([]*entities.Department, error) {
	panic("implement me")
	return nil, nil
}

func (q queryResolver) Department(ctx context.Context, id string) (*entities.Department, error) {
	panic("implement me")
	return nil, nil
}

func (m mutationResolver) CreateDepartment(ctx context.Context, input entities.CreateDepartmentInput) (*entities.Department, error) {
	panic("implement me")
}
