var getSubarrayBeauty = function(nums, k, x) {
    let l = 0, r = 0;
    let result = [];
    let freqArr = Array(50).fill(0);

    for (; r < nums.length; r++){
        if (nums[r] < 0 && r <= k - 1){
            // This needs to be considered by us. 
            freqArr[nums[r] + 50]++;
        }

        if (r >= k ){
            // when r === k, this condition first triggers. This means l = 0, r = k. In total we have k + 1 values. But we do not add the nums[k] value to the freqArr immediately. We process the freqArr for the range 0...k-1 get the result and then process the k index value for freqArr.
            // In the next run, 1...k will be processed first, result pushed and then k+1 index value will be pushed. This means at the end we will still need to do one final push for the last range.

            let count = 0;
            for (let i = 0; i < freqArr.length; i++){
                count += freqArr[i];

                if (count >= x) {
                    result.push(i - 50);
                    break;
                }
            }

            if (count < x){
                result.push(0);
            }

            if (nums[l] < 0){
                freqArr[nums[l] + 50]--;
            }
            l++;

            if (nums[r] < 0){
                freqArr[nums[r] + 50]++;
            }
        }
    }

    // One final range to be processed n-k-1....n-1
    let count = 0; 
    for (let i = 0; i < freqArr.length; i++){
        count += freqArr[i];

        if (count >= x){
            result.push(i - 50);
            break;
        }
    }

    if (count < x){
        result.push(0);
    }

    return result; 

};

console.log(getSubarrayBeauty([-50,14], 2, 2));
console.log(getSubarrayBeauty([-47,44,46], 2, 2));