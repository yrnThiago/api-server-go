package usecase

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/tests/mocks"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestUserUseCase_Add(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name        string
		input       dto.UserInputDto
		mockSetup   func(repo *mocks.MockIUserRepository)
		expected    *dto.UserOutputDto
		expectError bool
	}

	inputTest := dto.UserInputDto{Email: "test@test.com", Password: "123456"}
	inputTest.Password, _ = utils.GenerateHashPassword(inputTest.Password)

	newUser := entity.NewUser(inputTest.Email, inputTest.Password)

	tests := []testCase{
		{
			name:        "Add return error when input is not valid",
			mockSetup:   func(repo *mocks.MockIUserRepository) {},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "Add return database error",
			input: inputTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().Add(gomock.Any()).Return(nil, fmt.Errorf("db error"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "Add return new user output when everything good",
			input: inputTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().Add(gomock.Any()).Return(newUser, nil)
			},
			expected:    dto.NewUserOutputDto(newUser),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			config.Logger = zap.NewNop()
			usecase := NewUserUseCase(repo)
			user, err := usecase.Add(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Email, user.Email)
				assert.Equal(t, tt.expected.Orders, user.Orders)
			}
		})
	}
}

func TestUserUseCase_GetMany(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name        string
		mockSetup   func(repo *mocks.MockIUserRepository)
		expected    []*dto.UserOutputDto
		expectError bool
	}

	userTest := entity.NewUser("test@test.com", "123456")
	userTest2 := entity.NewUser("test2@test2.com", "654321")

	tests := []testCase{
		{
			name: "GetMany return users when everything looks good",
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetMany().Return(nil, fmt.Errorf("db error"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "GetMany return users when everything looks good",
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetMany().Return([]*entity.User{userTest, userTest2}, nil)
			},
			expected: []*dto.UserOutputDto{
				dto.NewUserOutputDto(userTest),
				dto.NewUserOutputDto(userTest2),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			usecase := NewUserUseCase(repo)
			user, err := usecase.GetMany()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}

func TestUserUseCase_GetById(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name        string
		user        *entity.User
		mockSetup   func(repo *mocks.MockIUserRepository)
		expected    *dto.UserOutputDto
		expectError bool
	}

	userTest := entity.NewUser("test@test.com", "123456")
	userNotFound := entity.NewUser("notfound@notfound.com", "123456")

	tests := []testCase{
		{
			name: "GetById return valid user when everything looks good",
			user: userTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userTest.ID).Return(userTest, nil)
			},
			expected:    dto.NewUserOutputDto(userTest),
			expectError: false,
		},
		{
			name: "GetById return error when user id not found",
			user: userNotFound,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userNotFound.ID).Return(nil, fmt.Errorf("id not found"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			usecase := NewUserUseCase(repo)
			user, err := usecase.GetById(tt.user.ID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}

func TestUserUseCase_GetByLogin(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name        string
		email       string
		mockSetup   func(repo *mocks.MockIUserRepository)
		expected    *entity.User
		expectError bool
	}

	userTest := entity.NewUser("test@test.com", "123456")
	invalidEmail := uuid.NewString()
	emptyEmail := ""

	tests := []testCase{
		{
			name:  "GetByLogin return user when email is registered",
			email: userTest.Email,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetByLogin(userTest.Email).Return(userTest, nil)
			},
			expected:    userTest,
			expectError: false,
		},
		{
			name:  "GetByLogin return error when email not found",
			email: invalidEmail,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetByLogin(invalidEmail).Return(nil, fmt.Errorf("email not found"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "GetByLogin return error when email is empty",
			email: emptyEmail,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetByLogin(emptyEmail).Return(nil, fmt.Errorf("email not found"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			usecase := NewUserUseCase(repo)
			user, err := usecase.GetByLogin(tt.email)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}

func TestUserUseCase_UpdateById(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name        string
		id          string
		input       dto.UserInputDto
		mockSetup   func(repo *mocks.MockIUserRepository)
		expected    *dto.UserOutputDto
		expectError bool
	}

	inputTest := dto.UserInputDto{Email: "test@test.com", Password: "123456"}
	userTest := entity.NewUser("test@test.com", "123456")

	updatedUser := &entity.User{
		ID:       userTest.ID,
		Email:    inputTest.Email,
		Password: inputTest.Password,
	}

	tests := []testCase{
		{
			name:        "UpdateById return error when input is not valid",
			id:          uuid.NewString(),
			mockSetup:   func(repo *mocks.MockIUserRepository) {},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "UpdateById return error when id not found",
			id:    "invalid-id",
			input: inputTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById("invalid-id").Return(nil, fmt.Errorf("id not found"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "UpdateById return database error",
			id:    userTest.ID,
			input: inputTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userTest.ID).Return(userTest, nil)
				repo.EXPECT().UpdateById(userTest).Return(nil, fmt.Errorf("db error"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "UpdateById return updated user output when everything good",
			id:    userTest.ID,
			input: inputTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userTest.ID).Return(userTest, nil)
				repo.EXPECT().UpdateById(userTest).Return(updatedUser, nil)
			},
			expected:    dto.NewUserOutputDto(updatedUser),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			config.Logger = zap.NewNop()
			usecase := NewUserUseCase(repo)
			user, err := usecase.UpdateById(tt.id, tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}

func TestUserUseCase_DeleteById(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name        string
		id          string
		mockSetup   func(repo *mocks.MockIUserRepository)
		expected    *dto.UserOutputDto
		expectError bool
	}

	userTest := entity.NewUser("test@test.com", "123456")
	userOutputTest := dto.NewUserOutputDto(userTest)

	tests := []testCase{
		{
			name: "DeleteById return error id not found",
			id:   "invalid-id",
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById("invalid-id").Return(nil, fmt.Errorf("id not found"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "DeleteById return database error",
			id:   userTest.ID,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userTest.ID).Return(userTest, nil)
				repo.EXPECT().DeleteById(userTest.ID).Return(fmt.Errorf("db error"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "DeleteById return updated user output when everything good",
			id:   userTest.ID,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userTest.ID).Return(userTest, nil)
				repo.EXPECT().DeleteById(userTest.ID).Return(nil)
			},
			expected:    userOutputTest,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			config.Logger = zap.NewNop()
			usecase := NewUserUseCase(repo)
			user, err := usecase.DeleteById(tt.id)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}
