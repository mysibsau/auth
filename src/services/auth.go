package services

import (
	"fmt"
	"github.com/skilld-labs/go-odoo"
	"log"
	"src/shemas"
)

func newClient(database string, login string, password string) (client *odoo.Client, err error) {
	client, err = odoo.NewClient(&odoo.ClientConfig{
		Admin:    login,
		Password: password,
		Database: database,
		URL:      fmt.Sprintf("https://%s.pallada.sibsau.ru", database),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func Auth(auth shemas.Auth) (err error) {
	_, err = newClient("portfolio", auth.Login, auth.Password)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
