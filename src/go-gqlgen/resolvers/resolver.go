package resolvers

import (
	"context"

    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Employee(ctx context.Context, id string) (*entities.Employee, error) {
	return &entities.Employee{
	    ID: id,
    }, nil
}

func (r *queryResolver) Employees(ctx context.Context) ([]*entities.Employee, error) {
	employees := []*entities.Employee{
	    &entities.Employee{
	        ID: "asdfasdfasdf",
        },
    }

	return employees, nil
}
