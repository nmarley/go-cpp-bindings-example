#include <inttypes.h>
#include <stdlib.h>

// num.h
#ifdef __cplusplus
extern "C" {
#endif

typedef void* Num;
Num NumInit(void);
void NumFree(Num);
void NumIncrement(Num);
int NumGetValue(Num);
void DoSomething(void*, int);

#ifdef __cplusplus
}
#endif
