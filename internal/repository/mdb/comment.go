package mdb

import (
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Get(urlQuery map[string][]string) ([]models.Comment, error) {
	var comments []models.Comment

	result := r.db.Model(&comments)

	if urlQuery["tour_id"] != nil {
		result.Where("tour_id IN ?", urlQuery["tour_id"])
	}

	if urlQuery["user_id"] != nil && urlQuery["user_id"][0] != "0" {
		result.Where("user_id = ?", urlQuery["user_id"][0])
	}

	// if urlQuery["date"] != nil {
	// 	date := urlQuery["date"][0]

	// 	dateSplit := strings.Split(date, " ")
	// 	if len(dateSplit) != 2 {
	// 		return nil, errors.New("bad request, error split")
	// 	}
	// 	result.Select("comments.*").Joins("left join trips on tours.id = trips.tour_id").Where("trips.start >= ? AND trips.end < ?", dateSplit[0], dateSplit[1]).Group("trips.tour_id")
	// }

	if urlQuery["order_by"] != nil {
		order := urlQuery["order_by"][0]
		result.Order(order)
	}

	if urlQuery["limit"] != nil {
		limit, err := strconv.Atoi(urlQuery["limit"][0])
		if err != nil {
			return nil, err
		}
		result.Limit(limit)
	}

	result.Find(&comments)
	return comments, result.Error
}

func (r *CommentRepository) GetAllCommentsByTourId(tourId int) ([]models.Comment, error) {
	var comments []models.Comment
	result := r.db.Where("tour_id = ?", tourId).Limit(5).Find(&comments)
	return comments, result.Error
}
func (r *CommentRepository) GetOneByID(id int) ([]models.Comment, error) {
	var comments []models.Comment
	result := r.db.Find(&comments, id)
	return comments, result.Error
}
func (r *CommentRepository) Create(comment models.Comment) error {
	result := r.db.Model(&models.Comment{}).Create(&comment)
	return result.Error
}
func (r *CommentRepository) Update(id int, comment models.Comment) error {
	result := r.db.Model(models.Comment{}).Where("id = ?", id).Updates(comment)
	return result.Error
}
func (r *CommentRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Delete(&models.Comment{})
	return result.Error
}
