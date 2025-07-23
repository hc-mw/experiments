#include <stdio.h>
#include <stdlib.h>

typedef enum {
  FLAG_A = (1 << 0),
  FLAG_B = (1 << 1),
  FLAG_C = (1 << 2),
  FLAG_D = (1 << 3),
} t_flag;

int ops(int a, t_flag flag) {
  if (flag & FLAG_A)
    a += a;
  if (flag & FLAG_B)
    a -= a;
  if (flag & FLAG_C)
    a *= a;
  if (flag & FLAG_D)
    a /= a;

  return a;
}

int main(int argc, char **argv) {
  int a = 10;

  // try all flags
  printf("a = %d\n", ops(a, FLAG_A));
  printf("a = %d\n", ops(a, FLAG_B));
  printf("a = %d\n", ops(a, FLAG_C));
  printf("a = %d\n", ops(a, FLAG_D));

  // try multiple flags
  printf("a = %d\n", ops(a, FLAG_A | FLAG_B));
  printf("a = %d\n", ops(a, FLAG_A | FLAG_C));
  printf("a = %d\n", ops(a, FLAG_A | FLAG_D));
  printf("a = %d\n", ops(a, FLAG_B | FLAG_C));
  printf("a = %d\n", ops(a, FLAG_B | FLAG_D));
  printf("a = %d\n", ops(a, FLAG_C | FLAG_D));

  return 0;
}