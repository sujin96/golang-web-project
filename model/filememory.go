package model

type memoryHandler5 struct {
	File map[string]*File
}

func (m *memoryHandler5) GetFile() []*File {
	list := []*File{}
	for _, v := range m.File {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler5) AddFile(file_id string, name string, location string, size string) *File {
	File := &File{file_id, name, location, size}
	m.File[file_id] = File
	return File
}

func (m *memoryHandler5) RemoveFile(file_id string) bool {
	if _, ok := m.File[file_id]; ok { // memberMap id 값이 있으면
		delete(m.File, file_id) //지우고
		return true
	}
	return false
}

func (m *memoryHandler5) Close5() {

}

func newMemoryHandler5() DBHandler5 {
	m := &memoryHandler5{}
	m.File = make(map[string]*File)
	return m
}
