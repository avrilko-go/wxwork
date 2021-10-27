package wxwork

// IBody 可转化为API请求体的trait
type IBody interface {
	intoBody() ([]byte, error)
}
