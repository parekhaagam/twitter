package storage

type Storage interface {
	GetToken()
	ValidateToken()

}
