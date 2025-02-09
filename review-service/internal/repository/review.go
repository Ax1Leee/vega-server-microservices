package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"review-service/internal/model"

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

func (r *Repository) Create(review *model.Review) error {
	if err := r.db.Create(review).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(review *model.Review) error {
	if err := r.db.Save(review).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(review *model.Review) error {
	if err := r.db.Delete(review).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) QueryReviewByID(id uint) (*model.Review, error) {
	review := &model.Review{}
	if err := r.db.Where("id = ?", id).First(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *Repository) QueryReviewByUserIDAndMovieID(userID uint, movieID uint) (*model.Review, error) {
	review := &model.Review{}
	if err := r.db.Where("user_id = ? AND movie_id = ?", userID, movieID).First(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}
