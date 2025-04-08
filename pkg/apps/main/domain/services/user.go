package services

import (
	"context"

	"piroux.dev/yoping/api/pkg/apps/main/domain/adapters"
	"piroux.dev/yoping/api/pkg/apps/main/domain/models"
)

// Domain Service
type UserService struct {
	userRespository adapters.UserRespository
}

func NewUserService(
	userRespository adapters.UserRespository,
) *UserService {

	return &UserService{
		userRespository: userRespository,
	}
}

//
// Create/Update/Delete

type CreateUserRequest struct {
	models.UserData
}

type CreateUserResponse struct {
	Status string
	User   models.User
}

func (svc *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (rsp CreateUserResponse, err error) {
	user, err := models.NewUser(
		req.NameFull,
		req.PhoneNumber,
	)
	if err != nil {
		return CreateUserResponse{Status: "failed to build user"}, err
	}

	user, err = svc.userRespository.Create(user)
	if err != nil {
		return CreateUserResponse{Status: "failed to store user"}, err
	}

	rsp.User = *user
	rsp.Status = "ok"
	return
}

type UpdateUserRequest struct {
	User models.User
}

type UpdateUserResponse struct {
	Status string
	User   models.User
}

func (svc *UserService) UpdateUser(ctx context.Context, req UpdateUserRequest) (rsp UpdateUserResponse, err error) {
	return
}

type DeleteUserRequest struct {
	User models.User
}

type DeleteUserResponse struct {
	Status string
}

func (svc *UserService) DeleteUser(ctx context.Context, req DeleteUserRequest) (rsp DeleteUserResponse, err error) {
	err = svc.userRespository.Delete(&req.User)
	if err != nil {
		return DeleteUserResponse{Status: "failed to delete user"}, err
	}

	rsp.Status = "ok"
	return
}

//
// Get

type GetOneUserRequest struct {
	models.UserKey
}

type GetOneUserResponse struct {
	ResponseMetadata
	User models.User
}

func (svc *UserService) GetOneUser(ctx context.Context, req GetOneUserRequest) (rsp GetOneUserResponse, err error) {
	user, err := svc.userRespository.GetOne(req.UserKey.String())
	if err != nil {
		return GetOneUserResponse{ResponseMetadata: ResponseMetadata{Status: "failed to retrieve user"}}, err
	}

	rsp.User = *user
	rsp.ResponseMetadata.Status = "ok"
	return
}

type GetAllUsersRequest struct {
}

type GetAllUsersResponse struct {
	ResponseMetadata
	Users []models.User
}

func (svc *UserService) GetAllUsers(ctx context.Context, req GetAllUsersRequest) (rsp GetAllUsersResponse, err error) {
	users, err := svc.userRespository.GetAll()
	if err != nil {
		return GetAllUsersResponse{ResponseMetadata: ResponseMetadata{Status: "failed to retrieve users"}}, err
	}

	rsp.Users = make([]models.User, 0, len(users))
	for _, user := range users {
		if user != nil {
			rsp.Users = append(rsp.Users, *user)
		}
	}

	rsp.ResponseMetadata.Status = "ok"
	return
}

type GetUserContactsRequest struct {
	models.UserKey
}

type GetUserContactsResponse struct {
	ResponseMetadata
	User     models.User
	Contacts []models.User
}

func (svc *UserService) GetUserContacts(ctx context.Context, req GetUserContactsRequest) (rsp GetUserContactsResponse, err error) {
	user, err := svc.userRespository.GetOne(req.UserKey.String())
	if err != nil {
		return GetUserContactsResponse{ResponseMetadata: ResponseMetadata{Status: "failed to retrieve user"}}, err
	}

	contacts, err := svc.userRespository.GetContacts(req.UserKey.String())
	if err != nil {
		return GetUserContactsResponse{ResponseMetadata: ResponseMetadata{Status: "failed to retrieve user's contacts"}}, err
	}

	rsp.Contacts = make([]models.User, 0, len(contacts))
	for _, contact := range contacts {
		if contact != nil {
			rsp.Contacts = append(rsp.Contacts, *contact)
		}
	}
	rsp.User = *user
	rsp.ResponseMetadata.Status = "ok"
	return
}
