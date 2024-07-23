#include "stdio.h"
#include "string.h"

int main() {
    char str[] = "this is a random string";
    short length = strlen(str);
    char temp;
    printf("original string: %s\n", str);
    for (int i = 0; i < length / 2; i++) {
        temp = str[i];
        str[i] = str[length - i - 1];
        str[length - i - 1] = temp;
    }
    printf("result: %s\n", str);
}
