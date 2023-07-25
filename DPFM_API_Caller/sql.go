package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-exchange-rate-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-exchange-rate-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var exchangeRate *[]dpfm_api_output_formatter.ExchangeRate
	for _, fn := range accepter {
		switch fn {
		case "ExchangeRate":
			func() {
				exchangeRate = c.ExchangeRate(mtx, input, output, errs, log)
			}()
		case "ExchangeRates":
			func() {
				exchangeRate = c.ExchangeRates(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ExchangeRate:     exchangeRate,
	}

	return data
}

func (c *DPFMAPICaller) ExchangeRate(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ExchangeRate {
	where := fmt.Sprintf("WHERE CurrencyTo = '%s'", input.ExchangeRate.CurrencyTo)
	where = fmt.Sprintf("%s\nAND CurrencyFrom = '%s'", where, *input.ExchangeRate.CurrencyFrom)
	where = fmt.Sprintf("%s\nAND ValidityStartDate = '%s'", where, *input.ExchangeRate.ValidityStartDate)
	where = fmt.Sprintf("%s\nAND ValidityEndDate = '%s'", where, *input.ExchangeRate.ValidityEndDate)

	if input.ExchangeRate.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.ExchangeRate.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_exchange_rate_exchange_rate_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, CurrencyTo DESC, CurrencyFrom DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToExchangeRate(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ExchangeRates(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ExchangeRate {

	if input.ExchangeRate.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.ExchangeRate.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_exchange_rate_exchange_rate_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, CurrencyTo DESC, CurrencyFrom DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToExchangeRate(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
