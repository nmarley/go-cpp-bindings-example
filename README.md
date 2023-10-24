eha 
ch a
e 
rah er
ahc 
ae r
ar 
a c
a dc  bikin owierr io nip iopnwe pino piowpoi epio 

# Go / C++ bindings example

This is an example of Go code calling to a C++ library with a C wrapper.

## Makefile build + run

```sh
make
```

## Build manually

```sh
(cd src/ && g++ -c nummer.cpp && /usr/bin/ar rcs libnummer.a nummer.o)
go build  # this only ensures it compiles
```


## Test manually

```sh
go test  # actually runs tests, c++ lib must be built
```
