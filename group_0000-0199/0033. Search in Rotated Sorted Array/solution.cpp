class Solution {
public:
    int search(vector<int>& nums, int target) {
        int size = nums.size();
        int left = 0, right = size-1;
        while(left < right)
        {
            int mid = left + (right - left) / 2;

            // check if sorted, then left is the pivot Index
            if(nums[left] < nums[right]) break;

            if(nums[left] <= nums[mid]) left = mid+1;
            else right = mid;
        }
        
        // now left is pivotedIndex so we need to change left, right index for binser
        // If data not rotated, then search from 0 to last index   
        // if the data[0] less than target value, then we need to search from 0 to pivotedIndex-1
        // else then we need to search from pivotedIndex to last index
        if(left != 0){
            if(nums[0]<=target){
                right = left-1;
                left=0;
            } else {
                right = size-1;
            }
        }

        while(left <= right)
        {
            int mid = (left + right) / 2;

            // check if middle of data is the target
            if(nums[mid] == target) return mid;

            if(target > nums[mid]) left = mid+1;
            else right = mid-1;
        }
        
        return -1;
    }
};