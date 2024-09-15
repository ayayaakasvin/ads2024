#include <iostream>
#include <vector>

void printVector (std::vector<int> input) {
    for (auto &&i : input)
    {
        std::cout << i << ' ';
    }

    std::cout << std::endl;    
}

int main () {
    int N;
    std::cin >> N;

    std::vector<int> ages(N);
    std::vector<int> result(N, -1);

    for (int i = 0; i < N; i++)
    {
        std::cin >> ages[i];
    }
    
    for (int i = N - 1; i >= 0; i--)
    {
        for (int j = i - 1; j >= 0 ;j--)
        {
            if (ages[i] >= ages[j])
            {
                result[i] = ages[j];
                break;
            }
        }
    }

    printVector(result);
    return 0;
}