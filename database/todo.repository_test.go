package database_test

import (
	"app/config"
	"app/database"
	"app/entity"
	"app/test_utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var todoRepository database.TodoRepository
var todoDb *config.DB

func todoSetUp(t *testing.T) func() {
	os.Setenv("GO_MODE", "test")
	todoDb = test_utils.NewDB(t)
	todoRepository = database.NewTodoRepository(todoDb)

	return func() {
		todoDb.Exec("DELETE FROM todos")
	}
}

func setIntialTodoData() {
	data := []entity.Todo{
		{ID: 1, UserId: uint64(1), Title: "title_1", Description: "description_1"},
		{ID: 2, UserId: uint64(2), Title: "title_2", Description: "description_2"},
		{ID: 3, UserId: uint64(3), Title: "title_3", Description: "description_3"},
	}
	todoDb.Create(&data)
}

func TestTodoFindAll(t *testing.T) {
	setup := todoSetUp(t)
	defer setup()

	setIntialTodoData()

	t.Run("success", func(t *testing.T) {
		result, err := todoRepository.FindAll(1)

		assert.NoError(t, err)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, uint64(1), result[0].UserId)
	})
}
