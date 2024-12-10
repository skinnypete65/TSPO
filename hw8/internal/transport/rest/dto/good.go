package dto

type Good struct {
	ID            string       `json:"id"`
	Name          string       `json:"name" validate:"required"`
	Price         int          `json:"price" validate:"required"`
	Desc          string       `json:"desc" validate:"required"`
	StockQuantity int          `json:"stock_quantity"`
	MeasureUnit   *MeasureUnit `json:"measure_unit" validate:"required"`
}

type MeasureUnit struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,oneof=METER KILOGRAM LITER PIECE"`
}
