// gcc -mavx -C -g -Wa,-a,-ad -S -o test.asm test.c
#include "test.h"

#include <stdlib.h>
#include <string.h>
#include <math.h>

int the_mask = 0; // global so the compiler can't be sure what its value is for opt.

static double frand()
{
    return (double)rand() / (double)RAND_MAX;
}

static void randmat(double *A)
{
    for (int i=0; i < 4; i++)
        for (int j=0; j < 4; j++)
            A[4*i+j] = frand();
}

static void randvec(double *x)
{
    for (int i=0; i < 4; i++)
        x[i] = frand();
}

void vecmatvec_avx(const double* restrict x, const double* restrict A,
                   const double* restrict y, double* restrict out)
{
    asm volatile ("# avx code begin");  // looking at assembly with gcc -S
    __m256d yrow = _mm256_loadu_pd(y);
    __m256d xrow = _mm256_loadu_pd(x);

    __m256d a = _mm256_mul_pd(_mm256_loadu_pd(A), yrow);
    __m256d b = _mm256_mul_pd(_mm256_loadu_pd(A+4), yrow);
    __m256d c = _mm256_mul_pd(_mm256_loadu_pd(A+8), yrow);
    __m256d d = _mm256_mul_pd(_mm256_loadu_pd(A+12), yrow);

    // our task now is to get {sum(a), sum(b), sum(c), sum(d)}
    // This is tricky because there is no hadd instruction for avx

    // {a[0]+a[1], b[0]+b[1], a[2]+a[3], b[2]+b[3]}
    __m256d sumab = _mm256_hadd_pd(a, b);

    // {c[0]+c[1], d[0]+d[1], c[2]+c[3], d[2]+d[3]}
    __m256d sumcd = _mm256_hadd_pd(c, d);

    // {a[0]+a[1], b[0]+b[1], c[2]+c[3], d[2]+d[3]}
    __m256d blend = _mm256_blend_pd(sumab, sumcd, 0b1100);

    // {a[2]+a[3], b[2]+b[3], c[0]+c[1], d[0]+d[1]}
    __m256d perm = _mm256_permute2f128_pd(sumab, sumcd, 0x21);

    // {sum(a), sum(b), sum(c), sum(d)}
    __m256d Ay = _mm256_add_pd(perm, blend);

    // Now we want to take the dotproduct: dot(x, Ay)

    __m256d e = _mm256_mul_pd(Ay, xrow);

    // horizontal sum of e: tricky again.
    // {e[0]+e[1], e[0]+e[1], e[2]+e[3], e[2]+e[3]}
    __m256d tmp = _mm256_hadd_pd(e, e);

    // {e[2]+e[3], e[2]+e[3]}
    __m128d e23 = _mm256_extractf128_pd(tmp, 1);
    __m128d result = _mm_add_pd(_mm256_castpd256_pd128(tmp), e23);

    _mm_storel_pd(out, result);
    asm volatile ("# avx code end");
}

static double run_avx(double *x, double *A, double *y, int count)
{
    double result = 0.0;
    for (int i=0; i < count; i++)
    {
        int j = i & the_mask;
        vecmatvec_avx(x+j, A+j, y+j, &result);
    }
    return result;
}

void print_buffer(unsigned char *buf, size_t size) {
    double Ap[16];
    double xp[4], yp[4];
    randmat(Ap);
    randvec(xp);
    randvec(yp);
    static const int muls_per_run = 4096;
    run_avx(xp, Ap, yp, muls_per_run);
    // i256 x;
    // set_zero(x);
    // x[0] = 2;

    // i256 y;
    // set_zero(y);
    // x[0] = 3;

    i256 r;
    set_zero(r);

    // mul256(&x, &y, &r);

    // u_int64_t hi;
    // u_int64_t lo;
    // mult64to128(2, 3, &hi, &lo);

    void* p = &r;

    for (uint i = 0; i < 31; i++) {
        // printf("%X.", r[i]);
        printf("%X.", *(unsigned int *)(p+i));
    }
    printf("\n");

    // for (uint i = 0; i < size; i++) {
    //     printf("%X", buf[i]);
    // }
    // printf("\n");

    // _mm256_zeroall();
    // /* Initialize the two argument vectors */
    // __m256 evens = _mm256_set_ps(2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0, 16.0);
    // __m256 odds = _mm256_set_ps(1.0, 3.0, 5.0, 7.0, 9.0, 11.0, 13.0, 15.0);

    // /* Compute the difference between the two vectors */
    // __m256 result = _mm256_sub_ps(evens, odds);

    // /* Display the elements of the result vector */
    // float* f = (float*)&result;
    // printf("%f %f %f %f %f %f %f %f\n",
    //     f[0], f[1], f[2], f[3], f[4], f[5], f[6], f[7]);
}
