package domain

type PaginateMeta struct {
	Count  int `json:"count"`
	Starts int `json:"starts"`
	Ends   int `json:"ends"`
}
