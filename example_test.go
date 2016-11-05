package symbol_test

import (
	"fmt"

	"github.com/ckaznocha/symbol"
)

func Example() {
	var (
		events      = map[symbol.Symbol]func(){}
		fooEvent    = symbol.New("")
		eventCaller = func(sym symbol.Symbol) {
			event, ok := events[sym]
			if !ok {
				fmt.Println("No event registered")
				return
			}
			event()
		}
	)

	events[fooEvent] = func() { fmt.Println("foo") }
	eventCaller(fooEvent)
	eventCaller(symbol.New(""))
	// Output: foo
	// No event registered
}

func ExampleNew() {
	var (
		mySymbol         = symbol.New("")
		mySymbolWithDesc = symbol.New("🌯")
	)
	fmt.Println(mySymbol == mySymbolWithDesc)
	fmt.Println(mySymbolWithDesc == symbol.New("🌯"))
	// Output: false
	// false
}

func ExampleSymbol_String() {
	var (
		mySymbol         = symbol.New("")
		mySymbolWithDesc = symbol.New("🛰")
	)
	fmt.Println(mySymbol.String())
	fmt.Println(mySymbolWithDesc.String())
	// Output: symbol.New()
	// symbol.New(🛰)
}

func ExampleFor() {
	var mySymbol = symbol.For("foo")
	fmt.Println(mySymbol == symbol.For("foo"))
	// Output: true
}

func ExampleKeyFor() {
	var mySymbol = symbol.For("foo")
	fmt.Println(symbol.KeyFor(mySymbol))
	// Output: foo true
}
