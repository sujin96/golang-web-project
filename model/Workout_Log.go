package model

type Workout_log struct {
	WORKOUT_ID   string  `json:"workout_id"`
	TotalTime    float64 `json:"totaltime"`
	Trytime      float64 `json:"trytime"`
	RecoveryTime float64 `json:"recoverytime"`
	FrontCount   float64 `json:"frontcount"`
	Backcount    float64 `json:"backcount"`
	AvaregyRPB   float64 `json:"avaregyRPB"`
	AverageSpeed float64 `json:"averagespeed"`
	Distance     float64 `json:"distance"`
	MuscleNum    float64 `json:"musclenum"`
	KcaloryNum   float64 `json:"kcalorynum"`
}

type DBHandler2 interface {
	GetWorkOutlog() []*Workout_log
	//AddWorkOutlog(workout_id string, work_id int, start int, end int, distance int) *Workout_log
	Close2() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func NewDBHandler2(filepath string) DBHandler2 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return newSqliteHandler2(filepath)
}
