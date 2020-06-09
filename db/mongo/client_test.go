package mongo

import (
	"os"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/db"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

func Test_client_DeleteEmployee(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "employee-system-mongo-db")

	c, _ := NewClient(db.Option{})
	employee := &models.Employee{Name: "uzair", Address: "islamabad", Age: 30, Salary: 60000, Phone: 123}
	_, _ = c.AddEmployee(employee)
	type fields struct {
		conn *mongo.Client
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
			c, _ := NewClient(db.Option{})
			if err := c.DeleteEmployee(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetEmployeeByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "employee-system-mongo-db")

	c, _ := NewClient(db.Option{})
	employee := &models.Employee{Name: "usman", Address: "attock", Age: 25, Salary: 40000, Phone: 123}
	_, _ = c.AddEmployee(employee)
	type fields struct {
		conn *mongo.Client
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

func Test_client_ListEmployees(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "employee-system-mongo-db")

	c, _ := NewClient(db.Option{})
	employee := &models.Employee{Name: "usman", Address: "attock", Age: 25, Salary: 40000, Phone: 123}
	_, _ = c.AddEmployee(employee)
	type fields struct {
		conn *mongo.Client
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
			c, _ := NewClient(db.Option{})
			got, err := c.ListEmployees()
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

func Test_client_AddEmployee(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "employee-system-mongo-db")

	type fields struct {
		conn *mongo.Client
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
			name:    "success - add employee in db",
			args:    args{employee: &models.Employee{Name: "beenish", Address: "lahore", Age: 35, Salary: 50000, Phone: 456}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddEmployee(tt.args.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_UpdateEmployee(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "employee-system-mongo-db")

	type fields struct {
		conn *mongo.Client
	}
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success - update employee in db",
			args:    args{employee: &models.Employee{Name: "shah", Address: "lahore", Age: 35, Salary: 50000, Phone: 456}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			if err := c.UpdateEmployee(tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
