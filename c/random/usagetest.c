#include <stdlib.h>
#include <string.h>
#include <sys/resource.h>
#include <unistd.h>
#include <stdio.h>

long get_mem_usage()
{
    struct rusage myusage;

    getrusage(RUSAGE_SELF, &myusage);
    return myusage.ru_maxrss;
}

int main()
{
    long baseline = get_mem_usage();

    printf("usage :%ld + %ld\n", baseline, get_mem_usage());
}