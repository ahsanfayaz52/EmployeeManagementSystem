package models

import (
	"github.com/fatih/structs"
)

// Employee holds information of employee.
type Employee struct {
	ID      string `json:"id" structs:"id"  bson:"id" db:"id"`
	Name    string `json:"name" structs:"name"  bson:"name" db:"name"`
	Address string `json:"address" structs:"address"  bson:"address" db:"address"`
	Age     int    `json:"age" structs:"age" bson:"age" db:"age"`
	Salary  int    `json:"salary" structs:"salary" bson:"salary" db:"salary"`
	Phone   int    `json:"phone" structs:"phone" bson:"phone" db:"phone"`
}

// Map converts structs to a map representation.
func (e *Employee) Map() map[string]interface{} {
	return structs.Map(e)
}

// Names returns the field names of Employee model.
func (e *Employee) Names() []string {
	fields := structs.Fields(e)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
