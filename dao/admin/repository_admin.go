package admin

import (
	"bookstore/auth"
	"bookstore/dao"
	"bookstore/dao/admin/model"
	"bookstore/helpers"
	"bookstore/serialize"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Client instance

type AdminRepository struct {
	adminsCollection *mongo.Collection
}

func NewAdminRepository() *AdminRepository {
	var DB *mongo.Client = dao.ConnectDB()

	return &AdminRepository{
		adminsCollection: dao.GetCollection(DB, "admins"),
	}
}

func (repo *AdminRepository) GetAdminByID(ctx context.Context, adminId string) (model.Admin, error) {

	var admin model.Admin

	objId, _ := primitive.ObjectIDFromHex(adminId)

	err := repo.adminsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&admin)
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil

}

func (repo *AdminRepository) LoginAccountAdmin(ctx context.Context, username, password string) (model.Admin, string, error) {
	var admin model.Admin
	var find bson.M
	err := repo.adminsCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
	if err != nil {
		return model.Admin{}, "", err
	}

	//Convert interface to string
	hashedPassword := fmt.Sprintf("%v", find["password"])
	err = helpers.CheckPasswordHash(hashedPassword, password)
	if err != nil {
		return model.Admin{}, "", err
	}

	token, errCreate := auth.GenerateJWT(admin.Email, admin.UserName)
	if errCreate != nil {
		return model.Admin{}, "", err
	}
	adminID := find["_id"]
	_, errAddToken := repo.adminsCollection.UpdateOne(context.TODO(), bson.M{"_id": adminID}, bson.M{"$set": bson.M{"token": token}})
	if errAddToken != nil {
		return model.Admin{}, "", errAddToken
	}

	return admin, token, nil

}

func (repo *AdminRepository) GetAdminByUserName(ctx context.Context, username string) (model.Admin, error) {

	var admin model.Admin

	err := repo.adminsCollection.FindOne(ctx, bson.M{"username": username}).Decode(&admin)
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil

}

func (repo *AdminRepository) GetAdminByEmail(ctx context.Context, email string) (model.Admin, error) {

	var admin model.Admin

	err := repo.adminsCollection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil

}

func (repo *AdminRepository) EditAdmin(ctx context.Context, id string, admin *serialize.Admin) (model.Admin, error) {
	//var admin model.Admin
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Admin{}, err
	}

	// Validate the request body

	update := model.Admin{
		Id:       objId,
		FullName: admin.FullName,
		Phone:    admin.Phone,
		Role:     admin.Role,
	}

	result, err := repo.adminsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return model.Admin{}, err
	}

	var updatedAdmin model.Admin
	if result.MatchedCount == 1 {
		err := repo.adminsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAdmin)
		if err != nil {
			return model.Admin{}, err
		}
	}

	return model.Admin{
		FullName: updatedAdmin.FullName,
		Phone:    updatedAdmin.Phone,
		Role:     updatedAdmin.Role,
	}, nil
}

func (repo *AdminRepository) DeleteAdmin(ctx context.Context, id string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "Deleted fail", err
	}

	result, err := repo.adminsCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return "Deleted fail", err
	}

	if result.DeletedCount == 0 {
		return "Deleted fail", mongo.ErrNoDocuments
	}

	return "Deleted successfully", nil
}

func (repo *AdminRepository) CreateAdmin(ctx context.Context, admin model.Admin) (model.Admin, error) {
	admin.Password, _ = helpers.Hash(admin.Password)

	_, err := repo.adminsCollection.InsertOne(ctx, admin)
	if err != nil {
		return model.Admin{}, err
	}
	return model.Admin{
		Id:       admin.Id,
		FullName: admin.FullName,
		Phone:    admin.Phone,
		Role:     admin.Role,
		UserName: admin.UserName,
		Password: admin.Password,
		Email:    admin.Email,
	}, nil
}

func (repo *AdminRepository) GetAllAdmins(ctx context.Context) ([]model.Admin, error) {

	var admins []model.Admin
	cursor, err := repo.adminsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var admin model.Admin
		if err := cursor.Decode(&admin); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}

func (repo *AdminRepository) GetAdminToken(ctx context.Context, token string) (model.Admin, error) {
	var admin model.Admin
	err := repo.adminsCollection.FindOne(ctx, bson.M{"token": token}).Decode(&admin)
	if err != nil {
		return model.Admin{}, err
	}
	return admin, nil
}

func (repo *AdminRepository) EditAminToken(ctx context.Context, token string) error {
	_, err := repo.adminsCollection.UpdateOne(ctx, bson.M{"token": token}, bson.M{"$unset": bson.M{"token": ""}})
	if err != nil {
		return err
	}
	return nil
}
