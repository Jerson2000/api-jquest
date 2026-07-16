package dtos

type CandidateCreateRequestDto struct {
	FirstName       string  `json:"firstName" validate:"required"`
	LastName        string  `json:"lastName" validate:"required"`
	Email           string  `json:"email" validate:"omitempty,email"`
	Phone           string  `json:"phone"`
	LinkedInURL     string  `json:"linkedinUrl"`
	ResumeURL       string  `json:"resumeUrl"`
	TotalExperience float32 `json:"totalExperience"`
	CurrentTitle    string  `json:"currentTitle"`
	CurrentLocation string  `json:"currentLocation"`
	Bio             string  `json:"bio"`
	ExpectedSalary  *int    `json:"expectedSalary"`
}

type CandidateResponseDto struct {
	Id              int                    `json:"id"`
	FirstName       string                 `json:"firstName"`
	LastName        string                 `json:"lastName"`
	Email           string                 `json:"email"`
	Phone           string                 `json:"phone"`
	LinkedInURL     string                 `json:"linkedinUrl"`
	ResumeURL       string                 `json:"resumeUrl"`
	TotalExperience float32                `json:"totalExperience"`
	CurrentTitle    string                 `json:"currentTitle"`
	CurrentLocation string                 `json:"currentLocation"`
	Bio             string                 `json:"bio,omitempty"`
	ExpectedSalary  *int                   `json:"expectedSalary,omitempty"`
	Skills          []string               `json:"skills,omitempty"`
	Educations      []EducationResponseDto `json:"educations,omitempty"`
}

