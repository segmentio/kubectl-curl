package curl

import "strings"

type Certificate string

func NewCertificate(value string) *Certificate {
	c := Certificate(value)
	return &c
}

func (c *Certificate) Set(value string) error {
	*c = Certificate(value)
	return nil
}

func (c Certificate) Get() interface{} { return string(c) }
func (c Certificate) String() string   { return string(c) }
func (c Certificate) Type() string     { return "certificate[:password]" }

func certificateOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewCertificate(defval)}
}

type CertificateKey string

func NewCertificateKey(value string) *CertificateKey {
	k := CertificateKey(value)
	return &k
}

func (k *CertificateKey) Set(value string) error {
	*k = CertificateKey(value)
	return nil
}

func (k CertificateKey) Get() interface{} { return string(k) }
func (k CertificateKey) String() string   { return string(k) }
func (k CertificateKey) Type() string     { return "key[:password]" }

func keyOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewCertificateKey(defval)}
}

type CipherList []string

func NewCipherList(value []string) *CipherList {
	c := CipherList(value)
	return &c
}

func (c *CipherList) Set(value string) error {
	*c = CipherList(strings.Split(value, ":"))
	return nil
}

func (c CipherList) Get() interface{} { return ([]string)(c) }
func (c CipherList) String() string   { return strings.Join(([]string)(c), ":") }
func (c CipherList) Type() string     { return "list-of-ciphers" }

func cipherListOption(name, short string, defval []string, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewCipherList(defval)}
}
