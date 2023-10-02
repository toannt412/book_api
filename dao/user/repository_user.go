package user

import (
	"bookstore/configs"
	"bookstore/dao/user/model"
	"bookstore/helpers"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetUserByUserName(ctx context.Context, username string) (model.User, error) {
	var user model.User

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func RegisterAccount(ctx context.Context, username, password, email string) (string, error) {
	hashedPassword, err := helpers.Hash(password)
	if err != nil {
		return "", err
	}

	_, error := userCollection.InsertOne(ctx, bson.M{
		"username": username,
		"password": hashedPassword,
		"email":    email,
	})
	if error != nil {
		return "", err
	}
	return "Register Success", nil

}

func LoginAccount(ctx context.Context, username, password string) (model.User, string, error) {
	var user model.User
	var find bson.M
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
	if err != nil {
		return model.User{}, "", err
	}

	//Convert interface to string
	hashedPassword := fmt.Sprintf("%v", find["password"])
	err = helpers.CheckPasswordHash(hashedPassword, password)
	if err != nil {
		return model.User{}, "", err
	}

	token, errCreate := helpers.CreateJWT(username)
	if errCreate != nil {
		return model.User{}, "", err
	}

	return user, token, nil

}
