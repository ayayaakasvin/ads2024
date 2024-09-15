#include <iostream>
#include <stack>

bool isBalanced(std::string input) {
    std::stack<char> s;
    for (auto &&i : input)
    {
        if (!s.empty() && s.top() == i)
        {
            s.pop();
        } else {
            s.push(i);
        }
    }
    
    return s.empty();
}

int main() {
    std::string input;
    std::cin >> input;

    std::cout << (isBalanced(input) ? "YES" : "NO") << std::endl;
    return 0;
}
