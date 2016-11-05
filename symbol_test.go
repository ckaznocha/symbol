package symbol_test

import (
	"reflect"
	"testing"

	"github.com/ckaznocha/symbol"
)

func TestNew(t *testing.T) {
	syms := map[symbol.Symbol]struct{}{}
	for i := 0; i < 1000000; i++ {
		sym := symbol.New("")
		if v, ok := syms[sym]; ok {
			t.Fatalf("Unique test. New() = %+v, want a unique symbol", v)
		}
		syms[sym] = struct{}{}
	}
}

func TestFor(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want symbol.Symbol
	}{
		{"Returns a symbol if its already in the registry", args{"foo"}, symbol.For("foo")},
	}
	for _, tt := range tests {
		if got := symbol.For(tt.args.key); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. For(%v) = %+v, want %+v", tt.name, tt.args.key, got, tt.want)
		}
	}
}

func TestKeyFor(t *testing.T) {
	type args struct {
		sym symbol.Symbol
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{"Symbol in the registry", args{symbol.For("foo")}, "foo", true},
		{"Symbol not in the registry", args{symbol.New("foo")}, "", false},
	}
	for _, tt := range tests {
		got, got1 := symbol.KeyFor(tt.args.sym)
		if got != tt.want {
			t.Errorf("%q. KeyFor(%+v) got = %v, want %v", tt.name, tt.args.sym, got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("%q. KeyFor(%+v) got1 = %v, want %v", tt.name, tt.args.sym, got1, tt.want1)
		}
	}
}

func TestSymbol_String(t *testing.T) {
	tests := []struct {
		name string
		sym  symbol.Symbol
		want string
	}{
		{"Returns the correct string", symbol.New("foo"), "symbol.New(foo)"},
		{"Returns the correct string", symbol.New(""), "symbol.New()"},
		{"Returns the correct string", symbol.For("foo"), "symbol.New(foo)"},
	}
	for _, tt := range tests {
		if got := tt.sym.String(); got != tt.want {
			t.Errorf("%q. symbol.Symbol.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
