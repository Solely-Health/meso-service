package repository

type standardRepository interface {
	Store(domainObject interface{}) error
	Find(identifier interface{}) (*interface{}, error)
}