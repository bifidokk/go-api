package fixtures

import (
	"fmt"
	"log"

	"github.com/bifidokk/go-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserMap map[string]entity.User

var UserFixtures = UserMap{
	"user@test.com": {
		Email:    "user@test.com",
		Password: "123456!",
	},
	"user2@test.com": {
		Email:    "user2@test.com",
		Password: "123456!",
	},
}

func CreateUserFixtures(db *gorm.DB) {
	fmt.Println("Create user fixtures")

	keys := make([]string, 0, len(UserFixtures))

	for k := range UserFixtures {
		keys = append(keys, k)
	}

	for _, key := range keys {
		entity := UserFixtures[key]
		password, err := bcrypt.GenerateFromPassword(
			[]byte(entity.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			log.Println("Error occurred during user fixture password generation ", err)
			break
		}

		entity.Password = string(password)

		db.Create(&entity)
	}
}
