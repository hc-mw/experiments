#include <stdio.h>
#include <stdlib.h>

void pointerArray()
{
    int *p = malloc(3 * sizeof(int));

    int *q = p;

    for (int i = 0; i < 3; i++)
    {
        *(p + i) = i + 1;
    }

    printf("[ ");
    for (int i = 0; i < 3; i++)
    {
        printf("%d ", *(q + i));
    }
    printf("]\n");
}

// void *error()
// {
//     int x = 100;
//     return (void *)&x;
// }

void vm()
{
    char *p = malloc(100);
    for (int i = 0; i < 100; i++)
    {
        *(p + i) = i;
    }

    while (1)
        ;
}

int main(int argc, char const *argv[])
{
    // int x = 1;

    // printf("size of int: %ld\n", sizeof(x));

    // int *p = &x;

    // printf("starting address of x p: %x\n", p);

    // char *cp = (char *)p;

    // cp += 2;

    // printf("x p + 2: %x\n", cp);

    // *cp = 1;

    // printf("number x: %d\n", x);

    // pointerArray();

    // void *ptr = error();
    // printf("%d\n", *((int *)ptr));
    vm();
    return 0;
}
