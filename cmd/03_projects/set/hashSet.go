package set

type hashSet struct {
	set map[interface{}]bool
}

func NewHashSet() Set {
	return &hashSet{set: make(map[interface{}]bool)}
}

func (m *hashSet) Add(e ...interface{}) {
	for _, v := range e {
		m.set[v] = true
	}
}

func (m *hashSet) Clean() {
	m.set = make(map[interface{}]bool)
}

func (m *hashSet) Contains(i ...interface{}) bool {
	contains := true
	for _, v := range i {
		if _, ok := m.set[v]; ok == false {
			contains = false
			break
		}
	}
	return contains
}

func (m *hashSet) Iterator() chan interface{} {
	c := make(chan interface{})
	go func() {
		for k := range m.set {
			c <- k
		}

		close(c)
	}()
	return c
}

func (m *hashSet) Length() int {
	return len(m.set)
}

func (m *hashSet) ToSlice() []interface{} {
	s := make([]interface{}, 0, m.Length())
	for k := range m.set {
		s = append(s, k)
	}
	return s
}
