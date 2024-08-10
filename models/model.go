package models

type Sale struct{
	Id				int
	CustomerEmail 	string
	DeliveryStatus string
	Total 			float64
}

type MarketSummary struct{
	MarketName string `json:"market_name"`
	SalesMetrics map[string]float64 `json:"sales_metrics"`
	MostExpensiveSale Sale `json:"most_expensive_sale"`
	AverageDeliveryTime float64  `json:"average_delivery_time"`
}