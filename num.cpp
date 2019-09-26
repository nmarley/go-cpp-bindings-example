// num.cpp
#include "nummer.hpp"
#include "num.h"
#include <stdio.h>

void DoSomething(void *p, int size) {
    // char* ptr = (char*)p;
    // printf("hello world : %d\n", *ptr);
    uint8_t* ptr = (uint8_t*)p;
    // printf("hello world : %d\n", *ptr);

    for (int i = 0 ; i < size; i++) {
        printf("hello world : %d\n", ptr[i]);
    }

// bls::PrivateKey sk = bls::PrivateKey::FromSeed(seed, sizeof(seed));
// bls::PublicKey pk = sk.GetPublicKey();

    printf("size : %d\n", size);
    return;
}

Num NumInit() {
    cxxNum * ret = new cxxNum(1);
    return (void*)ret;
}

void NumFree(Num n) {
    cxxNum * num = (cxxNum*)n;
    delete num;
}

void NumIncrement(Num n) {
    cxxNum * num = (cxxNum*)n;
    num->Increment();
}

int NumGetValue(Num n) {
    cxxNum * num = (cxxNum*)n;
    return num->GetValue();
}
