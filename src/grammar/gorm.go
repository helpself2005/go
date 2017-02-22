package main

import (
	//"fmt"
	"time"
	//"./gorm_model"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num      int    `gorm:"AUTO_INCREMENT"`

	//CreditCard CreditCard // One-To-One relationship (has one - use CreditCard's UserID as foreign key)
	//Emails []Email // One-To-Many relationship (has many - use Email's UserID as foreign key)

	//BillingAddress   Address // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	//BillingAddressID sql.NullInt64

	//ShippingAddress   Address // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int

	IgnoreMe int `gorm:"-"` // Ignore this field
	//Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:root@/wrc?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{}, &User{})
	/*
		// Create
		db.Create(&Product{Code: "L1212", Price: 1000})

		// Read
		var product Product
		db.First(&product, 1)                   // find product with id 1
		db.First(&product, "code = ?", "L1212") // find product with code l1212

		// Update - update product's price to 2000
		db.Model(&product).Update("Price", 2000)
		fmt.Println(&product)

		// Delete - delete product
		db.Delete(&product)
		db.DropTableIfExists("products")
	*/
}
