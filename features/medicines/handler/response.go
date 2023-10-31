package handler

type InputResponse struct {
	CategoryID   uint   `json:"category_id" form:"category_id"`
	Name         string `json:"name" form:"name"`
	StockMinimum int    `json:"stock_minimum" form:"stock_minimum"`
	Stock        int    `json:"stock" form:"stock"`
	Price        int    `json:"price" form:"price"`
	Photo        string `json:"photo" form:"photo"`
	File         string `json:"file" form:"file"`
}
