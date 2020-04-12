package models

import (
	"time"
)

// RawStockPriceWrapper - 企業ID・企業名・銘柄コードと共に、株価の生時系列データを保持
type RawStockPriceWrapper struct {

	// クライアントへレスポンスを送信する日時
	Timestamp time.Time `json:"timestamp,omitempty"`

	// システム内で設定する企業毎に一意となるコード
	CompanyID string `json:"companyId,omitempty"`

	// 企業名
	CompanyName string `json:"companyName,omitempty"`

	// 銘柄コード
	StockCode string `json:"stockCode,omitempty"`

	// 株価の時系列データ本体
	Data []RawStockPrice `json:"data,omitempty"`
}

// RawStockPrice - 株価の生データ(open/high/low/close/volume)
type RawStockPrice struct {

	// 日付
	Date string `json:"date,omitempty"`

	// 株価(始値)
	Open float32 `json:"open,omitempty"`

	// 株価(高値)
	High float32 `json:"high,omitempty"`

	// 株価(安値)
	Low float32 `json:"low,omitempty"`

	// 株価(終値)
	Close float32 `json:"close,omitempty"`

	// 出来高
	Volume float32 `json:"volume,omitempty"`
}

// StockPriceMAWrapper - 企業ID・企業名・銘柄コードと共に、各種移動平均株価データを保持
type StockPriceMAWrapper struct {

	// クライアントへレスポンスを送信する日時
	Timestamp time.Time `json:"timestamp,omitempty"`

	// システム内で設定する企業毎に一意となるコード
	CompanyID string `json:"companyId,omitempty"`

	// 企業名
	CompanyName string `json:"companyName,omitempty"`

	// 銘柄コード
	StockCode string `json:"stockCode,omitempty"`

	// 株価の時系列データ本体
	Data []StockPriceMA `json:"data,omitempty"`
}

// StockPriceMA - 移動平均株価データ
type StockPriceMA struct {

	// 日付
	Date string `json:"date,omitempty"`

	// 移動平均の計算方法 (計算方法と期間のセット)
	MAType string `json:"ma_type,omitempty"`

	// 移動平均株価
	MAValue float32 `json:"ma_value,omitempty"`
}
