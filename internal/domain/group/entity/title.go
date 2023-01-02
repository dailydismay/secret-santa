package entity

type Title string

func NewTitle(s string) Title {
	return Title(s)
}
