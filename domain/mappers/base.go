package mappers

type BaseMap[T any] interface {
	ToDomain(data any) (T, error)
	ToPersistence(data T) (any, error)
}
