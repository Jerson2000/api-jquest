package dtos

type CompanyCreateRequestDto struct {
	Name        string  `json:"name" binding:"required"`
	Industry    string  `json:"industry" binding:"required"`
	Website     *string `json:"website"`
	Location    *string `json:"location"`
	CompanySize *string `json:"companySize"`
}

type CompanyResponseDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	CompanySize string `json:"companySize"`
}

type CompanyApplyRequestDto struct {
	Company CompanyCreateRequestDto `json:"company" binding:"required"`
	User    UserCreateRequestDto    `json:"user" binding:"required"`
}
