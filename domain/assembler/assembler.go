package assembler

import (
	"github.com/AdiKhoironHasan/bookservices/domain/entity"
	protoUser "github.com/AdiKhoironHasan/bookservices/proto/user"
)

func ToResponseUserList(Users []entity.User, ch chan<- []*protoUser.User) {
	value := []*protoUser.User{}
	for _, val := range Users {
		value = append(value, &protoUser.User{
			Id:        val.ID,
			Name:      val.Name,
			Role:      val.Role,
			CreatedAt: val.CreatedAt.String(),
			UpdatedAt: val.UpdatedAt.String(),
		})
	}

	ch <- value
}
