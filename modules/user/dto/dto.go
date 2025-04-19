package dto

import "xyz/modules/user/domain"

type (
	RegisterRequest struct {
		Username string `json:"username" form:"username" validate:"required,min=3,max=20"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
		NIK string `json:"nik" form:"nik" validate:"required,min=16,max=16"`
		Full_Name string `json:"full_name" form:"full_name" validate:"required,min=3,max=50"`
		Legal_Name string `json:"legal_name" form:"legal_name" validate:"required,min=3,max=50"`
		Place_Of_Birth string `json:"place_of_birth" form:"place_of_birth" validate:"required,min=3,max=50"`
		Date_Of_Birth string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
		Salary int `json:"salary" form:"salary" validate:"required"`
		Photo_KTP string `json:"photo_ktp" form:"photo_ktp" validate:"required"`
		Photo_Selfie string `json:"photo_selfie" form:"photo_selfie" validate:"required"`
	}
	LoginRequest struct {
		Username string `json:"username" form:"username" validate:"required,min=3,max=20"`
		Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
	}
	ResetPasswordRequest struct {
		Email string `json:"email" form:"email" validate:"required,email"`
	}
	ChangePasswordRequest struct {
		UserID      string `json:"user_id" form:"user_id" validate:"required"`
		OldPassword string `json:"old_password" form:"old_password" validate:"required,min=6,max=20"`
		NewPassword string `json:"new_password" form:"new_password" validate:"required,min=6,max=20"`
	}
	RequestActivationRequest struct {
		Email string `json:"email" form:"email" validate:"required,email"`
	}
	ActivateUserRequest struct {
		UserID string `json:"user_id" form:"user_id" validate:"required"`
		Token  string `json:"token" form:"token" validate:"required"`
	}
)

type (
	RegisterResponse struct {
		UserID int `json:"user_id"`
	}

	LoginResponse struct {
		UserID int    `json:"user_id"`
		Token  string `json:"token"`
	}
)

func RegisterRequestIntoUserDomain(request RegisterRequest) domain.User {
	return domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserDomainIntoRegisterResponse(user domain.User) RegisterResponse {
	return RegisterResponse{
		UserID: user.ID,
	}
}

func LoginRequestIntoUserDomain(request LoginRequest) domain.User {
	return domain.User{
		Username: request.Username,
		Password: request.Password,
	}
}

func UserDomainIntoLoginResponse(user domain.User, token string) LoginResponse {
	return LoginResponse{
		UserID: user.ID,
		Token:  token,
	}
}

func RequestActivationRequestIntoUserDomain(request RequestActivationRequest) domain.User {
	return domain.User{
		Email: request.Email,
	}
}