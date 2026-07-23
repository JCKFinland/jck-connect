package service

import (
	"context"
	"time"

	authdto "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/dto"
	"github.com/JCKFinland/jck-connect/backend/internal/domain/auth/pi"
	userdto "github.com/JCKFinland/jck-connect/backend/internal/domain/user/dto"
	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
	userservice "github.com/JCKFinland/jck-connect/backend/internal/domain/user/service"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	"github.com/JCKFinland/jck-connect/backend/pkg/jwt"
)

type Service interface {
	Login(
		ctx context.Context,
		req *authdto.LoginRequest,
	) (*authdto.LoginResponse, error)
}

type service struct {
	userService userservice.Service
	jwtManager  *jwt.Manager
	verifier    pi.Verifier
}

func New(
	userService userservice.Service,
	jwtManager *jwt.Manager,
	verifier pi.Verifier,
) Service {

	return &service{
		userService: userService,
		jwtManager:  jwtManager,
		verifier:    verifier,
	}
}

func (s *service) Login(
	ctx context.Context,
	req *authdto.LoginRequest,
) (*authdto.LoginResponse, error) {

	if req == nil {
		return nil, sharedErrors.BadRequest(nil)
	}

	if req.AccessToken == "" {
		return nil, sharedErrors.BadRequest(nil)
	}

	//----------------------------------------------------------
	// Verify access token with Pi Platform
	//----------------------------------------------------------

	piUser, err := s.verifier.Verify(
		ctx,
		req.AccessToken,
	)
	if err != nil {
		return nil, err
	}

	//----------------------------------------------------------
	// Find local user
	//----------------------------------------------------------

	user, err := s.userService.GetByPiUID(
		ctx,
		piUser.UID,
	)

	if err != nil {

		appErr, ok := err.(*sharedErrors.AppError)

		if !ok || appErr.Code != sharedErrors.CodeNotFound {
			return nil, err
		}

		//------------------------------------------------------
		// First Login
		//------------------------------------------------------

		user = &entity.User{
			PiUID:       piUser.UID,
			PiUsername:  piUser.Username,
			DisplayName: piUser.Username,
		}

		if err := s.userService.Create(
			ctx,
			user,
		); err != nil {
			return nil, err
		}

		user, err = s.userService.GetByPiUID(
			ctx,
			piUser.UID,
		)

		if err != nil {
			return nil, err
		}

		//------------------------------------------------------
		// Wallet creation comes here in next step
		//------------------------------------------------------
	}

	//----------------------------------------------------------
	// Generate JWT
	//----------------------------------------------------------

	accessToken, err := s.jwtManager.GenerateAccessToken(
		user.ID,
		user.PiUID,
		user.PiUsername,
		string(user.Role),
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(
		user.ID,
	)
	if err != nil {
		return nil, err
	}

	return &authdto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresAt: time.Now().Add(
			s.jwtManager.AccessTokenDuration(),
		),
		User: toUserResponse(user),
	}, nil
}

func toUserResponse(
	user *entity.User,
) userdto.UserResponse {

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
