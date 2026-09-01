var maxFrequency = function(nums, k) {
    let l = 0, r = 1; 
    let maxFreq = -Infinity;
    let kBeforeWindow = k; 

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
            // Current window too big of a task for available k. So move l forward, reclaim invested k and proceed from there. 
 
            while (reqK > k){
                if (l < r - 1){
                    k += nums[r - 1] - nums[l];
                    l++;
                } else {
                    l++;
                    break;
                }
                
            }
        }
    }

    return (maxFreq === -Infinity) ? 1 : maxFreq;
};

// console.log(maxFrequency([1,2,4], 5)); //3
// console.log(maxFrequency([1,4,8,13], 5)); //2
// console.log(maxFrequency([3,9,6], 2)); //1
// console.log(maxFrequency([1, 100, 101], 10)); //2
console.log(maxFrequency([9930,9923,9983,9997,9934,9952,9945,9914,9985,9982,9970,9932,9985,9902,9975,9990,9922,9990,9994,9937,9996,9964,9943,9963,9911,9925,9935,9945,9933,9916,9930,9938,10000,9916,9911,9959,9957,9907,9913,9916,9993,9930,9975,9924,9988,9923,9910,9925,9977,9981,9927,9930,9927,9925,9923,9904,9928,9928,9986,9903,9985,9954,9938,9911,9952,9974,9926,9920,9972,9983,9973,9917,9995,9973,9977,9947,9936,9975,9954,9932,9964,9972,9935,9946,9966], 3056)); //73