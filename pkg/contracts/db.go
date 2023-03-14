package contracts

type DB interface {
	Connect()
	Close()
}
