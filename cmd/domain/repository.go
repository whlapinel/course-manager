package domain

type Repository[T any] interface {
	Saver[T]
	Aller[T]
	CSVer[T]
}

type Saver[T any] interface {
	Save(*T) error
}

type Aller[T any] interface {
	All() ([]*T, error)
}

type CSVer[T any] interface {
	WriteToCSV(*T) error
	ReadFromCSV() ([]*T, error)
}
