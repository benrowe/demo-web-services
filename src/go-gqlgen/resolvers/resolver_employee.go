package resolvers

import (
	"context"
	"github.com/benrowe/demo-web-services/src/app/models"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/resolvers/transformations"
	"log"
)

func (m mutationResolver) CreateEmployee(ctx context.Context, input entities.CreateEmployeeInput) (*entities.Employee, error) {
	employee, err := transformations.ConvertCreateEmployeeInputToEmployeeModel(&input)
	if err != nil {
		log.Printf("unable to convert input into employee model: %v", err)
		return nil, err
	}

	m.App.DB.Save(&employee)

	return transformations.ModelToGqlEntityEmployee(employee)
}

func (m mutationResolver) TerminateEmployee(ctx context.Context, id string) (*entities.Employee, error) {
	panic("implement me")
	return nil, nil
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*entities.Employee, error) {
	model, err := models.FindEmployee(r.App.DB, id)
	if err != nil {
		return nil, err
	}
	return transformations.ModelToGqlEntityEmployee(model)
}

func (r *queryResolver) Employees(ctx context.Context) ([]*entities.Employee, error) {
	var models []models.Employee
	r.App.DB.Find(&models)

	var entities []*entities.Employee

	for _, model := range models {
		entity, err := transformations.ModelToGqlEntityEmployee(&model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}
