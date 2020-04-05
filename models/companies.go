package models

import (
	"time"
)

// CompanyWrapper - 企業の基本情報を持つCompanyクラスをリストで保持するWrapperクラス
type CompanyWrapper struct {

	// クライアントへレスポンスを送信する日時
	Timestamp time.Time `json:"timestamp,omitempty"`

	// 企業名・銘柄コード・設立日などの会社の基本情報を保持するクラスの配列
	Companies []Company `json:"companies,omitempty"`
}

// Company - 企業名・銘柄コード・設立日などの会社の基本情報を保持する
type Company struct {

	// システム内で設定する企業毎に一意となるコード
	CompanyID string `json:"companyId,omitempty"`

	// 企業名
	CompanyName string `json:"companyName,omitempty"`

	// 銘柄コード
	StockCode string `json:"stockCode,omitempty"`

	// 国コード
	CountryCode string `json:"countryCode,omitempty"`

	// 上場市場
	ListedMarket string `json:"listedMarket,omitempty"`

	// 設立日(月)
	FoundationDate string `json:"foundationDate,omitempty"`

	// 本社所在地経度
	Longitude float32 `json:"longitude,omitempty"`

	// 本社所在地緯度
	Latitude float32 `json:"latitude,omitempty"`
}

// CompanyDetail - 企業の詳細情報を保持する
// 現時点では上記のCompanyと同等だが、今後詳細情報を応答するAPIと一覧を応答するAPIで
// レスポンスとして使用する構造体を変えるため、分離しておく
type CompanyDetail struct {

	// クライアントへレスポンスを送信する日時
	Timestamp time.Time `json:"timestamp,omitempty"`

	// システム内で設定する企業毎に一意となるコード
	CompanyID string `json:"companyId,omitempty"`

	// 企業名
	CompanyName string `json:"companyName,omitempty"`

	// 銘柄コード
	StockCode string `json:"stockCode,omitempty"`

	// 国コード
	CountryCode string `json:"countryCode,omitempty"`

	// 上場市場
	ListedMarket string `json:"listedMarket,omitempty"`

	// 設立日(月)
	FoundationDate string `json:"foundationDate,omitempty"`

	// 本社所在地経度
	Longitude float32 `json:"longitude,omitempty"`

	// 本社所在地緯度
	Latitude float32 `json:"latitude,omitempty"`
}
