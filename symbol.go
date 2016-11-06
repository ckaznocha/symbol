// Package symbol provides a Go implementation of JavaScript style symbols (atoms).
package symbol

import "sync"

var registry = &struct {
	sync.RWMutex
	keysToSymbols map[string]Symbol
	symbolKeys    map[Symbol]string
}{
	keysToSymbols: map[string]Symbol{},
	symbolKeys:    map[Symbol]string{},
}

// A Symbol is a unique and immutable data type.
type Symbol struct{ *string }

// New creates a new unique symbol. Description is a string which is useful
// for debugging but not to access the symbol itself.
func New(description string) Symbol { return Symbol{&description} }

// String method returns a string representing the specified Symbol.
func (s Symbol) String() string { return "symbol.New(" + *s.string + ")" }

// For searches for existing symbols in a runtime-wide symbol registry with the
// given key and returns it if found. Otherwise a new symbol gets created in the
// global symbol registry with this key.
func For(key string) Symbol {
	registry.RLock()
	sym, ok := registry.keysToSymbols[key]
	registry.RUnlock()
	if !ok {
		sym = New(key)
		registry.Lock()
		registry.keysToSymbols[key], registry.symbolKeys[sym] = sym, key
		registry.Unlock()
	}
	return sym
}

// KeyFor retrieves a shared symbol key from the global symbol registry for the
// given symbol. Returns the key and true if found, otherwise an empty string
// and false.
func KeyFor(sym Symbol) (string, bool) {
	registry.RLock()
	key, ok := registry.symbolKeys[sym]
	registry.RUnlock()
	return key, ok
}
