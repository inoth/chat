package request

type LoginRequest struct {
	UserName string `json:"username" yaml:"username"`
	Passwd   string `json:"passwd" yaml:"passwd"`
}
