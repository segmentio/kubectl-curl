package curl

type Connect string

func NewConnect(value string) *Connect {
	c := Connect(value)
	return &c
}

func (c *Connect) Set(value string) error {
	*c = Connect(value)
	return nil
}

func (c Connect) Get() interface{} { return string(c) }
func (c Connect) String() string   { return string(c) }
func (c Connect) Type() string     { return "HOST1:PORT1:HOST2:PORT2" }

func connectOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewConnect(defval)}
}

type HeaderField string

func NewHeaderField(value string) *HeaderField {
	h := HeaderField(value)
	return &h
}

func (h *HeaderField) Set(value string) error {
	*h = HeaderField(value)
	return nil
}

func (h HeaderField) Get() interface{} { return string(h) }
func (h HeaderField) String() string   { return string(h) }
func (h HeaderField) Type() string     { return "header/@file" }

func headerOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewHeaderField(defval)}
}

type ProxyAddr string

func NewProxyAddr(value string) *ProxyAddr {
	a := ProxyAddr(value)
	return &a
}

func (a *ProxyAddr) Set(value string) error {
	*a = ProxyAddr(value)
	return nil
}

func (a ProxyAddr) Get() interface{} { return string(a) }
func (a ProxyAddr) String() string   { return string(a) }
func (a ProxyAddr) Type() string     { return "[protocol://]host[:port]" }

func proxyAddrOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewProxyAddr(defval)}
}

type UserPassword string

func NewUserPassword(value string) *UserPassword {
	u := UserPassword(value)
	return &u
}

func (u *UserPassword) Set(value string) error {
	*u = UserPassword(value)
	return nil
}

func (u UserPassword) Get() interface{} { return string(u) }
func (u UserPassword) String() string   { return string(u) }
func (u UserPassword) Type() string     { return "user:password" }

func userOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewUserPassword(defval)}
}
