package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type sqliteHandler struct {
	db *sql.DB // 멤버변수로 가진다
}

func (s *sqliteHandler) GetMembers() []*Member {
	members := []*Member{}                                                                      //list를 만든다
	rows, err := s.db.Query("SELECT id, pswd, name, birth, gender, email, mobile FROM members") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var member Member                                                                                               //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&member.ID, &member.PSWD, &member.Name, &member.Birth, &member.Gender, &member.Email, &member.Mobile) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		members = append(members, &member)
	}
	log.Print(members[0])
	return members
}

func (s *sqliteHandler) AddMember(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Member { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO members (id, pswd, name, birth, gender, email, mobile) VALUES (?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id, pswd, name, birth, gender, email, mobile)
	if err != nil {
		panic(err)
	}
	var member Member
	member.ID = id
	member.PSWD = pswd
	member.Name = name
	member.Birth = birth
	member.Gender = gender
	member.Email = email
	member.Mobile = mobile
	return &member
}

func (s *sqliteHandler) RemoveMember(id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := s.db.Prepare("DELETE FROM members WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS members (
			id			TEXT PRIMARY KEY,
			pswd		TEXT NOT NULL,
			name		TEXT NOT NULL,
			birth		DATE NOT NULL,
			gender		TEXT NOT NULL,
			email		TEXT,
			mobile		TEXT NOT NULL
			)`)

	statement.Exec()
	return &sqliteHandler{db: database} // &sqliteHandler{}를 반환
}
