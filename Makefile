CXX = g++
CXXFLAGS = -std=c++14
AR = /usr/bin/ar
CFLAGS = -I./src
#LDFLAGS = "-lnummer -L./src"

test: src/libnummer.a
	go test

src/nummer.o: src/nummer.cpp
	$(CXX) $(CXXFLAGS) -c src/nummer.cpp -o src/nummer.o

src/libnummer.a: src/nummer.o
	$(AR) rcs src/libnummer.a src/nummer.o

clean:
	rm -f src/nummer.o src/libnummer.a
