#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <time.h>
#include <unistd.h>

#define MAX_THREADS 13
#define TZOLKIN 260
#define MAX_SEQUENCE_LENGTH 6

pthread_mutex_t lock;
int active = 0;

void *generateRandomSequence(void *arg) {
    char yinYang[2] = {'1', '0'};
    char sequence[MAX_SEQUENCE_LENGTH + 1];

    while (1) {
        pthread_mutex_lock(&lock);
        if (!active) {
            pthread_mutex_unlock(&lock);
            usleep(10000); // Sleep for 10ms when inactive
            continue;
        }
        pthread_mutex_unlock(&lock);

        int sequenceType = rand() % 3; // 0, 1, or 2
        int sequenceLength = (sequenceType == 2) ? 6 : sequenceType + 2;

        for (int i = 0; i < sequenceLength; i++) {
            int randIndex = rand() % 2;
            sequence[i] = yinYang[randIndex];
        }
        sequence[sequenceLength] = '\0';

        printf("%s ", sequence);

        int sleepTime = (rand() % TZOLKIN) + 1;
        usleep(sleepTime * 1000); // usleep takes microseconds
    }

    return NULL;
}

int main() {
    srand(time(NULL)); // Seed the random number generator
    pthread_mutex_init(&lock, NULL);
    pthread_t threads[MAX_THREADS];

    // Create a fixed number of threads
    for (int i = 0; i < MAX_THREADS; i++) {
        pthread_create(&threads[i], NULL, generateRandomSequence, NULL);
    }

    while (1) {
        pthread_mutex_lock(&lock);
        active = 1;
        pthread_mutex_unlock(&lock);

        int sleepTime = (rand() % TZOLKIN) + 1;
        sleep(sleepTime); // Sleep for a random duration

        pthread_mutex_lock(&lock);
        active = 0;
        pthread_mutex_unlock(&lock);

        // Let the threads rest for a bit before next activation
        sleep(1);
    }

    // Cleanup, although this part of the code is never reached
    for (int i = 0; i < MAX_THREADS; i++) {
        pthread_cancel(threads[i]);
        pthread_join(threads[i], NULL);
    }
    pthread_mutex_destroy(&lock);

    return 0;
}
