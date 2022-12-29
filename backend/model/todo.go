package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID 			uint	`gorm:"primaryKey";"AUTO_INCREMENT"` 
	Title		string	`gorm:"size:255"`
	Description string 
	Status 		string
	Priority 	uint
	ProjectID	uint
	Project		Project 
}

type Project struct {
	gorm.Model
	ID 				uint 	`gorm:"primaryKey";"AUTO_INCREMENT"`
	Project_name 	string 
}

type User struct {
	gorm.Model
	ID 			uint 	`gorm:"primaryKey";"AUTO_INCREMENT"`
	Full_name 	string 
	User_id 	*string `gorm:"unique;not null"`
	Email 		string 	`gorm:"unique"`
}

type UserProject struct {
	gorm.Model
	ID 			uint 	`gorm:"primaryKey";"AUTO_INCREMENT"`
	ProjectID	uint
	Project 	Project
	UserID		string
	User	 	User
}