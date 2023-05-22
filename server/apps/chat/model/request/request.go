package request

type LoginRequest struct {
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}

type RoomRequest struct {
	RoomName  string `json:"roomName"`
	UserLimit int    `json:"userLimit"`
	Passwd    string `json:"passwd"`
	Desc      string `json:"desc"`
}

type RoomQueryRequest struct {
	RoomId string `query:"roomId"`
}
