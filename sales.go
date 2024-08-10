package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// data
// id,customer_email,delivery_status,total
// 1,customer@example.com,processing,725.6
// int, string, strin, float  

type Sale struct{
	id				int
	customer_email 	string
	delivery_status string
	total 			float64
}

func (s Sale) Print(){
	fmt.Println("Sale: ", s.id)
	fmt.Println("Customer: ", s.customer_email)
	fmt.Println("Status: ", s.delivery_status)
	fmt.Println("Total: ", s.total)
}

// Function to calculate average response time for successful delivery shipments
func averageStatusDelivery(data []Sale) map[string]float64{
	return map[string]float64{
		data[0].customer_email: data[0].total,
	}
}

// Function to analyze sales metrics
func analyzeSalesMetrics(data []Sale) map[string]float64{
	var totalShipping, totalProcessing, totalPreparing, totalDelivered float64
	for _, sale := range data {
		switch sale.delivery_status {
		case "shipped":
			totalShipping+= sale.total
		case "delivered":
			totalDelivered+= sale.total
		case "processing":
			totalProcessing+= sale.total
		case "preparing":
			totalPreparing+= sale.total
		}
	}
}