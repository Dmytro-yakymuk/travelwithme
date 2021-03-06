package seeds

import (
	"crypto/sha1"
	"fmt"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

const salt = "skjdhgjSKALdjhgdsnjkdzjhxfdjk"

func CreateUser(db *gorm.DB, name, email, password, role string) error {
	return db.Create(&models.User{Name: name, Email: email, Password: generatePasswordHash(password), Role: role}).Error
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
