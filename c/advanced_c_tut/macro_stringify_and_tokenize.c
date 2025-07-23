#include <stdio.h>

#define FOO(x, y) printf("%s + %s = %d\n", #x, #y, (x) + (y))

#define PRINT_VAR_NAME(var) printf("%s\n", #var)

#define PREFIX(var_name) mylib_##var_name

#define GENERIC_ADD_FUNC(type)                                                 \
  type add_##type(type x, type y) { return (x) + (y); }

#define PRINTF_LOOP(count, ...)                                                \
  do {                                                                         \
    for (int i = 0; i < count; ++i)                                            \
      printf(__VA_ARGS__);                                                     \
  } while (0)

GENERIC_ADD_FUNC(int)

GENERIC_ADD_FUNC(float)

int main() {
  // 1) print var name
  int first = 10;
  int second = 20;

  FOO(first, second);
  // 2) create var names dynamically
  int PREFIX(size_t) = 10;
  printf("%d\n", mylib_size_t);

  // 3) generics
  float res = add_float(1, 2.2);
  printf("%.2f\n", res);

  // 4) macro variable arguments
  PRINTF_LOOP(3, "hello %d %s\n", 32, "bar");

  return 0;
}
