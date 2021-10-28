package model

type memoryHandler2 struct {
	Workout_log map[string]*Workout_log
}

func (m *memoryHandler2) GetWorkOutlog() []*Workout_log {
	list := []*Workout_log{}
	for _, v := range m.Workout_log {
		list = append(list, v)
	}
	return list
}

/*
func (m *memoryHandler2) AddWorkOutlog(workout_id string, work_id int, start int, end int, distance int) *Workout_log {
	workout := &Workout_log{workout_id, work_id, start, end, distance}
	m.Workout_log[workout_id] = workout
	return workout
}
*/

func (m *memoryHandler2) Close2() {

}

func newMemoryHandler2() DBHandler2 {
	m := &memoryHandler2{}
	m.Workout_log = make(map[string]*Workout_log)
	return m
}
