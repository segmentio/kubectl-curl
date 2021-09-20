package curl

import "strconv"

type File string

func NewFile(value string) *File {
	f := File(value)
	return &f
}

func (f *File) Set(value string) error {
	*f = File(value)
	return nil
}

func (f File) Get() interface{} { return string(f) }
func (f File) String() string   { return string(f) }
func (f File) Type() string     { return "file" }

func fileOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewFile(defval)}
}

type Dir string

func NewDir(value string) *Dir {
	d := Dir(value)
	return &d
}

func (d *Dir) Set(value string) error {
	*d = Dir(value)
	return nil
}

func (d Dir) Get() interface{} { return string(d) }
func (d Dir) String() string   { return string(d) }
func (d Dir) Type() string     { return "dir" }

func dirOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewDir(defval)}
}

type Offset int64

func NewOffset(value int64) *Offset {
	o := Offset(value)
	return &o
}

func (o *Offset) Set(value string) error {
	v, err := strconv.ParseInt(value, 10, 64)
	*o = Offset(v)
	return err
}

func (o Offset) Get() interface{} { return int64(o) }
func (o Offset) String() string   { return formatInt64(int64(o)) }
func (o Offset) Type() string     { return "offset" }

func offsetOption(name, short string, defval int64, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewOffset(defval)}
}

type Bytes int64

func NewBytes(value int64) *Bytes {
	b := Bytes(value)
	return &b
}

func (b *Bytes) Set(value string) error {
	v, err := strconv.ParseInt(value, 10, 64)
	*b = Bytes(v)
	return err
}

func (b Bytes) Get() interface{} { return int64(b) }
func (b Bytes) String() string   { return formatInt64(int64(b)) }
func (b Bytes) Type() string     { return "bytes" }

func bytesOption(name, short string, defval int64, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewBytes(defval)}
}

type FileRange string

func NewFileRange(value string) *FileRange {
	r := FileRange(value)
	return &r
}

func (r *FileRange) Set(value string) error {
	*r = FileRange(value)
	return nil
}

func (r FileRange) Get() interface{} { return string(r) }
func (r FileRange) String() string   { return string(r) }
func (r FileRange) Type() string     { return "range" }

func rangeOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewFileRange(defval)}
}
