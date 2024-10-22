package golang_strings_chain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringsChain_Replace(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		old string
		new string
		n   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StringsChain
	}{
		// TODO: Add test cases.
		{fields: fields{value: "Hello Kitty"},
			args: args{
				old: "Kitty",
				new: "World",
				n:   -1,
			},
			want: NewStringsChain("Hello World")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStringsChain(tt.fields.value)
			if got := s.Replace(tt.args.old, tt.args.new, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}

	chain := NewStringsChain("Hello World")
	value := chain.ReplaceAll("Hello", "Star").ReplaceAll("World", "Kitty").value
	fmt.Println(value)
}
