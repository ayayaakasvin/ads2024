#include <iostream>

bool isPrime (int numero) {
    if (numero <= 1)
    {
        return false;
    } else if (numero <= 3) {
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

int main () {
    int numero;
    std::cin >> numero;

    std::cout << (isPrime(numero) ? "YES" : "NO");

    return 0;
}