#include <sys/types.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <sys/socket.h>
#include <netdb.h>

int main(int argc, char** argv) {
  if (argc < 2) {
    fprintf(stderr, "usage: %s <address>\n", argv[0]); 
    exit(-1);
  }

  char *node = argv[1];
  int err;
  struct addrinfo *res, hints, *p;
  memset(&hints, 0, sizeof hints);
  hints.ai_family  = AF_UNSPEC;
  hints.ai_socktype = SOCK_STREAM;

  if ((err = getaddrinfo(node, "https", &hints, &res)) != 0) {
    fprintf(stderr, "getaddrinfo: %s\n", gai_strerror(err)); 
    exit(EXIT_FAILURE);
  }

 for (p = res; p != NULL; p = p->ai_next) {
   int fd = socket(p->ai_family, res->ai_socktype, res->ai_protocol);

   int conErr = connect(fd, p->ai_addr, p->ai_addrlen);
   if (conErr == -1)
     exit(EXIT_FAILURE);
  
   char res[4096 * 4];

   int rcvErr =  recv(fd,(void *) &res, 4096 * 4, 0);
   if (rcvErr == -1)
     exit(EXIT_FAILURE);
  
   printf("response: %s\n", res); 
 } 

 freeaddrinfo(res);
}
