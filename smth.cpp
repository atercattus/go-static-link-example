#include <stdio.h>

extern "C" int testX() {
    printf("Hello world from C++\n");
    fflush (stdout);
    return 42;
}
