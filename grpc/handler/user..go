package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/AdiKhoironHasan/bookservices-protobank/proto/book"
	"github.com/AdiKhoironHasan/bookservices-users/domain/assembler"
	"github.com/AdiKhoironHasan/bookservices-users/domain/entity"
	"github.com/AdiKhoironHasan/bookservices-users/proto/user"
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
	}).Select("id, name, email, password, role, created_at, updated_at").Rows()
	if err != nil {
		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	defer rows.Close()
	for rows.Next() {
		User := entity.User{}
		rows.Scan(&User.ID, &User.Name, &User.Email, &User.Password, &User.Role, &User.CreatedAt, &User.UpdatedAt)
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
		Name:     UserReq.Name,
		Email:    UserReq.Email,
		Password: UserReq.Password,
		Role:     UserReq.Role,
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
			Email:     UserEntity.Email,
			Password:  UserEntity.Password,
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
		Name:     UserReq.Name,
		Email:    UserReq.Email,
		Password: UserReq.Password,
		Role:     UserReq.Role,
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

func (c *Handler) Login(ctx context.Context, UserReq *user.UserLoginReq) (*user.UserLoginRes, error) {
	userData := entity.User{}

	err := c.repo.DB.WithContext(ctx).Where(&entity.User{
		Email:    UserReq.Email,
		Password: UserReq.Password,
	}).First(&userData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.New(http.StatusUnauthorized, "user unauthorized").Err()
		}

		return nil, status.New(http.StatusInternalServerError, err.Error()).Err()
	}

	userRes := &user.User{
		Id:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Role:      userData.Role,
		CreatedAt: userData.CreatedAt.String(),
		UpdatedAt: userData.UpdatedAt.String(),
	}

	return &user.UserLoginRes{
		User: userRes,
	}, nil
}
