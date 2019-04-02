#include <iostream>
#include <vector>
#include <set>

using namespace std;

class Solution {
public:
    void quickSort(vector<int> &nums,int lo,int hi){
        if(lo >= hi){
            return;
        }
        int temp = nums[lo];
        int i=lo,j=lo+1;
        while (j<=hi){
            if(nums[j] > temp){
                j++;
            } else{
                swap(nums[i+1],nums[j]);
                i++,j++;
            }
        }
        swap(nums[lo],nums[i]);
        quickSort(nums,i+1,hi);
        quickSort(nums,lo,i-1);
    }
};

int main() {
    vector<int> rem = {57,7,11,6,1,2,3,5};
    auto answer = new Solution;
    answer->quickSort(rem, 0,rem.size()-1);
    for (int i=0;i<rem.size();i++)
        cout<<rem.at(i)<<" ";
    return 0;
}