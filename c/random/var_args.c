#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

int sum_all(int count, ...) {
    printf("%d\n", count);
    va_list ap;
    va_start(ap, count);
    
    int sum = 0;
    
    for (int i = 0; i < count; ++i) {
        int x = va_arg(ap, int);
        sum += x;
     } 
    return sum;
}

int main(void) {
    printf("%d\n", sum_all(2,10,10));
    return 0;
}
