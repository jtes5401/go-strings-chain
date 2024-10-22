package golang_strings_chain

import "strings"

type StringsChain struct {
	value string
}

func NewStringsChain(value string) *StringsChain {
	return &StringsChain{value: value}
}

func (s *StringsChain) Replace(old, new string, n int) *StringsChain {
	s.value = strings.Replace(s.value, old, new, n)
	return s
}

func (s *StringsChain) ReplaceAll(old, new string) *StringsChain {
	s.value = strings.ReplaceAll(s.value, old, new)
	return s
}

func (s *StringsChain) TrimSpace() *StringsChain {
	s.value = strings.TrimSpace(s.value)
	return s
}

func (s *StringsChain) Trim(cutset string) *StringsChain {
	s.value = strings.Trim(s.value, cutset)
	return s
}

func (s *StringsChain) Nothing() *StringsChain {

	return s
}
