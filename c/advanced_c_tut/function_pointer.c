#include <stdio.h>
#include <stdlib.h>

typedef void (*array_map_cb)(const void *, const void *, int);
typedef void (*array_for_each_cb)(const void *, int);

typedef struct task {
  char *name;
  int id;
} task;

void double_number(const void *ele, const void *out, int index) {
  int *ip = (int *)ele;
  int *op = (int *)out;
  *ip *= 2;
  *op = *ip;
}

void assign_id_to_task(const void *ele, const void *out, int index) {
  char **t = (char **)ele;
  task *op = (task *)out;
  op->id = index;
  op->name = *t;
}

void *array_map(const void *arr, int size, int ip_size, int op_size,
                array_map_cb cb) {
  void *res = malloc(size * op_size);

  for (int i = 0; i < size; ++i) {
    const void *ip = arr + i * ip_size;
    const void *op = res + i * op_size;
    cb(ip, op, i);
  }

  return res;
}

void array_for_each(const void *arr, int size, int ip_size,
                    array_for_each_cb cb) {
  for (int i = 0; i < size; ++i) {
    const void *ip = arr + i * ip_size;
    cb(ip, i);
  }
}

int main(int argc, char **argv) {
  int arr[] = {1, 2, 3, 4, 5};
  char *arr2[] = {"task1", "task2", "task3", "task4", "task5"};

  void *res = array_map(arr, 5, sizeof(int), sizeof(int), double_number);
  void *res2 = array_map(arr2, 5, sizeof(char *), sizeof(struct task),
                         assign_id_to_task);

  int *nums = (int *)res;
  task *tasks = (task *)res2;

  for (int i = 0; i < 5; ++i) {
    int *x = nums + i;
    printf("%d\n", *x);
  }

  for (int i = 0; i < 5; ++i) {
    task *t = tasks + i;
    printf("Task %d: %s\n", t->id, t->name);
  }

  return 0;
}
