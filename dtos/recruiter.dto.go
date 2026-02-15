package dtos

type RecruiterCreateRequestDto struct {
	CompanyId int    `json:"companyId" binding:"required"`
	Position  string `json:"position" binding:"required"`
}

type RecruiterResponseDto struct {
	Id         int                 `json:"id"`
	UserId     int                 `json:"userId"`
	CompanyId  int                 `json:"companyId"`
	Position   string              `json:"position"`
	IsVerified bool                `json:"isVerified"`
	Company    *CompanyResponseDto `json:"company,omitempty"`
}
