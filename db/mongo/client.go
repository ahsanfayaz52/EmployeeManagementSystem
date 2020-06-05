package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/config"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/db"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

const (
	collectionName = "employee_info"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c client) SaveEmployee(employee *models.Employee) (string, error) {
	if employee.ID == "" {
		employee.ID = uuid.NewV4().String()
	}
	opts := options.Update().SetUpsert(true)
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"id": employee.ID}, bson.M{"$set": employee}, opts); err != nil {
		panic("oops cannot update employee")
	}
	return "", nil
}

func (c client) GetEmployeeByID(id string) (*models.Employee, error) {
	var employee *models.Employee
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&employee); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			panic("oops error occurred ")
		}
		return nil, err
	}
	return employee, nil
}

func (c client) DeleteEmployee(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"id": id}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			panic("oops error occurred ")
		}
	}
	return nil
}

func (c client) ListEmployee() (*[]models.Employee, error) {
	var employee []models.Employee
	var em models.Employee
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log().Info("oops no records fetched")
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&em); err != nil {
			log().Info("cannot add elements to list")
		}
		employee = append(employee, em)
	}
	return &employee, nil
}
