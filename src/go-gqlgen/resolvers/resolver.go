package resolvers

import (
    "context"
    "github.com/benrowe/demo-web-services/src/app"
    "github.com/benrowe/demo-web-services/src/app/models"
    "log"

    "github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
    App *app.App
}

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
    var e []models.Employee
    r.App.DB.Find(&e)

    log.Printf("%+v", e)

	employees := []*entities.Employee{
	    &entities.Employee{
	        ID: "asdfasdfasdf",
        },
    }

	return employees, nil
}
