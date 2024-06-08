package config

type Config struct {
	Google       Google
	CloudStorage CloudStorage
	Topics       Topics
	Subscribers  Subscribers
	Workers      Workers
	API          API
	Project      Project
	Firestorage  Firestorage
}

type Project struct {
	Id string
}
type Google struct {
	Credentials string
}

type CloudStorage struct {
	FilePathPattern string
	BucketName      string
}

type Topics struct {
	FileUpload string
}

type Subscribers struct {
	FileUpload string
}

type Workers struct {
	NumberOfWorks int
}

type API struct {
	ServerHost string
}

type Firestorage struct {
	StudentCollection string
}
