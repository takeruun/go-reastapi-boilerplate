package database_test

import (
	"app/config"
	"app/database"
	"app/entity"
	"app/test_utils"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var userRepository database.UserRepository
var userDb *config.DB

func userSetUp(t *testing.T) func() {
	os.Setenv("GO_MODE", "test")
	userDb = test_utils.NewDB(t)
	userRepository = database.NewUserRepository(userDb)

	return func() {
		userDb.Exec("DELETE FROM users")
	}
}

func setIntialUserData() {
	data := []entity.User{
		{ID: 1, Name: "test1", Email: "test1@example.com", HashPassword: "d2fka"},
		{ID: 2, Name: "test2", Email: "test2@example.com", HashPassword: "d2fasdlfka"},
		{ID: 3, Name: "test3", Email: "test3@example.com", HashPassword: "d2fka"},
	}
	userDb.Create(&data)
}

func TestUserFindAll(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	setIntialUserData()

	t.Run("success", func(t *testing.T) {
		result, err := userRepository.FindAll()

		assert.NoError(t, err)
		assert.Equal(t, len(result), 3)
	})

}

func TestUserFind(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	setIntialUserData()

	var userId uint64 = 1

	t.Run("success", func(t *testing.T) {
		result, err := userRepository.Find(userId)

		assert.NoError(t, err)
		assert.Equal(t, userId, result.ID)
	})

	t.Run("If the user is not found", func(t *testing.T) {
		userId = 0

		_, err := userRepository.Find(userId)

		assert.Error(t, err)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestUserCreate(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	var u = &entity.User{
		Name:         "test4",
		Email:        "test4@example.com",
		HashPassword: "d2fka",
	}

	t.Run("success", func(t *testing.T) {
		result, err := userRepository.Create(u)

		assert.NoError(t, err)
		assert.NotEmpty(t, result.ID)
		assert.Equal(t, u.Name, result.Name)
	})

	t.Run("If the same email address is registered", func(t *testing.T) {
		setIntialUserData()
		u.Email = "test3@example.com"

		_, err := userRepository.Create(u)

		assert.Error(t, err)
		assert.Equal(t, int(err.(*mysql.MySQLError).Number), 1062)
	})
}

func TestUserFindByEmail(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	setIntialUserData()

	var email = "test1@example.com"

	t.Run("success", func(t *testing.T) {
		result, err := userRepository.FindByEmail(email)

		assert.NoError(t, err)
		assert.Equal(t, email, result.Email)
		assert.NotEmpty(t, result.ID)
	})

	t.Run("If the user is not found", func(t *testing.T) {
		email = "test10@example.com"

		_, err := userRepository.FindByEmail(email)
		assert.Error(t, err)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestUserUpdate(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	setIntialUserData()

	var u = &entity.User{
		ID:   1,
		Name: "update_test1",
	}

	t.Run("success", func(t *testing.T) {
		result, err := userRepository.Update(u)

		assert.NoError(t, err)
		assert.Equal(t, u.Name, result.Name)
	})

	t.Run("If the same email address is registered", func(t *testing.T) {
		u.Email = "test2@example.com"

		_, err := userRepository.Update(u)

		assert.Error(t, err)
		assert.Equal(t, int(err.(*mysql.MySQLError).Number), 1062)
	})
}

func TestUserDelete(t *testing.T) {
	setup := userSetUp(t)
	defer setup()

	setIntialUserData()

	var deleteUserId uint64 = 1

	t.Run("success", func(t *testing.T) {
		err := userRepository.Delete(deleteUserId)

		assert.NoError(t, err)
	})
}
