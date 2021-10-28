package model

type memoryHandler3 struct {
	My_dataMap map[string]*My_data
}

func (m *memoryHandler3) GetMyData() []*My_data {
	list := []*My_data{}
	for _, v := range m.My_dataMap {
		list = append(list, v)
	}
	return list
}

/*
func (m *memoryHandler3) AddMyData(data_id string, totaltime float64, trytime float64, recoverytime float64, frontcount float64, backcount float64, avaregyRPB float64, averagespeed float64, distance float64, musclenum float64, kcalorynum float64, date int, week int, month int, year int) *My_data {
	My_data := &My_data{data_id, totaltime, trytime, recoverytime, frontcount, backcount, avaregyRPB, averagespeed, distance, musclenum, kcalorynum, date, week, month, year}
	m.My_dataMap[data_id] = My_data
	return My_data
}
*/

func (m *memoryHandler3) Close3() {

}

func newMemoryHandler3() DBHandler3 {
	m := &memoryHandler3{}
	m.My_dataMap = make(map[string]*My_data)
	return m
}
