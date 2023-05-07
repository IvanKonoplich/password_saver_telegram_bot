package usecase

type Storage interface {
	Set() error
	Get() (string, error)
	Del() error
}

type UseCase struct {
	st Storage
}

func New(st Storage) *UseCase {
	return &UseCase{
		st: st,
	}
}
