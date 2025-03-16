package core

import (
	"fmt"
	"log"

	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/modules/course"
	"github.com/irvanherz/gourze/modules/media"
	"github.com/irvanherz/gourze/modules/order"
	"github.com/irvanherz/gourze/modules/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvideDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.Database.Host,
		config.Database.User,
		config.Database.Pass,
		config.Database.Name,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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
