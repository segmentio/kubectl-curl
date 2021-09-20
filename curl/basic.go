package curl

import "strconv"

type Bool bool

func NewBool(value bool) *Bool {
	b := Bool(value)
	return &b
}

func (b *Bool) Set(value string) error {
	v, err := strconv.ParseBool(value)
	*b = Bool(v)
	return err
}

func (b Bool) Get() interface{} { return bool(b) }
func (b Bool) Type() string     { return "" }
func (b Bool) String() string   { return strconv.FormatBool(bool(b)) }
func (b Bool) IsBoolFlag() bool { return true }

func boolOption(name, short string, defval bool, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewBool(defval)}
}

type Binary []byte

func NewBinary(value []byte) *Binary {
	b := Binary(value)
	return &b
}

func (b *Binary) Set(value string) error {
	*b = Binary(value)
	return nil
}

func (b Binary) Get() interface{} { return ([]byte)(b) }
func (b Binary) String() string   { return string(b) }
func (b Binary) Type() string     { return "data" }

func dataOption(name, short string, defval []byte, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewBinary(defval)}
}

type Name string

func NewName(value string) *Name {
	n := Name(value)
	return &n
}

func (n *Name) Set(value string) error {
	*n = Name(value)
	return nil
}

func (n Name) Get() interface{} { return string(n) }
func (n Name) String() string   { return string(n) }
func (n Name) Type() string     { return "name" }

func nameOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewName(defval)}
}

type Number int

func NewNumber(value int) *Number {
	n := Number(value)
	return &n
}

func (n *Number) Set(value string) error {
	v, err := parseInt(value)
	*n = Number(v)
	return err
}

func (n Number) Get() interface{} { return int64(n) }
func (n Number) String() string   { return formatInt(int(n)) }
func (n Number) Type() string     { return "num" }

func numberOption(name, short string, defval int, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewNumber(defval)}
}

type String string

func NewString(value string) *String {
	s := String(value)
	return &s
}

func (s *String) Set(value string) error {
	*s = String(value)
	return nil
}

func (s String) Get() interface{} { return string(s) }
func (s String) String() string   { return string(s) }
func (s String) Type() string     { return "string" }

func stringOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewString(defval)}
}

type Token string

func NewToken(value string) *Token {
	t := Token(value)
	return &t
}

func (t *Token) Set(value string) error {
	*t = Token(value)
	return nil
}

func (t Token) Get() interface{} { return string(t) }
func (t Token) String() string   { return string(t) }
func (t Token) Type() string     { return "token" }

func tokenOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewToken(defval)}
}

type Type string

func NewType(value string) *Type {
	t := Type(value)
	return &t
}

func (t *Type) Set(value string) error {
	*t = Type(value)
	return nil
}

func (t Type) Get() interface{} { return string(t) }
func (t Type) String() string   { return string(t) }
func (t Type) Type() string     { return "type" }

func typeOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewType(defval)}
}

func parseInt(s string) (int, error) {
	i, err := parseInt64(s)
	return int(i), err
}

func formatInt(i int) string {
	return formatInt64(int64(i))
}

func parseInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(s, 10, 64)
}

func formatInt64(i int64) string {
	if i <= 0 {
		return ""
	}
	return strconv.FormatInt(i, 10)
}
