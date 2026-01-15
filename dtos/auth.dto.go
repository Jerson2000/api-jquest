package dtos

type AuthLoginRequestDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSignupRequestDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRefreshRequestDto struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type AuthResponseDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
