package model

import "time"

type ExchangeRate struct {
	Id           uint      `gorm:"primarykey" json:"_id"`
	FromCurrency string    `json:"fromCurrency"`
	ToCurrency   string    `json:"toCurrency"`
	Rate         float64   `json:"rate"`
	Data         time.Time `json:"data"`
}
