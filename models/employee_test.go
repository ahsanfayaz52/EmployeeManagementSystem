package models

import (
	"reflect"
	"testing"
)

func TestEmployee_Map(t *testing.T) {
	type fields struct {
		ID      string
		Name    string
		Address string
		Age     int
		Salary  int
		Phone   int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - employee struct to map",
			fields: fields{
				ID:      "12345",
				Name:    "ahsan",
				Address: "attock",
				Age:     22,
				Salary:  20000,
				Phone:   12345,
			},
			want: map[string]interface{}{
				"id":      "12345",
				"name":    "ahsan",
				"address": "attock",
				"age":     22,
				"salary":  20000,
				"phone":   12345,
			},
		},
		{
			name: " success - employee struct to map with fewer fields",
			fields: fields{
				ID:     "12345",
				Name:   "ahsan",
				Salary: 20000,
			},
			want: map[string]interface{}{
				"id":      "12345",
				"name":    "ahsan",
				"address": "",
				"age":     0,
				"salary":  20000,
				"phone":   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
				Age:     tt.fields.Age,
				Salary:  tt.fields.Salary,
				Phone:   tt.fields.Phone,
			}
			if got := e.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_Names(t *testing.T) {
	type fields struct {
		ID      string
		Name    string
		Address string
		Age     int
		Salary  int
		Phone   int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:      "12345",
				Name:    "ahsan",
				Address: "attock",
				Age:     22,
				Salary:  20000,
				Phone:   12345,
			},
			want: []string{"id", "name", "address", "age", "salary", "phone"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
				Age:     tt.fields.Age,
				Salary:  tt.fields.Salary,
				Phone:   tt.fields.Phone,
			}
			if got := e.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
