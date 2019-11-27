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
	return nil, nil
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*entities.Employee, error) {
	return &entities.Employee{
		ID: id,
	}, nil
}

func (r *queryResolver) Employees(ctx context.Context) ([]*entities.Employee, error) {
	var e []models.Employee
	r.App.DB.Find(&e)

	var ee []*entities.Employee

	for _, em := range e {
		eaaa, err := transformations.ModelToGqlEntityEmployee(&em)
		if err != nil {
			return nil, err
		}
		ee = append(ee, eaaa)
	}

	return ee, nil
}
