package database

type set map[string]SubSet

type Sets map[string]set

func NewSets() *Sets {
	return &Sets{}
}

func newSet() *set {
	return &set{}
}

func (s *Sets) AddSet(name string) {
	(*s)[name] = *newSet()
}

func (s set) AddSubSet(name string) {
	(s)[name] = *NewSubSet()
}
