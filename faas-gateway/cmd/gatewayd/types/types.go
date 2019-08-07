package types

type ContainerData struct {
	labels string
	status string
	ip     string
	port   string
}

type userError interface {
	error
	Message() string
}
