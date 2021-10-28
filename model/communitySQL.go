package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type sqliteHandler4 struct {
	db *sql.DB // 멤버변수로 가진다
}

func (s *sqliteHandler4) GetCommunity() []*Community {
	community := []*Community{}                                                                        //list를 만든다
	rows, err := s.db.Query("SELECT board_id, title, content, id, date, file_id, good FROM community") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var Community Community                                                                                                                   //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&Community.Board_Id, &Community.Title, &Community.Content, &Community.ID, &Community.Date, &Community.File_Id, &Community.Good) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		community = append(community, &Community)
	}
	log.Print(community[0])
	return community
}

func (s *sqliteHandler4) AddCommunity(board_id string, title string, content string, id string, date string, file_id string, good string) *Community { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO community (board_id, title, content, id, date, file_id, good) VALUES (?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(board_id, title, content, id, date, file_id, good)
	if err != nil {
		panic(err)
	}
	var Community Community
	Community.Board_Id = board_id
	Community.Title = title
	Community.Content = content
	Community.ID = id
	Community.Date = date
	Community.File_Id = file_id
	Community.Good = good

	return &Community
}

func (s *sqliteHandler4) RemoveCommunity(board_id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := s.db.Prepare("DELETE FROM community WHERE board_id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(board_id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *sqliteHandler4) Close4() {
	s.db.Close()
}

func newSqliteHandler4(filepath string) DBHandler4 {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS community (
			board_id			TEXT PRIMARY KEY,
			title		TEXT ,
			content		TEXT ,
			id		TEXT ,
			date		TEXT ,
			file_id		TEXT ,
			good		TEXT
			);`)

	statement.Exec()
	return &sqliteHandler4{db: database} // &sqliteHandler{}를 반환
}
