package manager

import (
	"context"
	"errors"
	"fmt"

	"golang-core/api/src/client/request"
	"golang-core/api/src/client/response"
	"golang-core/api/src/common/crypto"
	"golang-core/api/src/common/helper"
	"golang-core/api/src/common/orm"
	"golang-core/api/src/infrastructure/repository"
	"golang-core/api/src/infrastructure/repository/entity"
)

type UserManager interface {
	Insert(ctx context.Context, resquest request.UserInsertRequest) (response.UserInfoResponse, error)
	Update(ctx context.Context, request request.UserUpdateRequest) (*response.UserInfoResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (response.UserInfoResponse, error)
	ListByIds(ctx context.Context, limit int, offset int) (*orm.PaginationData[response.UserInfoResponse], error)
}

type DefaultUserManager struct {
	userRepository repository.UserRepository
}

func NewUserManager(userRepository repository.UserRepository) UserManager {
	return &DefaultUserManager{
		userRepository: userRepository,
	}
}

func (m *DefaultUserManager) Insert(ctx context.Context, request request.UserInsertRequest) (res response.UserInfoResponse, err error) {
	hashedPassword, err := crypto.HashPassword(request.Password, "df3ca6fc6575537ccd12beff45ab31afc23977a378d535efbbd0a5c5811a10d1")
	if err != nil {
		return res, fmt.Errorf("Failed to hash password: %w", err)
	}
	newUser := entity.User{
		Email:    &request.Email,
		Name:     &request.Name,
		Password: &hashedPassword,
	}
	result, err := m.userRepository.Insert(ctx, newUser)
	if err != nil {
		return res, err
	}
	return *response.NewUserInfoResponse(result), nil
}

func (m *DefaultUserManager) Update(ctx context.Context, request request.UserUpdateRequest) (res *response.UserInfoResponse, err error) {
	user, err := m.userRepository.FindById(ctx, request.ID)
	if err != nil {
		return res, err
	}
	if user == nil {
		return nil, nil
	}
	helper.FromUserUpdateRequest(request, user)
	_, err = m.userRepository.Update(ctx, user)
	if err != nil {
		return res, err
	}
	return response.NewUserInfoResponse(user), nil
}

func (m *DefaultUserManager) Delete(ctx context.Context, id string) error {
	user, err := m.userRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found.")
	}
	err = m.userRepository.Delete(ctx, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *DefaultUserManager) FindById(ctx context.Context, id string) (res response.UserInfoResponse, err error) {
	result, err := m.userRepository.FindById(ctx, id)
	if result == nil || err != nil {
		return res, err
	}
	return *response.NewUserInfoResponse(result), nil
}

func (m *DefaultUserManager) ListByIds(ctx context.Context, limit int, offset int) (*orm.PaginationData[response.UserInfoResponse], error) {
	res, err := m.userRepository.ListByIds(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	converted := &orm.PaginationData[response.UserInfoResponse]{
		Total:         res.Total,
		CurrentOffset: res.CurrentOffset,
		Data:          helper.ToListUserInforResponse(res.Data),
	}
	return converted, nil
}
