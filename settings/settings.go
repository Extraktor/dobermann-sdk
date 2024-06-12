package settings

type Setup struct {
	Token       string
	Environment Environment
}

type Environment int

const (
	Sandbox Environment = iota
	Production
)

func (e Environment) String() string {
	return [...]string{"https://sandbox.asaas.com/api", "https://api.asaas.com"}[e-1]
}

type Method int

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

func (m Method) String() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[m-1]
}
