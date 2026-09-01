var maxFrequency = function(nums, k) {
    let l = 0, r = 1; 
    let maxFreq = -Infinity;

    // Sort the array
    nums.sort((a, b) =>{
        return a - b;
    });

    for (; r < nums.length; r++){
        let reqK = (nums[r] - nums[r - 1]) * (r - l);
        if(reqK <= k){
            // The current window l...r can be completely transformed such that l...r-1 can be incremented to r element.

            k -= reqK;

            if ((r - l + 1) > maxFreq){
                maxFreq = r - l + 1;
            }
        } else {
            // Current window too big of a task for available k. So move left forward and proceed from there. 

            while (reqK > k){
                k += nums[r] - nums[l];

                l++;
            }
        }
    }

    return (maxFreq === -Infinity) ? 1 : maxFreq;
};

console.log(maxFrequency([3,9,6], 2));