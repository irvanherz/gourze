package core

import (
	"fmt"
	"log"
	"os"

	"github.com/irvanherz/gourze/modules/course"
	"github.com/irvanherz/gourze/modules/media"
	"github.com/irvanherz/gourze/modules/order"
	"github.com/irvanherz/gourze/modules/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("✅ Database connected successfully!")

	// **AutoMigrate all models**
	err = db.AutoMigrate(&user.User{}, &course.Category{}, &course.Course{}, &course.Chapter{}, &course.CourseUser{}, &media.Media{}, &order.Order{}, &order.OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	fmt.Println("✅ Database migration completed!")
	return db, nil
}
