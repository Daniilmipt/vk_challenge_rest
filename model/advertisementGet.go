package model

import (
	"rest/dto"

	"github.com/google/uuid"
)

type AdvertisementGetDto struct {
	Adv       Advertisement
	UserLogin string
}

func (advGet *AdvertisementGetDto) GetToSend(userId uuid.UUID) dto.AdvertisementSendDto {
	adv := dto.AdvertisementDto{Title: advGet.Adv.Title, Content: advGet.Adv.Content, ImgPth: advGet.Adv.ImgPth, Price: advGet.Adv.Price}
	return dto.AdvertisementSendDto{Adv: adv, IsUserAdv: userId == advGet.Adv.UserId}
}
