package mysql

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/db"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"

)

func Test_client_DeleteEmployee(t *testing.T) {
	c, _ := NewClient(db.Option{})
	employee := &models.Employee{Name: "talha", Address: "islamabad", Age: 30, Salary: 60000, Phone: 123}
	_, _ = c.SaveEmployee(employee)
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete employee from db",
			args:    args{id: employee.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteEmployee(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetEmployeeByID(t *testing.T) {
	c, _ := NewClient(db.Option{})
	employee := &models.Employee{Name: "usman", Address: "attock", Age: 25, Salary: 40000, Phone: 123}
	_, _ = c.SaveEmployee(employee)
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Employee
		wantErr bool
	}{
		{
			name:    "success - get employee from db",
			args:    args{id: employee.ID},
			want:    employee,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			got, err := c.GetEmployeeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployeeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_ListEmployee(t *testing.T) {
	c, _ := NewClient(db.Option{})
	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]models.Employee
		wantErr bool
	}{
		{
			name:    "success - list employees from db",
			want:    &[]models.Employee{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListEmployee()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				log().Info(got)
			}
		})
	}
}

func Test_client_SaveEmployee(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success - ADD employee in db",
			args:    args{employee: &models.Employee{Name: "beenish", Address: "lahore", Age: 23, Salary: 20000, Phone: 123}},
			wantErr: false,
		},
		{
			name:    "success - UPDATE employee in db",
			args:    args{employee: &models.Employee{ID: "adebef41-8a90-4f73-80cc-d122a2dadb54", Name: "osama khan", Address: "sakkar", Age: 23, Salary: 20000, Phone: 123}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			got, err := c.SaveEmployee(tt.args.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveEmployee() got = %v, want %v", got, tt.want)
			}
		})
	}
}