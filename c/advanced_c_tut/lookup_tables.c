#include <stdio.h>

// NOTE:
// Any values that can be represented as unsigned integers can be used as keys
// in an array (lookup table) (int, char, enum, etc.)

enum students { JOHN = 1, JANE, JIM, JESSICA, JERRY };

int marks[] = {
    [JOHN] = 80, [JANE] = 90, [JIM] = 85, [JESSICA] = 95, [JERRY] = 75};

int main() {
  printf("John's marks: %d\n", marks[JOHN]);
  printf("Jane's marks: %d\n", marks[JANE]);

  int n = sizeof(marks) / sizeof(marks[0]);
  printf("n: %d\n", n);
  for (int i = 0; i < n; ++i)
    printf("%d ", marks[i]);
  printf("\n");

  return 0;
}