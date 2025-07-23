#include <stdio.h>

int add_int(int a, int b) { return a + b; }

float add_float(float a, float b) { return a + b; }

double add_double(double a, double b) { return a + b; }

#define add(a, b)           \
    _Generic((a),           \
        int: add_int,       \
        float: add_float,   \
        double: add_double, \
        default: printf("Invalid type \n"))

int main(int argc, char const *argv[]) {
    int ai = 10, bi = 20;
    float af = 10.5, bf = 20.5;
    double ad = 10.5, bd = 20.5;

    printf("Sum of int: %d\n", add(ai, bi));
    printf("Sum of float: %f\n", add(af, bf));
    printf("Sum of double: %lf\n", add(ad, bd));
    return 0;
}
