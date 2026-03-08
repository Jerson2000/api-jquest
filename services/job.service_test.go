package services

import (
	"context"
	"net/http"
	"testing"

	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/enums"
	"github.com/jerson2000/jquest/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repositories
type MockJobRepo struct {
	mock.Mock
}

func (m *MockJobRepo) Create(ctx context.Context, job models.Job) (models.Job, error) {
	args := m.Called(ctx, job)
	return args.Get(0).(models.Job), args.Error(1)
}

func (m *MockJobRepo) GetByID(ctx context.Context, id int) (models.Job, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Job), args.Error(1)
}

func (m *MockJobRepo) GetAll(ctx context.Context, page, limit int) ([]models.Job, int64, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]models.Job), args.Get(1).(int64), args.Error(2)
}

func (m *MockJobRepo) GetByCompanyID(ctx context.Context, companyId int, page int, limit int) ([]models.Job, int64, error) {
	args := m.Called(ctx, companyId, page, limit)
	return args.Get(0).([]models.Job), args.Get(1).(int64), args.Error(2)
}

func (m *MockJobRepo) Update(ctx context.Context, id int, job models.Job) (models.Job, error) {
	args := m.Called(ctx, id, job)
	return args.Get(0).(models.Job), args.Error(1)
}

func (m *MockJobRepo) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockRecruiterRepo struct {
	mock.Mock
}

func (m *MockRecruiterRepo) Create(ctx context.Context, recruiter models.Recruiter) (models.Recruiter, error) {
	args := m.Called(ctx, recruiter)
	return args.Get(0).(models.Recruiter), args.Error(1)
}

func (m *MockRecruiterRepo) GetByUserID(ctx context.Context, userId int) (models.Recruiter, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(models.Recruiter), args.Error(1)
}

func (m *MockRecruiterRepo) GetAll(ctx context.Context) ([]models.Recruiter, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Recruiter), args.Error(1)
}

func (m *MockRecruiterRepo) GetByCompanyID(ctx context.Context, companyId int) ([]models.Recruiter, error) {
	args := m.Called(ctx, companyId)
	return args.Get(0).([]models.Recruiter), args.Error(1)
}

func (m *MockRecruiterRepo) GetRecruiters(ctx context.Context) ([]models.Recruiter, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Recruiter), args.Error(1)
}

func (m *MockRecruiterRepo) Update(ctx context.Context, userId int, recruiter models.Recruiter) (models.Recruiter, error) {
	args := m.Called(ctx, userId, recruiter)
	return args.Get(0).(models.Recruiter), args.Error(1)
}

func (m *MockRecruiterRepo) Delete(ctx context.Context, userId int) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func TestCreateJobPost_Success(t *testing.T) {
	mockJobRepo := new(MockJobRepo)
	mockRecruiterRepo := new(MockRecruiterRepo)
	service := &jobService{
		jobRepo:       mockJobRepo,
		recruiterRepo: mockRecruiterRepo,
	}

	ctx := context.Background()
	userId := 1
	companyId := 10
	dto := dtos.JobCreateJobRequestDto{
		CompanyId:   companyId,
		Title:       "Software Engineer",
		Description: "Go developer needed",
		Location:    "Remote",
		JobType:     "Full-time",
		Status:      "Open",
	}

	recruiter := models.Recruiter{UserId: userId, CompanyId: companyId}
	mockRecruiterRepo.On("GetByUserID", ctx, userId).Return(recruiter, nil)

	createdJob := models.Job{
		Id:          100,
		CompanyId:   companyId,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      enums.JobStatusOpen,
	}

	// Expect Create to be called. Note: Matcher for struct might need improved precision in real scenarios
	mockJobRepo.On("Create", ctx, mock.AnythingOfType("models.Job")).Return(createdJob, nil)

	// Expect GetByID to be called to fetch full details
	mockJobRepo.On("GetByID", ctx, createdJob.Id).Return(createdJob, nil)

	result := service.CreateJobPost(ctx, userId, dto)

	assert.Equal(t, http.StatusCreated, result.Status)
	assert.NotNil(t, result.Data)
	assert.Equal(t, "Software Engineer", result.Data.Title)
	mockRecruiterRepo.AssertExpectations(t)
	mockJobRepo.AssertExpectations(t)
}

func TestCreateJobPost_Forbidden_NotRecruiter(t *testing.T) {
	mockJobRepo := new(MockJobRepo)
	mockRecruiterRepo := new(MockRecruiterRepo)
	service := &jobService{
		jobRepo:       mockJobRepo,
		recruiterRepo: mockRecruiterRepo,
	}

	ctx := context.Background()
	userId := 2
	dto := dtos.JobCreateJobRequestDto{CompanyId: 10}

	// Mock recruiter not found or similar error
	mockRecruiterRepo.On("GetByUserID", ctx, userId).Return(models.Recruiter{}, assert.AnError)

	result := service.CreateJobPost(ctx, userId, dto)

	assert.Equal(t, http.StatusForbidden, result.Status)
	assert.NotNil(t, result.Error)
	assert.Contains(t, result.Error.Message, "not a registered recruiter")
	mockRecruiterRepo.AssertExpectations(t)
}

func TestGetJobPost_Success(t *testing.T) {
	mockJobRepo := new(MockJobRepo)
	mockRecruiterRepo := new(MockRecruiterRepo)
	service := &jobService{
		jobRepo:       mockJobRepo,
		recruiterRepo: mockRecruiterRepo,
	}

	ctx := context.Background()
	page, limit := 1, 10
	jobs := []models.Job{
		{Id: 1, Title: "Job 1"},
		{Id: 2, Title: "Job 2"},
	}
	count := int64(2)

	mockJobRepo.On("GetAll", ctx, page, limit).Return(jobs, count, nil)

	result := service.GetJobPost(ctx, page, limit)

	assert.Equal(t, http.StatusOK, result.Status)
	assert.NotNil(t, result.Data)
	assert.Equal(t, 2, len(result.Data.Data))
	assert.Equal(t, count, result.Data.TotalCount)
	mockJobRepo.AssertExpectations(t)
}
