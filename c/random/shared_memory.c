#include <wait.h>
#include <stdio.h>
#include <stdint.h>   // Change here
#include <sys/mman.h> // Required for mmap
#include <unistd.h>
#include <bits/mman-linux.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

#define PAGESIZE 4096

int v = 5;

int main(int argc, char const *argv[])
{
    uint8_t *shared_memory = mmap(NULL, PAGESIZE, PROT_READ | PROT_WRITE, MAP_SHARED | MAP_ANONYMOUS, -1, 0);

    *shared_memory = 34;

    if (fork() == 0)
    {
        v = 80;
        *shared_memory = 15;
    }
    else
    {
        wait(NULL);
    }
    printf("Not shared. %i\n", v);
    printf("Shared. %i\n", *shared_memory);
    return 0;
  Lc LLc
