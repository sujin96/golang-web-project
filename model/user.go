package model

type Member struct {
	ID       string `json:"id"`
	PSWD     string `json:"pswd"`
	Name     string `json:"name"`
	Birth    string `json:"birth"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Area     string `json:"area"`
	BikeINFO string `json:"bike_info"`
	Career   string `json:"career"`
	Club     string `json:"club"`
}

type DBHandler interface {
	GetMembers() []*Member
	AddMember(id string, pswd string, name string, birth string, gender string, email string, area string, bike_info string, career string, club string) *Member
	RemoveMember(id string) bool
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func NewDBHandler(filepath string) DBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return newSqliteHandler(filepath)
}
