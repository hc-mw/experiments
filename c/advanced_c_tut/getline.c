#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>

void read_file_line_by_line(const char *filename);

int main() {
    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    printf("Please enter a line of text:\n");

    read = getline(&line, &len, stdin);

    if (read != -1) {
        printf("You entered: %s", line);
    } else {
        printf("Error reading line\n");
    }

    free(line);
    return 0;
}

void read_file_line_by_line(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        perror("Error opening file");
        return;
    }

    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    while ((read = getline(&line, &len, file)) != -1) {
        printf("Retrieved line of length %zu:\n", read);
        printf("%s", line);
    }

    free(line);
    fclose(file);
}