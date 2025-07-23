CC=gcc
CFLAGS=-Wall -g

all: server temp

%.o: %.c
	$(CC) $(CFLAGS) $> -o $@

server: server.o %.o
	$(CC) $(CFLAGS) 
