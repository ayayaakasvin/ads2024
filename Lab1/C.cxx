#include <iostream>
#include <string>

void cleanOperation(std::string& input) {
    int writePos = input.size() - 1;
    int hashtagCount = 0;

    for (int i = input.size() - 1; i >= 0; i--) {
        if (input[i] == '#') {
            hashtagCount++;
        } else {
            if (hashtagCount > 0) {
                hashtagCount--;
            } else {
                input[writePos--] = input[i];
            }
        }
    }

    input.erase(0, writePos + 1);
}

int main() {
    std::string input1, input2;
    std::cin >> input1 >> input2;
    cleanOperation(input1);
    cleanOperation(input2);
    std::cout << (input1 == input2 ? "Yes" : "No") << std::endl;
    return 0;
}