package transformations

import (
	"github.com/benrowe/demo-web-services/src/app/models"
	"github.com/benrowe/demo-web-services/src/go-gqlgen/entities"
	"log"
	"time"
)

func ConvertCreateEmployeeInputToEmployeeModel(i *entities.CreateEmployeeInput) (*models.Employee, error) {
	log.Printf("converting CreateEmployeeInput to Employee model: %+v", i)
	t, err := time.Parse("2006/01/02", i.DateOfBirth)
	if err != nil {
		return nil, err
	}
	m := models.Employee{
		FirstName: i.FirstName,
		LastName:  i.LastName,
		BirthDate: &t,
		//Gender:    Gender,
		StartDate: nil,
		Salaries:  nil,
	}
	return &m, nil
}

func ModelToGqlEntityEmployee(m *models.Employee) (*entities.Employee, error) {
	return &entities.Employee{
		ID:          uint32ToStr(m.ID),
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		DateOfBirth: m.BirthDate.String(),
		Gender:      entities.GenderFemale,
		Age:         m.Age(),
	}, nil
}
