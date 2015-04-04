# Fileserver

This project is a file server written in [Go](http://golang.org). The difference
from the [standard library file server](http://golang.org/pkg/net/http/#FileServer)
is that you can specify custom `Error` and `NotFound` handlers to respond with
rather than being stuck using the built-in ones that only respond with text.

_Note, this project is a fork of the Go source for serving files. The difference
is just what was needed to implement the custom handlers._ 

## Go in Practice
This was inspired by the development of the book
[Go in Practice](http://manning.com/butcher/) when writing about file service
and solving some common problems.

## License
The license is the same as [Go itself](https://github.com/golang/go/blob/master/LICENSE).

## Todo
- Move the tests over.