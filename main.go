package main

import (
	"fmt"

	"example.com/go-basic-price-calculator/prices"
	"example.com/go-basic-price-calculator/utils"
)

const priceFileName = "prices.txt"

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.15}
	doneChs := make([]chan bool, len(taxRates))
	errChs := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		fm := utils.New(priceFileName, fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)

		doneChs[i] = make(chan bool)
		errChs[i] = make(chan error)
		go priceJob.Process(doneChs[i], errChs[i])
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <-errChs[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChs[i]:
		}
	}
}
