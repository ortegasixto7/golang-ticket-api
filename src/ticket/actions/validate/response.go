package validate

type Response struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Code    string `json:"code"`
	IsValid bool   `json:"is_valid"`
}
