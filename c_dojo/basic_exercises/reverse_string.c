#include "stdio.h"
#include "string.h"

int main() {
    char str[] = "this";
    short length = strlen(str);
    char temp;
    printf("orignal string: %s\n", str);
    for (int i; i < length / 2; i++) {
        temp = str[i];
        str[i] = str[length - i - 1];
        str[length - i - 1] = temp;
    }
    printf("result: %s\n", str);
}
