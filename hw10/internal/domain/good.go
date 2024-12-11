package domain

type Good struct {
	ID            string `gorm:"column:good_id;primaryKey"`
	Name          string `gorm:"column:good_name"`
	Price         int    `gorm:"column:price"`
	Desc          string `gorm:"column:description"`
	StockQuantity int    `gorm:"column:stock_quantity"`
	MeasureUnit   string `gorm:"column:measure_unit"`
}

func (Good) TableName() string {
	return "goods"
}
