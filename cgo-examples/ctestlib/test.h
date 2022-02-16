#ifndef _TESTLIB_H_
#define _TESTLIB_H_

#include <stdio.h>
#include <stdlib.h>

int sum(int a, int b);
const char* get_string();
void print_string(char* a);
void print_buffer(unsigned char *buf, size_t size);
typedef struct point
{
    int x;
    int y;
} point;

int point_diff(point p);
void pass_void_pointer(void *ptr);
enum status {
    PENDING,
    DONE,
};
void generate_numbers(uint num, uint callback);
void user_action(uint callback);

#endif
