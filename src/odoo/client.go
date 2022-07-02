package odoo

import (
	"fmt"
	"github.com/skilld-labs/go-odoo"
	"log"
)

type Client struct {
	client *odoo.Client
}

func NewClient(database string, login string, password string) (*Client, error) {
	client, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    login,
		Password: password,
		Database: database,
		URL:      fmt.Sprintf("https://%s.pallada.sibsau.ru", database),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Client{client: client}, nil
}
