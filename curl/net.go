package curl

import (
	"strings"
)

type Addr string

func NewAddr(value string) *Addr {
	a := Addr(value)
	return &a
}

func (a *Addr) Set(value string) error {
	*a = Addr(value)
	return nil
}

func (a Addr) Get() interface{} { return string(a) }
func (a Addr) String() string   { return string(a) }
func (a Addr) Type() string     { return "address" }

func addressOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewAddr(defval)}
}

type AddrList []string

func NewAddrList(value []string) *AddrList {
	a := AddrList(value)
	return &a
}

func (a *AddrList) Set(value string) error {
	*a = AddrList(strings.Split(value, ","))
	return nil
}

func (a AddrList) Get() interface{} { return ([]string)(a) }
func (a AddrList) String() string   { return strings.Join(([]string)(a), ",") }
func (a AddrList) Type() string     { return "addresses" }

func addressListOption(name, short string, defval []string, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewAddrList(defval)}
}

type HostPort string

func NewHostPort(value string) *HostPort {
	h := HostPort(value)
	return &h
}

func (h *HostPort) Set(value string) error {
	*h = HostPort(value)
	return nil
}

func (h HostPort) Get() interface{} { return string(h) }
func (h HostPort) String() string   { return string(h) }
func (h HostPort) Type() string     { return "address" }

func hostPortOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewHostPort(defval)}
}

type NetworkInterface string

func NewNetworkInterface(value string) *NetworkInterface {
	i := NetworkInterface(value)
	return &i
}

func (i *NetworkInterface) Set(value string) error {
	*i = NetworkInterface(value)
	return nil
}

func (i NetworkInterface) Get() interface{} { return string(i) }
func (i NetworkInterface) String() string   { return string(i) }
func (i NetworkInterface) Type() string     { return "interface" }

func interfaceOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewNetworkInterface(defval)}
}

type Port string

func NewPort(value string) *Port {
	p := Port(value)
	return &p
}

func (p *Port) Set(value string) error {
	*p = Port(value)
	return nil
}

func (p Port) Get() interface{} { return string(p) }
func (p Port) String() string   { return string(p) }
func (p Port) Type() string     { return "num/range" }

func portOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewPort(defval)}
}

type Protocol string

func NewProtocol(value string) *Protocol {
	p := Protocol(value)
	return &p
}

func (p *Protocol) Set(value string) error {
	*p = Protocol(value)
	return nil
}

func (p Protocol) Get() interface{} { return string(p) }
func (p Protocol) String() string   { return string(p) }
func (p Protocol) Type() string     { return "protocol" }

func protocolOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewProtocol(defval)}
}

type ProtocolList []string

func NewProtocolList(value []string) *ProtocolList {
	p := ProtocolList(value)
	return &p
}

func (p *ProtocolList) Set(value string) error {
	*p = ProtocolList(strings.Split(value, ","))
	return nil
}

func (p ProtocolList) Get() interface{} { return ([]string)(p) }
func (p ProtocolList) String() string   { return strings.Join(([]string)(p), ",") }
func (p ProtocolList) Type() string     { return "protocols" }

func protocolListOption(name, short string, defval []string, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewProtocolList(defval)}
}

type URL string

func NewURL(value string) *URL {
	u := URL(value)
	return &u
}

func (u *URL) Set(value string) error {
	*u = URL(value)
	return nil
}

func (u URL) Get() interface{} { return string(u) }
func (u URL) String() string   { return string(u) }
func (u URL) Type() string     { return "url" }

func urlOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewURL(defval)}
}
