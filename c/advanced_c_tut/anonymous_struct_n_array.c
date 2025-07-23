#include <stdio.h>

typedef struct Vec2 {
    int x;
    int y;
} Vec2;

int scale_sum(int scalar, Vec2 v) { return scalar * v.x + scalar * v.y; }

int sum(int arr[], int len) {
    int s = 0;
    for (int i = 0; i < len; ++i) s += arr[i];
    return s;
}

int main() {
    int r = scale_sum(4, (struct Vec2){2, 3});
    int s = sum((int[]){1, 2, 3, 4, 5}, 5);
    printf("%d, %d\n", r, s);
    return 0;
}
