package models

type StartSaleBarcodeReq struct {
	SaleID    string`json:"sale_id"`
	Barcode   string`json:"barcode"`
	Quantity     int`json:"quantity"`
}

type EndSaleResp struct {
	ProductName  string`json:"product_name"`
	ProductQuantity int`json:"product_quantity"`
	ProductPrice    int`json:"product_price"`
	TotalSum        int`json:"total_sum"`
}