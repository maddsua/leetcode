#include <iostream>
#include <vector>
#include <algorithm>

class Solution {
public:
    long long countOperationsToEmptyArray(std::vector<int> nums) {
        const int n = nums.size();
        std::vector<int> ind(n);
        for (int i = 0; i < n; ++i) {
            ind[i] = i;
        }

        std::sort(ind.begin(), ind.end(), [&](const int x, const int y) {
            return nums[x] < nums[y];
        });

		for(auto item : ind) {
			std::cout << item << "\n";
		}

        int m = n;
        long long r = 0;
        for (int i = 1; i < n; ++i) {
            if (ind[i] < ind[i - 1]) {
                r += m;
                m = n - i;
            }
        }
        r += m;
        return r;
        
    }
};

int main() {

	Solution sol;

	sol.countOperationsToEmptyArray({3, 4, -1});

	return 0;
}