var minimumSumSubarray = function(nums, l, r) {

    let min = Infinity;

    let lrWindows = [];
    let windowSum = [];

    // Setup windows
    
    let runningSum = 0; 
    for (let i = 0; i < nums.length; i++){
        if (i >= r){
            // Out of range already.
            break;
        }

        runningSum += nums[i];

        if (i >= (l - 1) && i <= (r - 1)){
            lrWindows.push([0, i]);
            windowSum.push(runningSum);
            if (runningSum > 0 && runningSum < min){
                min = runningSum; 
            }
        }
    }

    let removeWindow = [];

    // Move windows by 1 and adjust sums.
    while (lrWindows.length !== 0){
        for (let i = 0; i < lrWindows.length; i++){
            if (lrWindows[i][1] === nums.length - 1){
                // This window is done. Mark it to be deleted later.
                removeWindow.push(i);
            }

            windowSum[i] -= nums[lrWindows[i][0]];
            lrWindows[i][0]++;
            lrWindows[i][1]++;
            windowSum[i] += nums[lrWindows[i][1]];

            if (windowSum[i] > 0 && windowSum[i] < min){
                min = windowSum[i];
            }
        }

        for (let i = 0; i < removeWindow.length; i++){
            lrWindows.splice(removeWindow[i], 1);
            removeWindow.splice(i, 1);
        }
    }

    return min === Infinity ? -1 : min;
};

console.log(minimumSumSubarray([3, -2, 1, 4], 2, 3));
console.log(minimumSumSubarray([-2, 2, -3, 1], 2, 3));
console.log(minimumSumSubarray([1, 2, 3, 4], 2, 4));
