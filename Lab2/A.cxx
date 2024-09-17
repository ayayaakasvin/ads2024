#include <iostream>
#include <vector>

int NearestNumber (std::vector<int> numbers, int target)  {
    int difference, finalIndex = 0;
    difference = abs(target - numbers[0]);

    for (int i = 0; i < numbers.size(); i++)
    {
        int diff = abs(target - numbers[i]);

        if (diff == 0)
        {
            return i;
        }
        
        if (diff < difference)
        {
            difference = diff;
            finalIndex = i;
        }
    }
    
    return finalIndex;
}

int main () {
    int times;
    std::cin >> times;

    std::vector <int> numeros(times);
    for (int i = 0; i < times; i++)
    {
        std::cin >> numeros[i];
    }
    
    int target;
    std::cin >> target;

    std::cout << NearestNumber(numeros, target) << std::endl;
}