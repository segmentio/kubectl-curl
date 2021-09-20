package curl

type Method string

func NewMethod(value string) *Method {
	m := Method(value)
	return &m
}

func (m *Method) Set(value string) error {
	*m = Method(value)
	return nil
}

func (m Method) Get() interface{} { return string(m) }
func (m Method) String() string   { return string(m) }
func (m Method) Type() string     { return "method" }

func methodOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewMethod(defval)}
}
