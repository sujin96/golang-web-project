package model

type My_data struct {
	Data_id string `json:"data_id"`
	Date    int    `json:"date"`
	Week    int    `json:"week"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
}

type DBHandler3 interface {
	GetMyData() []*My_data
	//AddMyData(data_id string, distance int, date int, week int, month int, year int, time int) *My_data
	Close3() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func NewDBHandler3(filepath string) DBHandler3 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return newSqliteHandler3(filepath)
}
