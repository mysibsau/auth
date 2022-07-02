package odoo_shemas

import (
	"log"
	"strconv"
	"strings"
)

type Gradle struct {
	FIO   string `xmlrpc:"ID_student"`
	Grade string `xmlrpc:"grade"`
	Group string `xmlrpc:"group"`
}

type UserWithAverage struct {
	Name    string  `xmlrpc:"name"`
	Average float64 `xmlrpc:"average"`
}

type Grades []Gradle

func (g *Grades) Average() float64 {
	var sum float64
	count := 0
	for _, grade := range *g {
		mark, err := strconv.ParseFloat(strings.TrimSpace(grade.Grade), 64)
		if err != nil {
			sum += 0
			continue
		}
		sum += mark
		count += 1
	}
	log.Printf("count: %d", count)
	return sum / float64(count)
}

func (g *Grades) GetName() string {
	if len(*g) == 0 {
		return ""
	}
	return (*g)[0].FIO
}

func (g *Grades) GetGroup() string {
	if len(*g) == 0 {
		return ""
	}
	return (*g)[0].Group
}
