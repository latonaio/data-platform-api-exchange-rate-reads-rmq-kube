package dpfm_api_output_formatter

import (
	"data-platform-api-exchange-rate-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToExchangeRate(rows *sql.Rows) (*[]ExchangeRate, error) {
	defer rows.Close()
	exchangeRate := make([]ExchangeRate, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ExchangeRate{}

		err := rows.Scan(
			&pm.CurrencyTo,
			&pm.CurrencyFrom,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.ExchangeRate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &exchangeRate, nil
		}

		data := pm
		exchangeRate = append(exchangeRate, ExchangeRate{
			CurrencyTo:				data.CurrencyTo,
			CurrencyFrom:			data.CurrencyFrom,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			ExchangeRate:			data.ExchangeRate,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &exchangeRate, nil
}
