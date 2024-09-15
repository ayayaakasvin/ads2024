#include <iostream>
#include <stack>
#include <string>

bool teamDominance(const std::string& input) {
    char firstOne = input[0];
    for (char c : input) {
        if (firstOne != c) {
            return false;
        }
    }
    return true;
}

std::string EstablishResult(std::string voteOrder) {
    std::stack<int> KStack;
    std::stack<int> SStack;

    while (!teamDominance(voteOrder)) {
        for (size_t i = 0; i < voteOrder.length(); ) {
            if (voteOrder[i] == 'K') {
                if (!SStack.empty()) {
                    voteOrder.erase(i, 1);
                    SStack.pop();
                } else {
                    KStack.push(1);
                    i++;
                }
            } else {
                if (!KStack.empty()) {
                    voteOrder.erase(i, 1);
                    KStack.pop();
                } else {
                    SStack.push(1);
                    i++;
                }
            }
        }
    }

    if (voteOrder[0] == 'K') {
        return "KATSURAGI";
    }
    return "SAKAYANAGI";
}

int main() {
    int size;
    std::cin >> size;
    std::string voteOrder;
    std::cin >> voteOrder;
    std::cout << EstablishResult(voteOrder) << std::endl;
    return 0;
}
