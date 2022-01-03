package bpl

type BplSection struct {
	Name          string        `json:"name"`
	KeyValuePairs []BplKeyValue `json:"keyValuePairs"`
}