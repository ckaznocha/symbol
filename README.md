# symbol

[![Build Status](https://travis-ci.org/ckaznocha/symbol.svg?branch=master)](https://travis-ci.org/ckaznocha/symbol)
[![Coverage Status](https://coveralls.io/repos/github/ckaznocha/symbol/badge.svg?branch=master)](https://coveralls.io/github/ckaznocha/symbol?branch=master)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://ckaznocha.mit-license.org)
[![GoDoc](https://godoc.org/github.com/ckaznocha/symbol?status.svg)](https://godoc.org/github.com/ckaznocha/symbol)
[![Go Report Card](https://goreportcard.com/badge/ckaznocha/symbol)](https://goreportcard.com/report/ckaznocha/symbol)

--
    import "github.com/ckaznocha/symbol"

A Go implementation of [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol) style [symbols](https://en.wikipedia.org/wiki/Symbol_%28programming%29) (atoms).

For more info see the example files and the [GoDoc](https://godoc.org/github.com/ckaznocha/symbol) page.

## Usage

#### func  KeyFor

```go
func KeyFor(sym Symbol) (string, bool)
```
KeyFor retrieves a shared symbol key from the global symbol registry for the
given symbol. Returns the key and true if found, otherwise an empty string and
false.

#### type Symbol

```go
type Symbol struct {}
```

A Symbol is a unique and immutable data type.

#### func  For

```go
func For(key string) Symbol
```
For searches for existing symbols in a runtime-wide symbol registry with the
given key and returns it if found. Otherwise a new symbol gets created in the
global symbol registry with this key.

#### func  New

```go
func New(description string) Symbol
```
New creates a new unique symbol. Description is a string which is useful for
debugging but not to access the symbol itself.

#### func (Symbol) String

```go
func (s Symbol) String() string
```
String method returns a string representing the specified Symbol.

## Contributing

See the `CONTRIBUTING` file.

## License
See `LICENSE` file
