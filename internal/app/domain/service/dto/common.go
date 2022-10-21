package dto

type PageParams struct {
  Current int `json:"current"` // 当前页
  Size    int `json:"size"`    // 页的大小
}

type Ids struct {
  Ids []int `json:"ids"` // id 切片
}
