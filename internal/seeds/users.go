package seeds

import (
	"crypto/sha1"
	"fmt"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

const salt = "skjdhgjSKALdjhgdsnjkdzjhxfdjk"

func CreateUser(db *gorm.DB, name, surname, patronymic, email, password, phone string, role_id int) error {
	return db.Create(&models.User{Name: name, Surname: surname, Patronymic: patronymic, Email: email, Password: generatePasswordHash(password), Phone: phone, RoleId: role_id}).Error
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
