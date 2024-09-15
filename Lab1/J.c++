#include <deque>
#include <vector>
#include <iostream>
#include <string>

// for checking if there are any errors
class output {
private:
    int data;
    bool error;

public:
    output(int inputData, bool isError) : data(inputData), error(isError) {}

    bool isError() const {
        return error;
    }

    int getData() const {
        return data;
    }
};

std::vector<output> JSolution(std::deque<int>& q) {
    std::vector<output> result;

    while (true) {
        std::string operation;
        std::cin >> operation;

        if (operation == "!") {
            break;
        }

        if (operation == "*") {
            if (q.empty()) {
                output goingToPush(0, true);
                result.push_back(goingToPush);
            } else if (q.size() == 1) {
                int front = q.front();
                result.push_back(output(front * 2, false));
                q.pop_front();
            } else {
                int front = q.front();
                int back = q.back();
                q.pop_front();
                q.pop_back();
                result.push_back(output(front + back, false));
            }
        } else if (operation == "+") {
            int numero;
            std::cin >> numero;
            q.push_front(numero);
        } else if (operation == "-") {
            int numero;
            std::cin >> numero;
            q.push_back(numero);
        } else {
            std::cout << "Unexpected operation: " << operation << std::endl;
        }
    }

    return result;
}

int main() {
    std::deque<int> a;
    
    std::vector<output> result = JSolution(a);
    for (const auto& i : result) {
        if (i.isError()) {
            std::cout << "error" << std::endl;
        } else {
            std::cout << i.getData() << std::endl;
        }
    }
    
    return 0;
}
