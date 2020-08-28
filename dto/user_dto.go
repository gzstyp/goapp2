package dto

import "com.fwtai/app2/model"

type UserDto struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:        user.ID,
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
