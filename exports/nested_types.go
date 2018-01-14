package exports

type NestedTypes struct {
	Id        int32              `json:"Id"`
	DictTest1 map[string]float32 `json:"DictTest1"`
	DictTest2 map[string]int32   `json:"DictTest2"`
	ListTest1 []int32            `json:"ListTest1"`
	ListTest2 []float32          `json:"ListTest2"`
}
