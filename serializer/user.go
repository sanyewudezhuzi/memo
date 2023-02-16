package serializer

import "github.com/sanyewudezhuzi/memo/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	CreateAt int64  `json:"create_at"`
}

// 序列化 User
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
