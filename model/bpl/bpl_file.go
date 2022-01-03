package bpl

type BplFile struct {
	FileName string       `json:"fileName"`
	Sections []BplSection `json:"sections"`
}