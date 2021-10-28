package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type sqliteHandler5 struct {
	db *sql.DB // 멤버변수로 가진다
}

func (s *sqliteHandler5) GetFile() []*File {
	file := []*File{}                                                        //list를 만든다
	rows, err := s.db.Query("SELECT file_id, name, content, location, size") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var File File                                                    //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&File.File_ID, &File.Name, &File.Location, &File.Size) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		file = append(file, &File)
	}
	log.Print(file[0])
	return file
}

func (s *sqliteHandler5) AddFile(file_id string, name string, location string, size string) *File { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO file (file_id, name, location, size) VALUES (?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(file_id, name, location, size)
	if err != nil {
		panic(err)
	}
	var File File
	File.File_ID = file_id
	File.Name = name
	File.Location = location
	File.Size = size

	return &File
}

func (s *sqliteHandler5) RemoveFile(file_id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := s.db.Prepare("DELETE FROM file WHERE file_id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(file_id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *sqliteHandler5) Close5() {
	s.db.Close()
}

func newSqliteHandler5(filepath string) DBHandler5 {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS file (
			file_id			TEXT PRIMARY KEY,
			name		TEXT NOT NULL,
			location		TEXT NOT NULL,
			size		TEXT NOT NULL
			);`)

	statement.Exec()
	return &sqliteHandler5{db: database} // &sqliteHandler{}를 반환
}
