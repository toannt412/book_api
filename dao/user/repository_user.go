package user

import (
	"bookstore/configs"
	"bookstore/dao/user/model"
	"bookstore/helpers"
	"bookstore/serialize"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func DeleteUser(ctx context.Context, id string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "Deleted fail", err
	}

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return "Deleted fail", err
	}

	if result.DeletedCount == 0 {
		return "Deleted fail", mongo.ErrNoDocuments
	}

	return "Deleted successfully", nil
}

func EditUser(ctx context.Context, id string, sth *serialize.User) (model.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, err
	}

	// Validate the request body

	update := model.User{
		Id:          objId,
		FullName:    sth.FullName,
		Location:    sth.Location,
		DateOfBirth: sth.DateOfBirth,
		Phone:       sth.Phone,
	}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return model.User{}, err
	}

	var updatedUser model.User
	if result.MatchedCount == 1 {
		err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
		if err != nil {
			return model.User{}, err
		}
	}

	return model.User{
		Id:          updatedUser.Id,
		FullName:    updatedUser.FullName,
		Location:    updatedUser.Location,
		DateOfBirth: updatedUser.DateOfBirth,
		Phone:       updatedUser.Phone,
	}, nil
}
func GetUserByID(ctx context.Context, userId string) (model.User, error) {

	var user model.User

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil

}
func GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
