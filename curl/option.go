package curl

import (
	"flag"
	"sort"
)

type OptionSet []Option

func (opts OptionSet) Len() int { return len(opts) }

func (opts OptionSet) Less(i, j int) bool { return opts[i].Name < opts[j].Name }

func (opts OptionSet) Swap(i, j int) { opts[i], opts[j] = opts[j], opts[i] }

func (opts OptionSet) Search(option string) int {
	return sort.Search(len(opts), func(i int) bool {
		return opts[i].Name >= option
	})
}

type Option struct {
	Name  string
	Help  string
	Short string
	Value Value
}

func (opt *Option) String() string {
	if IsBoolFlag(opt.Value) {
		if on, _ := opt.Value.Get().(bool); on {
			return opt.Name
		}
	} else if val := opt.Value.String(); val != "" {
		return opt.Name + "=" + val
	}
	return ""
}

func IsBoolFlag(v Value) bool {
	x, _ := v.(boolFlag)
	return x != nil && x.IsBoolFlag()
}

type Value interface {
	flag.Value
	flag.Getter
	Type() string
}

type boolFlag interface {
	IsBoolFlag() bool
}
