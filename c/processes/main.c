#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

int main(int argc, char* argv[]) {
	pid_t pid = fork();

	if (pid == 0) {
		// Child process
		printf("Child Process executing\n");
		sleep(120);
		exit(0);
	} else if (pid > 0) {
		// Parent Process
		printf("%d: Parent process waiting for child\n", pid);
		int status;
		int c_pid = wait(&status);
		printf("Child process terminated,PID: %d, status: %d\n", c_pid, status);		
	} else {
		// fork faiiled
		perror("Fork Failed");
		return 1;
	}
	return 0;
}

