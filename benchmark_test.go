package symbol_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ckaznocha/symbol"
)

func BenchmarkNew(b *testing.B) {
	const foo = "foo"
	syms := map[symbol.Symbol]struct{}{}
	for i := 0; i < b.N; i++ {
		sym := symbol.New(foo)
		if v, ok := syms[sym]; ok {
			b.Errorf("New() = %+v, want a unique symbol", v)
		}
		syms[sym] = struct{}{}
	}
}

func BenchmarkSymbol_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sym := symbol.New(strconv.Itoa(i))
		want := fmt.Sprintf("symbol.New(%d)", i)
		if sym.String() != want {
			b.Errorf("Symbol.String() = %s, want %s", sym.String(), want)
		}
	}
}

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := symbol.For("foo")
		got2 := symbol.For("foo")
		if got != got2 {
			b.Errorf("For(\"foo\") = %+v, want %+v", got, got2)
		}
		if got == symbol.For(strconv.Itoa(i)) {
			b.Errorf("For(\"%d\") = %+v, want a unique symbol", i, got)
		}
	}
}
func BenchmarkKeyFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v, ok := symbol.KeyFor(symbol.For(strconv.Itoa(i)))
		if !ok {
			b.Errorf("KeyFor(symbol.For(%d)) = %q, want %q", i, v, i)
		}
		v, ok = symbol.KeyFor(symbol.New(""))
		if ok {
			b.Errorf("KeyFor(symbol.New(\"\")) = %q, want \"\"", v)
		}
	}
}
