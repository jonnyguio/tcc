#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

void *print_working(void *_) {
    int a[1000][1000];
    for (int i=0; i < 1000; i++) {
        for (int j = 0; j < 1000; j++) {
            a[i][j] = rand() % 100;
        }
    }
    while (1) {
        printf("working...\n");
        sleep(3);
    }
}

int main(int argc, char* argv[]) {
    srand((unsigned) time(NULL));
    pthread_t thread;    
    pthread_create(&thread, NULL, &print_working, NULL);
    sleep(10);
}
