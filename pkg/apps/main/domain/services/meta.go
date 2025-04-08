package services

type ResponseMetadata struct {
	Status string
	Count  uint64
	Errors []error
}
