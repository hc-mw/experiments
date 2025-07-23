#include <stdio.h>
#include <sys/socket.h>

int main(int argc, char const *argv[])
{
    // int fd = socket(AF_INET, SOCK_STREAM, 0);

    // short num = 1;

    // printf("%d\n", htons(num));

    // printf("%d\n", fd);

    int t;

    scanf("%d", &t);

    while (t--)
    {
        printf("Hello\n");
    }

    return 0;
}
