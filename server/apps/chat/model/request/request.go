package request

type LoginRequest struct {
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}

type RoomRequest struct {
	Rid    string `json:"rid"`
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Passwd string `json:"passwd"` // default: nil
	Desc   string `json:"desc"`
	Tags   string `json:"tags"`
}

type RoomQueryRequest struct {
	Rid  string `query:"rid"`
	Name string `query:"name"`
	Page
}

type Page struct {
	Index int `query:"i"`
	Size  int `query:"s"`
}
