package usecase_test

import (
	"app/controller/dto"
	"app/entity"
	"app/test_utils/mock_database"
	"app/test_utils/mock_service"
	"app/usecase"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockTodoRepository *mock_database.MockTodoRepository

func todoSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepository = mock_database.NewMockTodoRepository(ctrl)
	mockSessionService = mock_service.NewMockSessionService(ctrl)

	return func() {}
}

func TestCreate(t *testing.T) {
	todoSetUp := todoSetUp(t)
	defer todoSetUp()

	var (
		title        = "test-test"
		description  = "description-test"
		expectedTodo = entity.Todo{Title: title, Description: description, UserId: uint64(1), UpdatedAt: time.Now(), CreatedAt: time.Now()}
	)

	t.Run("success", func(t *testing.T) {
		mockSessionService.EXPECT().GetSessionValue(gomock.Any(), "userId").Return(uint64(1), nil)
		mockTodoRepository.EXPECT().Create(&entity.Todo{Title: title, Description: description, UserId: 1}).Return(&expectedTodo, nil)
		todoUsecase := usecase.NewTodoUsecase(mockTodoRepository, mockSessionService)

		result, err := todoUsecase.Create(context.TODO(), &dto.TodoCreateRequestDto{Title: title, Description: description})

		assert.NoError(t, err)
		assert.Equal(t, expectedTodo.Title, result.Title)
		assert.Equal(t, expectedTodo.Description, result.Description)
	})
}
