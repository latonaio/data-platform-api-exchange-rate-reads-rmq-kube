package requests

type ExchangeRate struct {
	CurrencyTo				string	`json:"CurrencyTo"`
	CurrencyFrom			string	`json:"CurrencyFrom"`
	ValidityStartDate		string	`json:"ValidityStartDate"`
	ValidityEndDate			string	`json:"ValidityEndDate"`
	ExchangeRate			float32	`json:"ExchangeRate"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
