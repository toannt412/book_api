package admin

import (
	"bookstore/configs"
	"bookstore/dao/admin/model"
	"bookstore/helpers"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminsCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

func GetAdminByID(ctx context.Context, adminId string) (model.Admin, error) {

	var admin model.Admin

	objId, _ := primitive.ObjectIDFromHex(adminId)

	err := adminsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&admin)
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil

}

func LoginAccountAdmin(ctx context.Context, username, password string) (model.Admin, string, error) {
	var admin model.Admin
	var find bson.M
	err := adminsCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
	if err != nil {
		return model.Admin{},"", err
	}

	//Convert interface to string
	hashedPassword := fmt.Sprintf("%v", find["password"])
	err = helpers.CheckPasswordHash(hashedPassword, password)
	if err != nil {
		return model.Admin{}, "", err
	}

	token, errCreate := helpers.CreateJWT(username)
	if errCreate != nil {
		return model.Admin{}, "", err
	}

	return admin, token, nil

}

func GetAdminByUserName(ctx context.Context, username string) (model.Admin, error) {

	var admin model.Admin

	err := adminsCollection.FindOne(ctx, bson.M{"username": username}).Decode(&admin)
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil

}

// func GetAdminByEmail(ctx context.Context, email string) (model.Admin, error) {

// 	var admin model.Admin

// 	err := adminsCollection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
// 	if err != nil {
// 		return model.Admin{}, err
// 	}

// 	return admin, nil

// }
