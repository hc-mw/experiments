#include <stdio.h>
#include <string.h>
#include "libstr.h"

void usage(char *command)
{
    printf("Usage: %s <string>\n", command);
}

int main(int argc, char **argv)
{
    if (argc < 2)
    {
        usage(argv[0]);
        return 0;
    }
    int length = strlen(argv[1]);

    if (length < 1)
        usage(argv[0]);

    char *reversed = reverse(argv[1], length);

    printf("Reversed String:\n%s\n", reversed);

    return 0;
}
