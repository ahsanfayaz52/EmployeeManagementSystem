package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/config"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/db"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

// EmployeeTableName contains name of table.
const (
	EmployeeTableName = "employee_info"
)

func init() {
	db.Register("mysql", NewClient)
}

//The first implementation.
type client struct {
	db *sqlx.DB
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)

	return cfg.FormatDSN()
}

// NewClient initializes a mysql database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	log().Info("data = ", cli.Ping())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{db: cli}, nil
}

func (c client) SaveEmployee(employee *models.Employee) (string, error) {
	if employee.ID == "" {
		employee.ID = uuid.NewV4().String()
	}
	names := employee.Names()
	log().Info("names = ", names)
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT IGNORE INTO %s (%s) VALUES(%s)`, EmployeeTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), employee); err != nil {
		return "", errors.Wrap(err, "failed to add/update employee")
	}
	return employee.ID, nil
}

func (c client) GetEmployeeByID(id string) (*models.Employee, error) {
	var em models.Employee
	if err := c.db.Get(&em, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, EmployeeTableName, id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "failed to get employee")
		}
	}
	return &em, nil
}

func (c client) DeleteEmployee(id string) (string, error) {
	rows, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, EmployeeTableName, id))
	if err != nil {
		if errors.Is(err, rows.Err()) {
			return "", errors.Wrap(err, "failed to delete employee")
		}
	}
	return "employee is removed successfully", nil
}

<<<<<<< HEAD
func (c client) ListEmployee() ([]*models.Employee, error) {
=======
func (c client) ListEmployees() ([]*models.Employee, error) {
>>>>>>> feature/service
	var emList []*models.Employee
	var em models.Employee
	rows, err := c.db.Query(fmt.Sprintf(`SELECT * FROM %s`, EmployeeTableName))
	if err != nil {
		if errors.Is(err, rows.Err()) {
			return nil, errors.Wrap(err, "failed to retrieve employees")
		}
	}
	for rows.Next() {
		err = rows.Scan(&em.ID, &em.Name, &em.Address, &em.Age, &em.Salary, &em.Phone)
		if err != nil {
<<<<<<< HEAD
			return nil, errors.Wrap(err, "oops error occurred while scanning.....")
=======
			return nil, errors.Wrap(err, "Error occurred while scanning rows")
>>>>>>> feature/service
		}
		emList = append(emList, &em)
	}
	return emList, nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}
	return ph
}
