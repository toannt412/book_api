package user

import (
	"bookstore/auth"
	"bookstore/dao"
	"bookstore/dao/user/model"
	"bookstore/helpers"
	"bookstore/serialize"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	usersCollection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	var DB *mongo.Client = dao.ConnectDB()
	return &UserRepository{
		usersCollection: dao.GetCollection(DB, "users"),
	}
}

func (repo *UserRepository) GetUserByUserName(ctx context.Context, username string) (model.User, error) {
	var user model.User

	err := repo.usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := repo.usersCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) RegisterAccount(ctx context.Context, username, password, email, phone string) (string, error) {
	hashedPassword, err := helpers.Hash(password)
	if err != nil {
		return "", err
	}

	_, error := repo.usersCollection.InsertOne(ctx, bson.M{
		"username": username,
		"password": hashedPassword,
		"email":    email,
		"phone":    phone,
	})
	if error != nil {
		return "", err
	}
	return "Register Success", nil

}

func (repo *UserRepository) LoginAccount(ctx context.Context, username, password string) (model.User, string, error) {
	var user model.User
	var find bson.M
	err := repo.usersCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
	if err != nil {
		return model.User{}, "", err
	}

	//Convert interface to string
	hashedPassword := fmt.Sprintf("%v", find["password"])
	err = helpers.CheckPasswordHash(hashedPassword, password)
	if err != nil {
		return model.User{}, "", err
	}

	//token, errCreate := helpers.CreateJWT(username)
	token, err := auth.GenerateJWT(user.Email, user.UserName)
	if err != nil {
		return model.User{}, "", err
	}
	userID := find["_id"]
	_, errAddToken := repo.usersCollection.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"token": token}})
	if errAddToken != nil {
		return model.User{}, "", errAddToken
	}
	return user, token, nil

}

func (repo *UserRepository) DeleteUser(ctx context.Context, id string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "Deleted fail", err
	}

	result, err := repo.usersCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return "Deleted fail", err
	}

	if result.DeletedCount == 0 {
		return "Deleted fail", mongo.ErrNoDocuments
	}

	return "Deleted successfully", nil
}

func (repo *UserRepository) EditUser(ctx context.Context, id string, sth *serialize.User) (model.User, error) {
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

	result, err := repo.usersCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return model.User{}, err
	}

	var updatedUser model.User
	if result.MatchedCount == 1 {
		err = repo.usersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
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
func (repo *UserRepository) GetUserByID(ctx context.Context, userId string) (model.User, error) {

	var user model.User

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := repo.usersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil

}
func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := repo.usersCollection.Find(ctx, bson.M{})
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

func (repo *UserRepository) GetUserToken(ctx context.Context, token string) (model.User, error) {
	var user model.User
	err := repo.usersCollection.FindOne(ctx, bson.M{"token": token}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) EditUserToken(ctx context.Context, token string) error {
	_, err := repo.usersCollection.UpdateOne(ctx, bson.M{"token": token}, bson.M{"$unset": bson.M{"token": ""}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User

	err := repo.usersCollection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) ForgotPassword(ctx context.Context, phone, otp string) error {
	//helpers.SetOTP(otp, &model.User{OTP: otp, OTPExpiry: time.Now().Add(5 * time.Minute)})
	_, err := repo.usersCollection.UpdateOne(ctx, bson.M{"phone": phone}, bson.M{"$set": bson.M{"otp": otp, "otpexpiry": time.Now().Add(1 * time.Minute)}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) ResetPassword(ctx context.Context, otp, password string) error {
	hashedPassword, err := helpers.Hash(password)
	if err != nil {
		return err
	}

	_, resetPass := repo.usersCollection.UpdateOne(ctx, bson.M{"otp": otp}, bson.M{"$set": bson.M{"password": hashedPassword}})
	if resetPass != nil {
		return resetPass
	}

	_, deleteOTP := repo.usersCollection.UpdateOne(ctx, bson.M{"otp": otp}, bson.M{"$set": bson.M{"otp": "", "otpexpiry": time.Now()}})
	if deleteOTP != nil {
		return deleteOTP
	}
	return nil

}

func (repo *UserRepository) ForgotPasswordUseEmail(ctx context.Context, email, otp string) error {
	_, errGenerate := repo.usersCollection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"otp": otp, "otpexpiry": time.Now().Add(1 * time.Minute)}})
	if errGenerate != nil {
		return errGenerate
	}
	return nil
}
