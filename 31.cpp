#include<iostream>
#include<string>
#include <limits>

using namespace std;

class Solution {
public:
    void nextPermutation(vector<int> &nums) {
        vector<int> vec_rem;
        int rem = 0;
        for (int i = nums.size()-1; i >= 0; i--) {
            if (nums[i] < rem) {
                sort(vec_rem.begin(), vec_rem.end());
                for (vector<int>::iterator it = vec_rem.begin(); it != vec_rem.end(); it++) {
                    if (*it > nums[i]) {
                        int temp = nums[i];
                        nums.pop_back();
                        nums.push_back(*it);
                        it = vec_rem.erase(it);
                        vec_rem.insert(it, temp);
                        nums.insert(nums.end(), vec_rem.begin(), vec_rem.end());
                        return;
                    }
                }
            } else {
                vec_rem.push_back(nums[i]);
                rem = nums[i];
                nums.pop_back();
            }
        }
        sort(vec_rem.begin(), vec_rem.end());
        nums.insert(nums.end(), vec_rem.begin(), vec_rem.end());
        return;
    }
};