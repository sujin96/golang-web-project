package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type sqliteHandler2 struct {
	db *sql.DB // 멤버변수로 가진다
}

func (s *sqliteHandler2) GetWorkOutlog() []*Workout_log {
	workoutlog := []*Workout_log{}                                                                                                                                                   //list를 만든다
	rows, err := s.db.Query("SELECT workout_id, totaltime, trytime, recoverytime, frontcount, backcount, avaregyRPB, averagespeed, distance, musclenum, kcalorynum FROM workoutlog") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var workout Workout_log                                                                                                                                                                                                                      //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&workout.WORKOUT_ID, &workout.TotalTime, &workout.Trytime, &workout.RecoveryTime, &workout.FrontCount, &workout.Backcount, &workout.AvaregyRPB, &workout.AverageSpeed, &workout.Distance, &workout.MuscleNum, &workout.KcaloryNum) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		workoutlog = append(workoutlog, &workout)
	}
	return workoutlog
}

/*
func (s *sqliteHandler2) AddWorkOutlog(workout_id string, work_id int, start int, end int, distance int) *Workout_log { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO workout (workout_id, work_id, start, end, distance) VALUES (?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(workout_id, work_id, start, end, distance)
	if err != nil {
		panic(err)
	}
	var workout Workout_log
	workout.WORKOUT_ID = workout_id
	workout.ID = work_id
	workout.Start = start
	workout.End = end
	workout.Distance = distance

	return &workout
}
*/
// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *sqliteHandler2) Close2() {
	s.db.Close()
}

func newSqliteHandler2(filepath string) DBHandler2 {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS workout (
			workout_id			TEXT PRIMARY KEY,
			totaltime		DOUBLE,
			trytime		DOUBLE,
			recoverytime		DOUBLE,
			frontcount		DOUBLE,
			backcount		DOUBLE,
			avaregyRPB		DOUBLE,
			averagespeed	DOUBLE,
			distance	DOUBLE,
			musclenum	DOUBLE,
			kcalorynum	DOUBLE
			);`)

	statement.Exec()
	return &sqliteHandler2{db: database} // &sqliteHandler{}를 반환
}
