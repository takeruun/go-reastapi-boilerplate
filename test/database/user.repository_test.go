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
var db *config.DB

func setUp(t *testing.T) func() {
	os.Setenv("GO_MODE", "test")
	db = test_utils.NewDB(t)
	userRepository = database.NewUserRepository(db)

	return func() {
		db.Exec("DELETE FROM users")
	}
}

func setIntialData() {
	data := []entity.User{
		{Name: "test1", Email: "test1@example.com", HashPassword: "d2fka"},
		{Name: "test2", Email: "test2@example.com", HashPassword: "d2fasdlfka"},
		{Name: "test3", Email: "test3@example.com", HashPassword: "d2fka"},
	}
	db.Create(&data)
}

func TestFindAll(t *testing.T) {
	setup := setUp(t)
	defer setup()

	setIntialData()

	t.Run("success", func(t *testing.T) {
		result, err := userRepository.FindAll()

		assert.NoError(t, err)
		assert.Equal(t, len(result), 3)
	})
}

func TestCreate(t *testing.T) {
	setup := setUp(t)
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
		setIntialData()
		u.Email = "test3@example.com"

		_, err := userRepository.Create(u)

		assert.Error(t, err)
		assert.Equal(t, int(err.(*mysql.MySQLError).Number), 1062)
	})
}

func TestFindByEmail(t *testing.T) {
	setup := setUp(t)
	defer setup()

	setIntialData()

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
