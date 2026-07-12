package dtos

type SkillCreateRequestDto struct {
	Name string `json:"name" binding:"required"`
}

type SkillResponseDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
