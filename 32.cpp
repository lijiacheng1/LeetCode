#include<iostream>
#include<string>
#include <limits>

using namespace std;

class Solution {
public:
    int longestValidParentheses(string s) {
        auto it = s.begin();
        vector<int> rem(s.size());
        stack<int> leftrem;
        int left = 0, result = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '(') {
                leftrem.push(i);
                rem[i] = 0;
            } else {
                if (leftrem.empty()) {
                    rem[i] = -1;
                } else {
                    rem[i] = rem[i - 1] + rem[leftrem.top()] + 2;
                    if (leftrem.top() - 1 >= 0 && rem[leftrem.top() - 1] != -1) {
                        rem[i] += rem[leftrem.top() - 1];
                    }
                    leftrem.pop();
                }
                result = result > rem[i] ? result : rem[i];
            }
        }
        return result;
    }
};