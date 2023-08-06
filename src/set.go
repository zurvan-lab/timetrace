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

// Get Sets Name
func (s *Sets) GetSetsName() []string {
	keys := make([]string, 0, len(*s))
	for key := range *s {
		keys = append(keys, key)
	}
	return keys
}

// Drop Set
func (s *Sets) DropSet(SetName string) {
	delete(*s, SetName)
}

// Clean Set
func (s *Sets) CleanSet(SetName string) {
	if _, ok := (*s)[SetName]; ok {
		(*s)[SetName] = Set{}
	}
}

// Get Elements of a set
func (s *Set) Get(Key string) Value {
	return s.Elements[Key]
}

// Count elements of a set
func (s *Set) Count() int {
	return len((*s).Elements)
}

// Remove Element from a set
func (s *Set) Remove(Key string) {
	delete((*s).Elements, Key)
}

// Update Element from a set
func (s *Set) Upadte(Key string, NewValue Value) {
	(s.Elements)[Key] = NewValue
}

// Get All Keys
func (s *Set) GetKeys() []string {
	keys := make([]string, 0, len(s.Elements))
	for key := range s.Elements {
		keys = append(keys, key)
	}
	return keys
}

// Get All Values
func (s *Set) GetValues() []Value {
	values := make([]Value, 0, len(s.Elements))

	for v := range s.Elements {
		values = append(values, s.Elements[v])
	}
	return values
}

// TODO: Create Snapshot
// TODO: End Snapshot
// TODO: Find List of Elements from a set
// TODO: Sort With Limit And Offset
// TODO: regex search
// TODO: Update list of elements from a set
// TODO: Remove list of elements from a set
// TODO: handles foreach operation on a set
