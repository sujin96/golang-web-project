package model

type File struct {
	File_ID  string `json:"file_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Size     string `json:"size"`
}

type DBHandler5 interface {
	GetFile() []*File
	AddFile(file_id string, name string, location string, size string) *File
	RemoveFile(file_id string) bool
	Close5() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func NewDBHandler5(filepath string) DBHandler5 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return newSqliteHandler5(filepath)
}
