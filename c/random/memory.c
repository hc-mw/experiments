#include <stdio.h>
#include <stdint.h>   // Change here
#include <sys/mman.h> // Required for mmap
#include <unistd.h>
#include <bits/mman-linux.h>
#include <stdlib.h>
#include <string.h>
#define PAGESIZE 4096

void sbrkDemo();
void mmapDemo();
void memoryAlloc();
void copyMem();
void memoryMappedIO();

int main(int argc, char const *argv[])
{
    mmapDemo();
    sbrkDemo();
    copyMem();
    return 0;
}

void mmapDemo()
{
    printf("-----------------mmap-------------------\n");
    uint8_t *first = mmap(NULL, PAGESIZE, PROT_READ | PROT_WRITE, MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
    uint8_t *second = mmap(NULL, PAGESIZE, PROT_READ | PROT_WRITE, MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);

    printf("first: %p\n", first);
    printf("second: %p\n", second);
}

void sbrkDemo()
{
    printf("-----------------sbrk-------------------\n");
    void *first = sbrk(0);
    void *second = sbrk(0xbc);
    void *third = sbrk(0);

    printf("First: %p\n", first);
    printf("second: %p\n", second);
    printf("third : %p\n", third);
}

void memoryAlloc()
{
    printf("-----------------memory allocation-------------------\n");

    int *arr = malloc(sizeof(int) * 100); // give me space for 100 ints

    int *arr1 = calloc(sizeof(int), 100); // same as above

    arr = realloc(arr1, sizeof(int) * 500); // reallocates, returns new address, will copy old contents to new content

    free(arr);

    free(arr1);
}

void copyMem()
{
    printf("-----------------copy memory-------------------\n");
    typedef struct node
    {
        unsigned int isValid : 1;
        unsigned int size : 15;
        struct node *next;
    } node_t;

    node_t n = {.isValid = 1, .size = 367, .next = NULL};
    node_t n2 = n; // u can copy structs like this
    // another way is to use, memcpy
    memcpy(&n2, &n, sizeof(node_t)); // copies n bytes from source address to dest address

    const int BUFFER_SIZE = 4096;
    const int NODE_OFFSET = 10;

    char buffer[BUFFER_SIZE];

    memcpy(buffer + NODE_OFFSET, &n, sizeof(n));

    node_t *n3 = (node_t *)(buffer + NODE_OFFSET);

    // memset(buffer, 0, BUFFER_SIZE); // sets all bytes of buffer to zero

    printf("n3 = %d, %d, %p\n", n3->isValid, n3->size, n3->next);
}

void memoryMappedIO()
{
    //     int fd = open("./file.txt", O_RDWR, S_IRUSR | S_IWUSR);
    //     struct stat sb;
}