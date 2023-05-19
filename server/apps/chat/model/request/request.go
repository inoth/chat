package request

type LoginRequest struct {
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}
