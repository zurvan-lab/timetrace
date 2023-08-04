package src

type Set struct {
	Elements Elements
}

type Sets map[string]Set

func NewSets() *Sets {
	return &Sets{}
}

func (s *Sets) NewSet(key string) {
	(*s)[key] = Set{
		Elements: *NewElements(),
	}
}

func (s *Set) NewElement(key string, v Value) {
	s.Elements.AddElement(key, v)
}
