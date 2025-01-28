package models

type PaginationData struct {
	Size     int         `json:"size"`
	Filtered int         `json:"filtered"`
	Total    int         `json:"total"`
	Data     interface{} `json:"data"`
}