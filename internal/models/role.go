package models

type Role struct {
	Id   int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"-"`
	Name string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`

	Users []User `gorm:"ForeignKey:RoleId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}
