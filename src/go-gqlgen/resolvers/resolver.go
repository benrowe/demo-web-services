package resolvers

import (
	"github.com/benrowe/demo-web-services/src/app"
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

func (r *Resolver) Employee() gen.EmployeeResolver {
	return &employeeResolver{}
}
