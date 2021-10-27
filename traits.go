package wxwork

// IBody 可转化为 API 请求体的 trait
type IBody interface {
	intoBody() ([]byte, error)
}
