#include <stdio.h>

// static functions and globals means
// they are only visible in this translation unit
// basically private to this file only
// and 
// static local variable means 
// they variable sticks to that function scope
// it doesnt get redefined or destroyed

void inc() {
	static int count = 0;
	count++;
	if (count > 10)
					return;
	printf("[info] count: %d\n", count);
	inc();
}

int main(int argc, const char **argv) {
	inc();
	return 0;
}

