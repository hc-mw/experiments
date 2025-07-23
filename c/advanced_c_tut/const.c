// const
#include <stdio.h>

int main(int argc, const char **argv) {
	// const int means, contents of a cannot be reassigned.
	// but you can assign its memory address to other pointer and
	// derefence it and assign it
	const int a = 1;
	// a = 2; // this will throw error
	// take its pointer
	int *pa = (int*)&a;
	// derefernce pointer and assign another value
	*pa = 2;
	printf("a: %d\n", a);

	// Now, we can also have const pointer, which defines that
	// memory address cannot be reassigned, but you can directly
	// assign another value
	int c = 10, d = 15;
	int * const e = &c;
	// assign another value to pointer
	*e = d;
	// e = &d; // this will throw error
	printf("e: %d\n", *e);

	return 0;
}
