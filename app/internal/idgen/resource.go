package idgen

type Resource interface {
	Prefix() string
}

type resource struct{}

func (r resource) Prefix() string { return "" }

type Request struct{}

func (r Request) Prefix() string { return "req" }

type User struct{}

func (r User) Prefix() string { return "usr" }
