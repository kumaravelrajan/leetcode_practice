/**
 * @param {number[]} nums
 * @param {number} m
 * @param {number} k
 * @return {number}
 */
var maxSum = function(nums, m, k) {
    let result = -Infinity; 
    let windowSum = 0; 
    let l = 0, r = 0; 
    let uniqueWindowElems = {};
    let countUniqueWindowElems = 0;

    for (; r < nums.length; r++){
        let curr = nums[r];
        windowSum += curr;

        if (uniqueWindowElems[curr]){
            uniqueWindowElems [curr]++;
        } else {
            countUniqueWindowElems++;
            uniqueWindowElems[curr] = 1;
        }

        if (r > k - 1){
            uniqueWindowElems[nums[l]]--;
            windowSum -= nums[l];
            if (uniqueWindowElems[nums[l]] === 0){
                delete uniqueWindowElems[nums[l]];
                countUniqueWindowElems--;
            } 
            l++;
        }

        if (countUniqueWindowElems >= m && (r-l+1) === k && windowSum > result){
            result = windowSum; 
        }
    }

    return result === -Infinity? 0 : result; 
};