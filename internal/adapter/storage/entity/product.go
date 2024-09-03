package entity

// Product structure to describe goods table
type Product struct {
	Id            int64   `json:"id" gorm:"column:id"`
	Name          string  `json:"name" gorm:"column:name"`
	Amount        float64 `json:"amount" gorm:"column:amount"`
	CategoryId    int64   `json:"categoryId" gorm:"column:category_id"`
	SubCategoryId int64   `json:"subCategoryId" gorm:"column:sub_category_id"`
	IsEnabled     bool    `json:"isEnabled" gorm:"column:is_enabled"`
}

func (Product) TableName() string {
	return "public.Product"
}
