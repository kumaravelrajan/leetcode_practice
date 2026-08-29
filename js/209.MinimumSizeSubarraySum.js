var minSubArrayLen = function(target, nums) {

    let left = 0, right = 0; 

    let curLen = 0, minLen = Infinity;

    let runningSum = nums[0];

    while (left <= right && right < nums.length){
        curLen = right - left + 1

        if (runningSum >= target){
            if (curLen < minLen ){
                minLen = right - left + 1
            }
            runningSum -= nums[left]
            left++
        } else {
            right++
            runningSum += nums[right]
        }
    }

    return minLen;
};

console.log(minSubArrayLen(7, [2,3,1,2,4,3]))