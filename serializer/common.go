package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
type DataList struct {
	Item  interface{}
	Total uint
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Data: DataList{
			Item:  items,
			Total: total,
		},
	}
}
