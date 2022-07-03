package odoo

import (
	"github.com/google/uuid"
	"log"
	"src/models"
	odoo_shemas "src/odoo/shemas"

	"src/shemas"
)

func (c *Client) GetGrades() (grades *odoo_shemas.Grades, err error) {
	grades = &odoo_shemas.Grades{}
	err = c.client.SearchRead(
		"portfolio_science.grade_view",
		nil,
		nil,
		grades,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func Auth(auth shemas.Auth) (*models.User, error) {
	client, err := NewClient("portfolio", auth.Login, auth.Password)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	grades, err := client.GetGrades()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if len(*grades) == 0 {
		log.Fatal("Оценок не найдено")
	}
	return &models.User{
		Id:      uuid.NewString(),
		Group:   grades.GetGroup(),
		Name:    grades.GetName(),
		Login:   auth.Login,
		Average: grades.Average(),
	}, nil
}
