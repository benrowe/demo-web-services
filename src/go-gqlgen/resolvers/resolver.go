package resolvers

import (
    "context"
    "github.com/benrowe/demo-web-services/src/app"
    "github.com/benrowe/demo-web-services/src/app/models"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/resolvers/transformations"
    log "github.com/sirupsen/logrus"

    "github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
    App *app.App
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *Resolver) Query() gen.QueryResolver {
    return &queryResolver{r}
}

func (r *Resolver) Mutation() gen.MutationResolver {
    return &mutationResolver{r}
}

func (m mutationResolver) CreateEmployee(ctx context.Context, input entities.CreateEmployeeInput) (*entities.Employee, error) {
    employee, err := transformations.ConvertCreateEmployeeInputToEmployeeModel(&input)
    if err != nil {
        log.Printf("unable to convert input into employee model: %v", err)
        return nil, err
    }

    m.App.DB.Save(&employee)

    return transformations.ModelToGqlEntityEmployee(employee)
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
