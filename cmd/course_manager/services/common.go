package services


type CreatorService[T any] interface {
	Create(*T) error
}

type GetAllService[T any] interface {
	All() ([]*T, error)
}

type UpdaterService[T any] interface {
	Update(*T) error
}

type CSVService[T any] interface {
	ReadFromCSV() ([]*T, error)
}

type DeleterService[T any] interface {
	Delete(*T) error
}