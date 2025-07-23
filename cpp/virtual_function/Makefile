
CC=g++
CFLAGS=-Wall

all: main

%.o: %.cpp
	$(CC) $(CFLAGS) -o $@ -c $<

main: main.cpp virtual_function.o
	$(CC) $(CFLAGS) -o $@ $^

.PHONY: clean
clean:
	rm -f main *.o *.so