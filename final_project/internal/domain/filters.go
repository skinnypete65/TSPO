package domain

type GormFilter struct {
	Query  string
	Params []any
}

type GormOrders []string
