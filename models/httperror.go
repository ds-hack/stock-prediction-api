package models

// HTTPError - HTTP Statusコードとセットでユーザーにメッセージを表示する
type HTTPError struct {

	// HTTP Status Code
	Status int `json:"status,omitempty"`

	// エラーに対する付加的な情報
	Message string `json:"message,omitempty"`
}
