#define _POSIX_C_SOURCE 199309L
#include <stdio.h>
#include <time.h>

int main()
{
    struct timespec ts;

    if (clock_gettime(CLOCK_REALTIME, &ts) == -1)
    {
        perror("clock_gettime");
        return 1;
    }

    printf("Current time: %ld seconds, %ld nanoseconds\n",
           (long)ts.tv_sec, ts.tv_nsec);

    return 0;
}
