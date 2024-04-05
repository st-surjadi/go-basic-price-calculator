package main

import (
	"fmt"

	"example.com/go-basic-price-calculator/prices"
	"example.com/go-basic-price-calculator/utils"
)

const priceFileName = "prices.txt"

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.15}

	for _, taxRate := range taxRates {
		fm := utils.New(priceFileName, fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		priceJob.Process()
	}

}
