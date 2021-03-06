package models

type User struct {
	Id       int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"-"`
	Name     string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(255); not null; unique" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(255); not null" json:"password" binding:"required"`
	Role     string `gorm:"type:varchar(20); not null" json:"role" binding:"required"`

	Tour    []Tour    `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
	Comment []Comment `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
	Order   []Order   `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}
