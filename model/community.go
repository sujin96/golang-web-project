package model

type Community struct {
	Board_Id string `json:"board_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ID       string `json:"id"`
	Date     string `json:"date"`
	File_Id  string `json:"file_id"`
	Good     string `json:"good"`
}

type DBHandler4 interface {
	GetCommunity() []*Community
	AddCommunity(board_id string, title string, content string, id string, date string, file_id string, good string) *Community
	RemoveCommunity(board_id string) bool
	Close4() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func NewDBHandler4(filepath string) DBHandler4 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return newSqliteHandler4(filepath)
}
