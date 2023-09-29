package repositories

type Repositories interface {
	Account() Account
	Message() Messaging
}

type Account interface {
}

type Messaging interface {
}
