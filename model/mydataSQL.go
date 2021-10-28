package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type sqliteHandler3 struct {
	db *sql.DB // 멤버변수로 가진다
}

func (s *sqliteHandler3) GetMyData() []*My_data {
	MyData := []*My_data{}                                                         //list를 만든다
	rows, err := s.db.Query("SELECT data_id, date, week, month, year FROM mydata") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var My_data My_data                                                                      //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&My_data.Data_id, &My_data.Date, &My_data.Week, &My_data.Month, &My_data.Year) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		MyData = append(MyData, &My_data)
	}
	log.Print(MyData[0])
	return MyData
}

/*
func (s *sqliteHandler3) AddMyData(data_id string, distance int, date int, week int, month int, year int, time int) *My_data { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO mydata (data_id, totaltime, trytime, recoverytime, frontcount, backcount, avaregyRPB, averagespeed, distance, musclenum, kcalorynum, date, week, month, year) VALUES (?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(data_id, distance, date, week, month, year, time)
	if err != nil {
		panic(err)
	}
	var My_data My_data
	My_data.Data_id = data_id
	My_data.Distance = distance
	My_data.Date = date
	My_data.Week = week
	My_data.Month = month
	My_data.Year = year
	My_data.Time = time

	return &My_data
}
*/

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *sqliteHandler3) Close3() {
	s.db.Close()
}

func newSqliteHandler3(filepath string) DBHandler3 {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS mydata (
			data_id			TEXT PRIMARY KEY,
			date	INTEGER DEFAULT '0',
			week	INTEGER DEFAULT '0',
			month	INTEGER DEFAULT '0',
			year	INTEGER DEFAULT '0'
			);`)

	statement.Exec()
	return &sqliteHandler3{db: database} // &sqliteHandler{}를 반환
}
