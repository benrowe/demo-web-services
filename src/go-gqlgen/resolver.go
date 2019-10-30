package main

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Employee(ctx context.Context, id string) (*Employee, error) {
	panic("not implemented")
}
func (r *queryResolver) Employees(ctx context.Context) ([]*Employee, error) {
	panic("not implemented")
}
