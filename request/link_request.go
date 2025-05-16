package request

type GetCode struct {
	Url        string `json:"url"`
	CustomCode string `json:"customCode"`
	ExpireAt   int    `json:"expireAt"`
}
