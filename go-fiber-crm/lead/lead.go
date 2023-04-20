package lead


import(

)

type Lead struct{
	gorm.Model
	Name string
	Company string
	email string
	Phone int
}