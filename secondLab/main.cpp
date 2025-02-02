#include <iostream>
#include "stack.h"
#include "list.h"
using namespace std;


bool SolutionLabTwo(const string &s) {
    Stack* stack = stack_create();

    for (char c : s) {
        switch (c) {
            case '(': 
                stack_push(stack, ')'); 
                break;
            case '[': 
                stack_push(stack, ']'); 
                break;
            case '{': 
                stack_push(stack, '}'); 
                break;
            case '<': 
                stack_push(stack, '>'); 
                break;
            case ')':
            case ']':
            case '}':
            case '>':
                if (stack_empty(stack) || stack_get(stack) != c) {
                    stack_delete(stack);
                    return false;
                }
                stack_pop(stack);
                break;
            default:
                break;
            }
    }

    auto flag = stack_empty(stack);
    stack_delete(stack);
    return flag;
}

void WorkingQuestionMark(const string& s) {
    cout << '"' << s << "\" " << (SolutionLabTwo(s) ? "YES" : "NO") << '\n';
}

int main() {
    WorkingQuestionMark("[](");
    WorkingQuestionMark("([])");
    WorkingQuestionMark("([)");
}
