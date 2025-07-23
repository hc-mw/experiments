CC=gcc
CFLAGS=-g -Wall

all: showip

showip: showip.o
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

.PHONY: clean
clean:
	rm -f main *.o *.so
