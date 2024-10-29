package domain

type Good struct {
	ID            string
	Name          string
	Price         int
	Desc          string
	StockQuantity int
	MeasureUnit   MeasureUnit
}

type MeasureUnit struct {
	ID   string
	Name string
}
