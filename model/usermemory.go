package model

type memoryHandler struct {
	memberMap map[string]*Member
}

//4개 func을 만든다
func (m *memoryHandler) GetMembers() []*Member {
	list := []*Member{}
	for _, v := range m.memberMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddMember(id string, pswd string, name string, birth string, gender string, email string, area string, bike_info string, career string, club string) *Member {
	member := &Member{id, pswd, name, birth, gender, email, area, bike_info, career, club}
	m.memberMap[id] = member
	return member
}

/*
func (m *memoryHandler) UpdateMember(id string, pswd string, email string, area string, bike_info string, career string, club string) *Member {
	member := Member{}
	if _, ok := m.memberMap[id]; ok { // memberMap id 값이 있으면
		delete(m.memberMap, id) //지우고
		return member
	}
	return
}
*/
func (m *memoryHandler) RemoveMember(id string) bool {
	if _, ok := m.memberMap[id]; ok { // memberMap id 값이 있으면
		delete(m.memberMap, id) //지우고
		return true
	}
	return false
}

func (m *memoryHandler) Close() {

}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.memberMap = make(map[string]*Member) // map을 초기화
	return m
}
