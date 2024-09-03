package model

// ProductRequest request to create product
type ProductRequest struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
	CategoryId    int64   `json:"categoryId"`
	SubCategoryId int64   `json:"subCategoryId"`
	IsEnabled     bool    `json:"isEnabled"`
}

// ProductResponse response to create product
type ProductResponse struct {
	Id         int64  `json:"id"`
	StatusCode string `json:"statusCode"`
}
