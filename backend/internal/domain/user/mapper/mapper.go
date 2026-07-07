package mapper

import (
	userdto "github.com/JCKFinland/jck-connect/backend/internal/domain/user/dto"
	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
)

// ToUserResponse converts a User entity into a UserResponse DTO.
func ToUserResponse(
	user *entity.User,
) userdto.UserResponse {

	if user == nil {
		return userdto.UserResponse{}
	}

	return userdto.UserResponse{
		ID:          user.ID,
		PiUID:       user.PiUID,
		PiUsername:  user.PiUsername,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        string(user.Role),
		Status:      string(user.Status),
		CreatedAt:   user.CreatedAt,
	}
}