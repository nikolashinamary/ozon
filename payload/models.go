package payload

type RequestEncode struct {
	Url string `json:"url"`
}

type ResponseEncode struct {
	Hash string `json:"hash"`
}
type ResponseDecode struct {
	Url string `json:"url"`
}
