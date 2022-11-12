package usecase_test

import (
	"app/test_utils/mock_database"
	"app/test_utils/mock_service"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockTodoRepository *mock_database.MockTodoRepository

func todoSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepository = mock_database.NewMockTodoRepository(ctrl)
	mockSessionService = mock_service.NewMockSessionService(ctrl)

	return func() {}
}
