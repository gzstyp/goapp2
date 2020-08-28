package dto

import "com.fwtai/app2/model"

type UserDto struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// 即把密码隐藏或移除掉,通过[强制]类型转换??? 调用方式:dto.ToUserDto(user.(model.User))
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:        user.ID,
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
