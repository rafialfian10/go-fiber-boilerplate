package handlerDisaster

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/repositories"
)

type handlerDisaster struct {
	DisasterRepository repositories.DisasterRepository
}

func HandlerDisaster(disasterRepository repositories.DisasterRepository) *handlerDisaster {
	return &handlerDisaster{disasterRepository}
}

func convertDisasterResponse(disaster *models.Disaster) *dto.DisasterResponse {
	return &dto.DisasterResponse{
		ID:           disaster.ID,
		Title:        disaster.Title,
		Description:  disaster.Description,
		Location:     disaster.Location,
		CategoryID:   disaster.CategoryID,
		Category:     disaster.Category.Category,
		Date:         disaster.Date.Format("2006-01-02"),
		Donate:       disaster.Donate,
		DonateTarget: disaster.DonateTarget,
		Image:        disaster.Image,
		IsTrending:   disaster.IsTrending,
		UserID:       disaster.UserID,
		User: dto.UserResponse{
			ID:       disaster.User.ID,
			FullName: disaster.User.FullName,
			Email:    disaster.User.Email,
			Phone:    disaster.User.Phone,
			Gender:   disaster.User.Gender,
			Address:  disaster.User.Address,
			Image:    disaster.User.Image,
			Role: dto.RoleResponse{
				ID:   disaster.User.Role.ID,
				Role: disaster.User.Role.Role,
			},
		},
	}
}

func convertMultipleDisasterResponse(disasterDatas *[]models.Disaster) *[]dto.DisasterResponse {
	var disasters []dto.DisasterResponse

	for _, t := range *disasterDatas {
		disasters = append(disasters, *convertDisasterResponse(&t))
	}

	return &disasters
}
