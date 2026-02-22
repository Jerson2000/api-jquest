package dtos

type CandidateCreateRequestDto struct {
	FirstName       string  `json:"firstName" validate:"required"`
	LastName        string  `json:"lastName" validate:"required"`
	Email           string  `json:"email" validate:"required,email"`
	Phone           string  `json:"phone"`
	LinkedInURL     string  `json:"linkedinUrl"`
	ResumeURL       string  `json:"resumeUrl"`
	TotalExperience float32 `json:"totalExperience"`
	CurrentTitle    string  `json:"currentTitle"`
	CurrentLocation string  `json:"currentLocation"`
}

type CandidateResponseDto struct {
	Id              int     `json:"id"`
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	Email           string  `json:"email"`
	Phone           string  `json:"phone"`
	LinkedInURL     string  `json:"linkedinUrl"`
	ResumeURL       string  `json:"resumeUrl"`
	TotalExperience float32 `json:"totalExperience"`
	CurrentTitle    string  `json:"currentTitle"`
	CurrentLocation string  `json:"currentLocation"`
}
