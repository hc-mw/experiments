#include <assert.h>
#include <errno.h>
#include <stdio.h>

void errno_demo(const char *filename);
void perror_demo(const char *filename);
void strerror_demo(const char *filename);
void assert_demo();

int main(int argc, char const *argv[]) {
    // cli app: show list of demos ,let user choose which demo to run, run the
    // chosen demo
    int choice;
    printf("Choose a demo to run:\n");
    printf("1. errno demo\n");
    printf("2. perror demo\n");
    printf("3. assert demo\n");
    scanf("%d", &choice);

    char filename[100];
    switch (choice) {
        case 1:
            // take file name input
            printf("Enter file name: ");
            scanf("%s", filename);
            errno_demo(filename);
            break;
        case 2:
            printf("Enter file name: ");
            scanf("%s", filename);
            perror_demo(filename);
            break;
        case 3:
            assert_demo();
            break;
        default:
            printf("Invalid choice\n");
            break;
    }

    return 0;
}

// errno is a global variable that is set by system calls and some library
// functions in the event of an error to indicate what went wrong. You can see
// full list of erros with 'errno --list'
void errno_demo(const char *filename) {
    errno = 0;
    FILE *f = fopen(filename, "r");
    if (f == NULL) {
        printf("Error opening file: %d\n", errno);
        return;
    }

    char c;
    while ((c = fgetc(f)) != EOF) fputc(c, stdout);
}

// The perror function is used to print an error message to the standard error
// stream (stderr). The error message printed by perror is based on the current
// value of errno.
void perror_demo(const char *filename) {
    FILE *f = fopen(filename, "r");
    if (f == NULL) {
        perror("fopen");
        return;
    }

    char c;
    while ((c = fgetc(f)) != EOF) fputc(c, stdout);
}

// strerror
void strerror_demo(const char *filename) {
    FILE *f = fopen(filename, "r");
    if (f == NULL) {
        printf("Error opening file: %s\n", strerror(errno));
        return;
    }

    char c;
    while ((c = fgetc(f)) != EOF) fputc(c, stdout);
}

// The assert macro is used to check if the condition is true. If the condition
// is false, the program will terminate with an error message. The assert macro
// is useful for debugging purposes.
void assert_demo() {
    int x = 5;
    assert(x == 5);   // This will not terminate the program
    assert(x == 10);  // This will terminate the program

    FILE *f = fopen("error_handling.c", "r");
    assert(f != NULL);

    char c;
    while ((c = fgetc(f)) != EOF) fputc(c, stdout);
}