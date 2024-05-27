package service

import "bookstore/dao/admin"

type AdminService struct {
	adminRepo *admin.AdminRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminRepo: admin.NewAdminRepository(),
	}
}
