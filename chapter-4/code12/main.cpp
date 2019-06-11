#include <cstdlib>
#include <iostream>
#include <ctime>
#include <chrono>
#include <thread>

void MultiplyMatrix(int **A, int **B, int **C, int n, int threads, int id) {
	for (int i = id; i < n; i+=threads) {
		for (int j = 0; j < n; j++) {
			for (int k = 0; k < n; k++) {
				C[i][j] = C[i][j] + A[i][k] * B[k][j];
			}
		}
	}
}

void InitMatrix(int **A, int n, bool random) {
	for (int i = 0; i < n; i++) {
        A[i] = (int*) malloc(sizeof(int) * n);
		for (int j = 0; j < n; j++) {
            if (random) {
                A[i][j] = rand() % 10000;
            } else {
                A[i][j] = 0;
            }
		}
	}
}

void PrintMatrix(int **A, int n) {
    for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
            std::cout << A[i][j] << ",";
        }
        std::cout << std::endl;
    }
}

int main(int argc, char *argv[]) {
    srand(time(NULL));
	if (argc < 2) {
		std::cout << "Usage: " << argv[0] << " <size of matrix> <number of threads>" << std::endl;
		exit(1);
	}
	int matrix_size = atoi(argv[1]);
	if (matrix_size == 0) {
		std::cout << "Cannot use 0 as matrix size" << std::endl;
		exit(0);		
	}
    int threads = atoi(argv[2]);
	if (threads < 1 || threads > 8) {
		std::cout << "Number of threads must be between 1 and 8" << std::endl;
		exit(0);		
	}
    int **A = (int**) malloc(sizeof(int*) * matrix_size);
    int **C = (int**) malloc(sizeof(int*) * matrix_size);
    InitMatrix(A, matrix_size, true);
    InitMatrix(C, matrix_size, false);

    auto start = std::chrono::high_resolution_clock::now();
    auto ref_threads = new std::thread[matrix_size];
    for (int i = 0; i < threads; i++) {
        ref_threads[i] = std::thread(MultiplyMatrix, A, A, C, matrix_size, threads, i);
    }
    for (int i = 0; i < threads; i++) {
        ref_threads[i].join();
    }
    auto end = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double, std::milli> dur = end - start;
    std::cout << "Elapsed in seconds:" << dur.count() / 1000 << std::endl;
}
