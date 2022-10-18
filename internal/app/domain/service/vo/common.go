package vo

type ListResult struct {
	List    any   `json:"list"`
	Total   int64 `json:"total"`
	Current int   `json:"current"`
	Size    int   `json:"size"`
}
