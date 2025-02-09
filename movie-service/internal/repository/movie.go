package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"movie-service/internal/model"

	"fmt"
	"os"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (r *Repository) Create(movie *model.Movie) error {
	if err := r.db.Create(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(movie *model.Movie) error {
	if err := r.db.Save(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(movie *model.Movie) error {
	if err := r.db.Delete(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) QueryMovieByID(id uint) (*model.Movie, error) {
	movie := &model.Movie{}
	if err := r.db.Preload("Genres").Preload("Stars").Where("id = ?", id).First(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}
