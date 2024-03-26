package dto

type AdvertisementDto struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	ImgPth  string  `json:"image_path"`
	Price   float64 `json:"price"`
}
