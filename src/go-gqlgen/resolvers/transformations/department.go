package transformations

import (
	"github.com/benrowe/demo-web-services/src/app/models"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
)

func ModelToGqlEntityDepartment(d *models.Department) (*entities.Department, error) {
	manager, err := ModelToGqlEntityEmployee(&d.Manager)

	return &entities.Department{
		ID:      uint32ToStr(d.ID),
		Name:    d.Name,
		Manager: manager,
	}, err
}
