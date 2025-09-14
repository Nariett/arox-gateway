package models

type Image struct {
	Id        int64  `json:"id"`
	IdProduct int64  `json:"id_product"`
	Url       string `json:"url"`
	IsMain    bool   `json:"is_main"`
	IsActive  bool   `json:"is_active"`
}
