var numOfSubarrays = function(arr, k, threshold) {
    let l = 0, r = 0; 
    let windowSum = 0; 
    let result = 0; 

    for (; r < arr.length; r++){
        windowSum += arr[r];

        if(r > k - 1){
            windowSum -= arr[l];
            l++;
        }

        if ((r - l + 1 === k) && (windowSum / k) >= threshold){
            result++;
        }
    }
    return result;
};

console.log(numOfSubarrays([2,2,2,2,5,5,5,8], 3, 4));