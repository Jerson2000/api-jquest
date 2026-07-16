package dtos

type CompanyCreateRequestDto struct {
	Name        string  `json:"name" binding:"required"`
	Industry    string  `json:"industry" binding:"required"`
	Website     *string `json:"website"`
	Location    *string `json:"location"`
	CompanySize *string `json:"companySize"`
	Description *string `json:"description"`
	LogoURL     *string `json:"logoUrl"`
}

type CompanyResponseDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	CompanySize string `json:"companySize"`
	Description string `json:"description,omitempty"`
	LogoURL     string `json:"logoUrl,omitempty"`
}

type CompanyApplyRequestDto struct {
	Company CompanyCreateRequestDto `json:"company" binding:"required"`
	User    UserCreateRequestDto    `json:"user" binding:"required"`
}
