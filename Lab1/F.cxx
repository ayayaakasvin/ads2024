#include <iostream>

bool isPrime (int numero) {
    if (numero <= 3)
    {
        return true;
    } else if (numero % 2 == 0 || numero % 3 == 0) {
        return false;
    } else {
        for (int i = 5; i * i <= numero; i += 6)
        {
            if (numero % i == 0 || numero % (i + 2) == 0)
            {
                return false;
            }
        }
    }
    
    return true;
}

int getPrime (int times) {
    int counter = 0;
    int result;

    for (int i = 2; counter < times; i++)
    {
        if (isPrime(i))
        {
            counter++;
            result = i;
        }
    }
    
    return result;
}

int main () {
    int times;
    std::cin >> times;

    std::cout << getPrime(times);

    return 0;
}