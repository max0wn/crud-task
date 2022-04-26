package repositories

import (
	"context"
	"crud-task/internal/entities"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "toor"
)

const DATABASE = "app"
const COLLECTION = "users"

type UserRepository struct {
	client *mongo.Client
}

func (repository *UserRepository) Create(user *entities.UserEntity) error {
	collection := repository.client.Database(DATABASE).Collection(COLLECTION)

	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (repository *UserRepository) Update(user *entities.UserEntity) error {
	collection := repository.client.Database(DATABASE).Collection(COLLECTION)

	filter := bson.D{{"id", user.Id}}
	update := bson.D{{
		"$set", bson.D{
			{"name", user.Name},
			{"surname", user.Surname},
		},
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (repository *UserRepository) Delete(user *entities.UserEntity) error {
	collection := repository.client.Database(DATABASE).Collection(COLLECTION)

	_, err := collection.DeleteOne(context.TODO(), user)

	if err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func NewUserRepository() *UserRepository {
	uri := fmt.Sprint("mongodb://mongodb")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username: USERNAME,
		Password: PASSWORD,
	}

	options := options.Client().ApplyURI(uri).SetAuth(credentials)
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return &UserRepository{
		client: client,
	}
}
