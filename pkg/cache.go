package bridge

type Cache interface {
	Record(object interface{})
}
