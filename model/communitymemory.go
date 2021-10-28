package model

type memoryHandler4 struct {
	Community map[string]*Community
}

func (m *memoryHandler4) GetCommunity() []*Community {
	list := []*Community{}
	for _, v := range m.Community {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler4) AddCommunity(board_id string, title string, content string, id string, date string, file_id string, good string) *Community {
	Community := &Community{board_id, title, content, id, date, file_id, good}
	m.Community[board_id] = Community
	return Community
}

func (m *memoryHandler4) RemoveCommunity(board_id string) bool {
	if _, ok := m.Community[board_id]; ok { // memberMap id 값이 있으면
		delete(m.Community, board_id) //지우고
		return true
	}
	return false
}

func (m *memoryHandler4) Close4() {

}

func newMemoryHandler4() DBHandler4 {
	m := &memoryHandler4{}
	m.Community = make(map[string]*Community)
	return m
}
