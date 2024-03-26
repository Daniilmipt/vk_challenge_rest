package model

import (
	"errors"
	constan "rest/constant"
	"rest/dto"
	"time"

	"github.com/google/uuid"
)

type Advertisement struct {
	ID          uuid.UUID `json:"id"`
	CreatedDtTm time.Time `json:"dateTime"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ImgPth      string    `json:"image_path"`
	Price       float64   `json:"price"`
}

func (adv *Advertisement) CreateId() {
	adv.ID = uuid.New()
}

func (adv *Advertisement) ModelToDto() dto.AdvertisementDto {
	return dto.AdvertisementDto{Content: adv.Content, ImgPth: adv.ImgPth, Price: adv.Price, Title: adv.Title}
}

func ModelToDtoArray(advs []Advertisement) []dto.AdvertisementDto {
	advsDto := make([]dto.AdvertisementDto, len(advs))
	for i, adv := range advs {
		advsDto[i] = adv.ModelToDto()
	}
	return advsDto
}

func (adv *Advertisement) FilterFields() error {
	if len(adv.Title) >= constan.TITLE_MAX_LEN || len(adv.Title) <= constan.TITLE_MIN_LEN {
		return errors.New("incorrect title length")
	}
	if len(adv.Content) >= constan.CONTENT_MAX_LEN || len(adv.Content) <= constan.CONTENT_MIN_LEN {
		return errors.New("incorrect content length")
	}
	if adv.Price >= constan.PRICE_MAX || adv.Price <= constan.PRICE_MIN {
		return errors.New("incorrect price value")
	}
	return nil
}
