package database

import (
	"github.com/jinzhu/gorm"
)

// Company Companyテーブルのエンティティモデル
type Company struct {
	Model
	// システム内で設定する企業毎に一意となるコード
	CompanyID string `gorm:"column:company_id;type:varchar(16);not null;primary_key"`

	// 企業名
	CompanyName string `gorm:"column:company_name;type:varchar(255);not null"`

	// 銘柄コード
	StockCode string `gorm:"column:stock_code;type:varchar(16);not null"`

	// 国コード
	CountryCode string `gorm:"column:country_code;type:varchar(16);not null"`

	// 上場市場
	ListedMarket string `gorm:"column:listed_market;type:varchar(32)"`

	// 設立日(月)
	FoundationDate string `gorm:"column:foundation_date;type:timestamp"`

	// 本社所在地経度
	Longitude float32 `gorm:"column:longitude;type:number"`

	// 本社所在地緯度
	Latitude float32 `gorm:"column:latitude;type:number"`
}

// GetCompanies Companiesテーブルから企業情報を取得する
func GetCompanies() ([]*Company, error) {
	var companies []*Company
	// レコードが増えてきたら、クエリパラメータにより1回のレスポンスで返すデータ量を制限する
	err := db.Find(&companies).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return companies, nil
}
