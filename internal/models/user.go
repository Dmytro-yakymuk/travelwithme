package models

type User struct {
	Id         int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Name       string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Surname    string `gorm:"type:varchar(255); not null" json:"surname" binding:"required"`
	Patronymic string `gorm:"type:varchar(255); not null" json:"patronymic" binding:"required"`
	Email      string `gorm:"type:varchar(255); not null; unique" json:"email" binding:"required"`
	Password   string `gorm:"type:varchar(255); not null" json:"password" binding:"required"`
	Phone      string `gorm:"type:varchar(25); not null" json:"phone"`
	RoleId     int    `gorm:"not null" json:"role_id" binding:"required"`

	Tour    []Tour    `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
	Comment []Comment `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
	Order   []Order   `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type UpdateUserInput struct {
	Surname     string  `json:"surname" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Patronymic  string  `json:"patronymic" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	OldPassword string  `json:"old_password" binding:"required"`
	NewPassword *string `json:"new_password"`
	Phone       *string `json:"phone"`
}

type UserOutput struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}
