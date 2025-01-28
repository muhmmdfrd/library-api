package models

type TableFilter struct {
	Keyword string `json:"keyword"`
	Size    int    `json:"size"`
	Index   int    `json:"index"`
}