package prices

import (
	"fmt"

	"example.com/go-basic-price-calculator/iomanager"
	"example.com/go-basic-price-calculator/utils"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_price"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := utils.StringsToFloats(lines)
	if err != nil {
		return err
	}
	fmt.Println(prices)
	job.InputPrices = prices

	return nil
}

func (job TaxIncludedPriceJob) Process(doneCh chan bool, errCh chan error) {
	err := job.LoadData()
	if err != nil {
		errCh <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", (price * (1 + job.TaxRate)))
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteJSON(job)

	doneCh <- true
}

func NewTaxIncludedPriceJob(taxRate float64, fm iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
