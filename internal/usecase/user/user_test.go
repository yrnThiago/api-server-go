package usecase

import (
	"testing"

	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/tests/mocks"
	"go.uber.org/mock/gomock"
)

func Test_FindUserByID(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	id := "12346"
	repository := mocks.NewMockIUserRepository(control)
	service := NewUserUseCase(repository)

	repository.EXPECT().GetById(id).Return(entity.NewUser("test@test.com", "123456"), nil)

	user, err := service.GetById(id)
	if err != nil {
		t.FailNow()
		return
	}

	if user.Email != "test@test.com" {
		t.FailNow()
		return
	}
}
