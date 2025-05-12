package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/tests/mocks"
	"go.uber.org/mock/gomock"
)

func TestUserUseCase_Add(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	type testCase struct {
		name         string
		userInputDto dto.UserInputDto
		mockSetup    func(repo *mocks.MockIUserRepository)
		expected     *dto.UserOutputDto
		expectError  bool
	}

	userInputTest := dto.UserInputDto{Email: "test@email.com", Password: "123456"}
	userTest := entity.NewUser(userInputTest.Email, userInputTest.Password)

	tests := []testCase{
		{
			name:         "return user created",
			userInputDto: userInputTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().Add(userInputTest).Return(userTest, nil)
			},
			expected:    dto.NewUserOutputDto(userTest),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockIUserRepository(control)
			tt.mockSetup(repo)

			usecase := NewUserUseCase(repo)
			user, err := usecase.Add(userInputTest)

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
			name: "return valid user when everything looks good",
			user: userTest,
			mockSetup: func(repo *mocks.MockIUserRepository) {
				repo.EXPECT().GetById(userTest.ID).Return(userTest, nil)
			},
			expected:    dto.NewUserOutputDto(userTest),
			expectError: false,
		},
		{
			name: "return error when user id not found",
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
