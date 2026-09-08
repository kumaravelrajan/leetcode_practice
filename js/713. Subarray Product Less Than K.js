var numSubarrayProductLessThanK = function(nums, k) {
    let res = 0; 

    let l = 0, r = 0; 
    let currProd = 1;

    for (; r < nums.length; r++){
        currProd *= nums[r];

        if (currProd < k){
            // Valid product. This means we can add the new nums[r] and all ensuing subarrays to res. 
            res += (r - l + 1);
        } else {
            while (currProd >= k){
                // Shrink window from left
                currProd /= nums[l];
                l++;
            }

            res += (r - l + 1);
        }
    }

    return res;
};

console.log(numSubarrayProductLessThanK([2], 1));