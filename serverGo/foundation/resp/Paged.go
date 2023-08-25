package resp

type Paged[T any] struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	TotalCount  int `json:"totalCount"`
	PageCount   int `json:"pageCount"`
	RowDataList T   `json:"rowDataList"`
}
