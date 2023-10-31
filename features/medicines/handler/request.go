package handler

import "mime/multipart"

type InputRequest struct {
	CategoryID   uint           `json:"category_id" form:"category_id"`
	Name         string         `json:"name" form:"name"`
	StockMinimum int            `json:"stock_minimum" form:"stock_minimum"`
	Stock        int            `json:"stock" form:"stock"`
	Price        int            `json:"price" form:"price"`
	Photo        multipart.File `json:"photo,omitempty" form:"photo,omitempty"`
	File         multipart.File `json:"file,omitempty" form:"file,omitempty"`
}

type UpdateRequest struct {
	CategoryID   uint   `json:"category_id" form:"category_id"`
	Name         string `json:"name" form:"name"`
	StockMinimum int    `json:"stock_minimum" form:"stock_minimum"`
	Stock        int    `json:"stock" form:"stock"`
	Price        int    `json:"price" form:"price"`
}
