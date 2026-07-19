package service

import (
	"context"
	"errors"
	"strings"
	"time"

	authdto "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/dto"
	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
	userdto "github.com/JCKFinland/jck-connect/backend/internal/domain/user/dto"
	userservice "github.com/JCKFinland/jck-connect/backend/internal/domain/user/service"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	"github.com/JCKFinland/jck-connect/backend/pkg/jwt"
)

// Service defines authentication business logic.
type Service interface {
	// Login authenticates a Pi user.
	//
	// If the user does not yet exist in the local database,
	// a new account is automatically created.
	Login(
		ctx context.Context,
		req *authdto.LoginRequest,
	) (*authdto.LoginResponse, error)
}

type service struct {
	userService userservice.Service
	jwtManager  *jwt.Manager
}

// New creates a new authentication service.
func New(
	userService userservice.Service,
	jwtManager *jwt.Manager,
) Service {
	return &service{
		userService: userService,
		jwtManager:  jwtManager,
	}
}


// Login authenticates a Pi user.
//
// If the user does not already exist, a new local account
// is automatically created.
func (s *service) Login(
	ctx context.Context,
	req *authdto.LoginRequest,
) (*authdto.LoginResponse, error) {

	if req == nil {
		return nil, sharedErrors.BadRequest(
			errors.New("login request cannot be nil"),
		)
	}

	req.PiUID = strings.TrimSpace(req.PiUID)
	req.PiUsername = strings.TrimSpace(req.PiUsername)

	if req.PiUID == "" {
		return nil, sharedErrors.PiUIDRequired(nil)
	}

	if req.PiUsername == "" {
		return nil, sharedErrors.PiUsernameRequired(nil)
	}

	// Attempt to locate the user.
	user, err := s.userService.GetByPiUID(ctx, req.PiUID)
	if err != nil {

		// User not found.
		// Create a new local account.
		appErr, ok := err.(*sharedErrors.AppError)
		if !ok || appErr.Code != sharedErrors.CodeNotFound {
			return nil, err
		}

		user = &entity.User{
			PiUID:       req.PiUID,
			PiUsername:  req.PiUsername,
			DisplayName: req.PiUsername,
		}

		if err := s.userService.Create(ctx, user); err != nil {
			return nil, err
		}

		// Reload the user so we have the generated fields.
		user, err = s.userService.GetByPiUID(ctx, req.PiUID)
		if err != nil {
			return nil, err
		}
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(
		user.ID,
		user.PiUID,
		user.PiUsername,
		string(user.Role),
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	response := &authdto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresAt:    time.Now().Add(s.jwtManager.AccessTokenDuration()),
		User:         toUserResponse(user),
	}

	return response, nil
}

// toUserResponse converts a User entity into its public DTO.
func toUserResponse(user *entity.User) userdto.UserResponse {
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
