package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"github.com/AdiKhoironHasan/bookservices/domain/assembler"
	"github.com/AdiKhoironHasan/bookservices/domain/entity"
	"github.com/AdiKhoironHasan/bookservices/proto/user"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (c *Handler) Ping(ctx context.Context, _ *user.PingReq) (*user.PingRes, error) {
	var now string
	err := c.repo.DB.WithContext(ctx).Raw("select now ()").Scan(&now).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &user.PingRes{
		Message: now,
	}, nil
}

// List is a function
func (c *Handler) List(ctx context.Context, userReq *user.UserListReq) (*user.UserListRes, error) {
	users := []entity.User{}

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
		users = append(users, User)
	}

	ch := make(chan []*user.User)
	defer close(ch)
	go assembler.ToResponseUserList(users, ch)

	return &user.UserListRes{
		Users: <-ch,
	}, nil
}

func (c *Handler) Store(ctx context.Context, UserReq *user.UserStoreReq) (*user.UserStoreRes, error) {
	UserEntity := entity.User{
		Name: UserReq.Name,
		Role: UserReq.Role,
	}

	err := c.repo.DB.WithContext(ctx).Create(&UserEntity).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &user.UserStoreRes{}, nil
}

func (c *Handler) Detail(ctx context.Context, UserReq *user.UserDetailReq) (*user.UserDetailRes, error) {
	UserEntity := entity.User{}

	err := c.repo.DB.WithContext(ctx).First(&UserEntity, UserReq.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.New(http.StatusNotFound, "record not found").Err()
		}

		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &user.UserDetailRes{
		User: &user.User{
			Id:        UserEntity.ID,
			Name:      UserEntity.Name,
			Role:      UserEntity.Role,
			CreatedAt: UserEntity.CreatedAt.String(),
			UpdatedAt: UserEntity.UpdatedAt.String(),
		},
	}, nil
}

func (c *Handler) Update(ctx context.Context, UserReq *user.UserUpdateReq) (*user.UserUpdateRes, error) {
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

	return &user.UserUpdateRes{}, nil
}

func (c *Handler) Delete(ctx context.Context, UserReq *user.UserDeleteReq) (*user.UserDeleteRes, error) {
	err := c.repo.DB.WithContext(ctx).First(&entity.User{}, UserReq.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.New(http.StatusNotFound, "record not found").Err()
		}

		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	_, err = c.grpcClient.Book.Delete(ctx, &book.BookDeleteReq{Id: UserReq.Id})
	if err != nil {
		if status.Code(err) != http.StatusNotFound {
			return nil, err
		}
	}

	err = c.repo.DB.WithContext(ctx).Delete(&entity.User{}, UserReq.Id).Error
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	return &user.UserDeleteRes{}, nil
}
