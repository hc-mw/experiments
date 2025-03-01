CC=gcc
CFLAGS=-g -Wall

all: main

%.o: %.c %.h
	$(CC) $(CFLAGS) -c $^

main: printer.o main.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm *.o main