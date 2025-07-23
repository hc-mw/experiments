#include <stdio.h>
#include <alloca.h>

int main()
{
    // thats variable length array
    int n = 5;
    int arr[n];
    // problem is that it is not supported by C89
    // and also it is not supported by C++ standard
    // and it can cause stack overflow
    // so it is better to use alloca
    int *arr2 = alloca(n * sizeof(int));

    return 0;
}