package session

type Repository interface {
	GetByToken(token Token) *Session
	Store(session Session) error
}
