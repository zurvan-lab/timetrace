package database

type IDataBase interface {
	SetsMap() Sets

	Connect([]string) string
	AddSet([]string) string
	AddSubSet([]string) string
	PushElement([]string) string
	DropSet([]string) string
	DropSubSet([]string) string
	CleanSet([]string) string
	CleanSets([]string) string
	CleanSubSet([]string) string
	CountSets([]string) string
	CountSubSets([]string) string
	CountElements([]string) string
	GetElements([]string) string
}
