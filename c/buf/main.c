#include <stdio.h>
#include <arpa/inet.h>

int main() {
    short n = 1; // Host short integer
    unsigned short network_order = htons(n); // Convert to network byte order

    printf("Host value: %d\n", n);            // Display the original value
    printf("Network byte order: %d\n", (int)network_order); // Display the converted value

    return 0;
}

