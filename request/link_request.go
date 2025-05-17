package request

type GenerateCode struct {
	Url        string `json:"url"`
	CustomCode string `json:"customCode"`
	ExpireAt   int    `json:"expireAt"`
}

type GetUrl struct {
	Code string `json:"code"`
}
