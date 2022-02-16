#include <stdio.h>
#include <stdlib.h>
#include <immintrin.h>

// written for clarity, not conciseness
#define OSXSAVEFlag (1UL<<27)
#define AVXFlag     ((1UL<<28)|OSXSAVEFlag)
#define VAESFlag    ((1UL<<25)|AVXFlag|OSXSAVEFlag)
#define FMAFlag     ((1UL<<12)|AVXFlag|OSXSAVEFlag)
#define CLMULFlag   ((1UL<< 1)|AVXFlag|OSXSAVEFlag)

#define set_zero(var) for (uint i = 0; i < 31; i++) { var[i] = 0; }
typedef unsigned char byte;
typedef byte i256[32] __attribute__ ((aligned (512)));

// void mul256(i256 *x, i256 *y, i256 *r);

void print_buffer(unsigned char *buf, size_t size);
