#include <regex.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char ip[16];
    char datetime[32];
    char request[64];
    int status;
    int bytes;
} LogEntry;

void parse_log_entry(const char *log, LogEntry *entry) {
    regex_t regex;
    regmatch_t matches[6];
    // const char *pattern = "^(\\S+) - - \\[(.*?)\\] \"(.*?)\" (\\d+) (\\d+)";
    const char *pattern =
        "^(\\S+) - - \\[([^\\]]+)\\] \"([^\"]+)\" (\\d+) (\\d+)";

    if (regcomp(&regex, pattern, REG_EXTENDED) != 0) {
        fprintf(stderr, "Could not compile regex\n");
        return;
    }

    if (regexec(&regex, log, 6, matches, 0) == 0) {
        // Extract IP
        int len = matches[1].rm_eo - matches[1].rm_so;
        strncpy(entry->ip, log + matches[1].rm_so, len);
        entry->ip[len] = '\0';

        // Extract datetime
        len = matches[2].rm_eo - matches[2].rm_so;
        strncpy(entry->datetime, log + matches[2].rm_so, len);
        entry->datetime[len] = '\0';

        // Extract request
        len = matches[3].rm_eo - matches[3].rm_so;
        strncpy(entry->request, log + matches[3].rm_so, len);
        entry->request[len] = '\0';

        // Extract status
        len = matches[4].rm_eo - matches[4].rm_so;
        char status_str[len + 1];
        strncpy(status_str, log + matches[4].rm_so, len);
        status_str[len] = '\0';
        entry->status = atoi(status_str);

        // Extract bytes
        len = matches[5].rm_eo - matches[5].rm_so;
        char bytes_str[len + 1];
        strncpy(bytes_str, log + matches[5].rm_so, len);
        bytes_str[len] = '\0';
        entry->bytes = atoi(bytes_str);
    } else {
        fprintf(stderr, "No match found\n");
    }

    regfree(&regex);
}

void print_log_entry(const LogEntry *entry) {
    printf("IP: %s\n", entry->ip);
    printf("Datetime: %s\n", entry->datetime);
    printf("Request: %s\n", entry->request);
    printf("Status: %d\n", entry->status);
    printf("Bytes: %d\n", entry->bytes);
}

int main() {
    const char *dummy_log =
        "127.0.0.1 - - [10/Oct/2023:13:55:36 -0700] \"GET /index.html "
        "HTTP/1.1\" 200 2326\n";
    LogEntry entry;

    parse_log_entry(dummy_log, &entry);
    print_log_entry(&entry);

    return 0;
}