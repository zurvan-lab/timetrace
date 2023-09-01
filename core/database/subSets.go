package database

type (
	SubSets map[string]SubSet
	SubSet  []Element
)

func NewSubSets() SubSets {
	return SubSets{}
}

func NewSubSet() *SubSet {
	return &SubSet{}
}

func (db *Database) PushElement(s string, sb string, e Element) {
	r := db.Sets[s][sb]
	db.Sets[s][sb] = append(r, e)
}
