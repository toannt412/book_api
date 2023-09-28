package service

import (
	"bookstore/serialize"
	"context"
	"testing"
)

// type GetAdminTestSuite struct {
// 	suite.Suite
// }

// func WrongID(ctx context.Context, id string) (*serialize.Admin, error) {
// 	// id = "650d4080be6273000b09ffea"
// 	// user, err := admin.GetAdminByID(ctx, id)
// 	return nil, nil
// }

func TestPassGetAdminUserByID(t *testing.T) {
	var admin *serialize.Admin
	adminID := "650d4080be6273000b09ffea"
	//admin := model.Admin{}
	got, _ := GetAdminUserByID(context.TODO(), adminID)
	t.Log(got)
	if got == admin {
		t.Log("ok")
	} else {
		t.Errorf("got %v, want %v", got, admin)
	}

}
