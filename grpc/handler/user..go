package handler

import (
	"context"
	"errors"
	"net/http"

	protoUser "github.com/AdiKhoironHasan/bookservice-protobank/proto/user"
	"github.com/AdiKhoironHasan/bookservices/domain/entity"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (c *Handler) Ping(ctx context.Context, _ *protoUser.PingReq) (*protoUser.PingRes, error) {
	var now string
	err := c.repo.DB.Raw("select now ()").Scan(&now).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &protoUser.PingRes{
		Message: now,
	}, nil
}


// List is a function
func (c *Handler) List(ctx context.Context, userReq *protoUser.UserListReq) (*protoUser.UserListRes, error) {
	Users := []entity.User{}

	rows, err := c.repo.DB.WithContext(ctx).Model(&entity.User{}).Where(&entity.User{
		Name: userReq.Name,
		Role: userReq.Role,
	}).Select("id, name, role, created_at, updated_at").Rows()
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	defer rows.Close()
	for rows.Next() {
		User := entity.User{}
		rows.Scan(&User.ID, &User.Name, &User.Role, &User.CreatedAt, &User.UpdatedAt)
		Users = append(Users, User)
	}

	ch := make(chan []*protoUser.User)
	defer close(ch)

	go func(Users []entity.User, ch chan<- []*protoUser.User) {
		value := []*protoUser.User{}
		for _, val := range Users {
			value = append(value, &protoUser.User{
				Id:          val.ID,
				Name:        val.Name,
				Role:        val.Role,
				CreatedAt:   val.CreatedAt.String(),
				UpdatedAt:   val.UpdatedAt.String(),
			})
		}

		ch <- value
	}(Users, ch)

	return &protoUser.UserListRes{
		Users: <-ch,
	}, nil
}

func (c *Handler) Store(ctx context.Context, UserReq *protoUser.UserStoreReq) (*protoUser.UserStoreRes, error) {
	UserEntity := entity.User{
		Name: UserReq.Name,
		Role: UserReq.Role,
	}

	err := c.repo.DB.WithContext(ctx).Create(&UserEntity).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &protoUser.UserStoreRes{}, nil
}

func (c *Handler) Detail(ctx context.Context, UserReq *protoUser.UserDetailReq) (*protoUser.UserDetailRes, error) {
	UserEntity := entity.User{}

	err := c.repo.DB.WithContext(ctx).First(&UserEntity, UserReq.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.New(http.StatusNotFound, "record not found").Err()
		}

		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &protoUser.UserDetailRes{
		User: &protoUser.User{
			Id:          UserEntity.ID,
			Name:        UserEntity.Name,
			Role:        UserEntity.Role,
			CreatedAt:   UserEntity.CreatedAt.String(),
			UpdatedAt:   UserEntity.UpdatedAt.String(),
		},
	}, nil
}

func (c *Handler) Update(ctx context.Context, UserReq *protoUser.UserUpdateReq) (*protoUser.UserUpdateRes, error) {
	err := c.repo.DB.WithContext(ctx).First(&entity.User{}, UserReq.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.New(http.StatusNotFound, "record not found").Err()
		}

		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	err = c.repo.DB.WithContext(ctx).Model(&entity.User{ID: UserReq.Id}).Updates(&entity.User{
		Name: UserReq.Name,
		Role: UserReq.Role,
	}).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &protoUser.UserUpdateRes{}, nil
}

func (c *Handler) Delete(ctx context.Context, UserReq *protoUser.UserDeleteReq) (*protoUser.UserDeleteRes, error) {
	err := c.repo.DB.WithContext(ctx).First(&entity.User{}, UserReq.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.New(http.StatusNotFound, "record not found").Err()
		}

		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	err = c.repo.DB.WithContext(ctx).Delete(&entity.User{}, UserReq.Id).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &protoUser.UserDeleteRes{}, nil
}
