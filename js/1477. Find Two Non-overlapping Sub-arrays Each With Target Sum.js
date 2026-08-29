/**
 * @param {number[]} arr
 * @param {number} target
 * @return {number}
 */
var minSumOfLengths = function(arr, target) {
    const n = arr.length;
    // dp[i] = length of the shortest valid subarray (sum === target)
    // that ends at or before index i. Infinity means none found yet.
    const dp = new Array(n).fill(Infinity);
    let left = 0, currSum = 0, ans = Infinity;

    for (let right = 0; right < n; right++) {
        currSum += arr[right];

        // Shrink window from the left while sum is too big
        while (currSum > target) {
            currSum -= arr[left];
            left++;
        }

        // Carry forward the best (shortest) match found so far,
        // up to (but not including) the current index
        dp[right] = right > 0 ? dp[right - 1] : Infinity;

        if (currSum === target) {
            const len = right - left + 1;

            // If there's a valid, non-overlapping match ending before
            // this window starts, combine it with the current window
            if (left > 0 && dp[left - 1] !== Infinity) {
                ans = Math.min(ans, len + dp[left - 1]);
            }

            // Update dp[right] in case this window is shorter than
            // whatever was carried forward
            dp[right] = Math.min(dp[right], len);
        }
    }

    return ans === Infinity ? -1 : ans;
};

console.log(minSumOfLengths([3,2,2,4,3], 3));
console.log(minSumOfLengths([7,3,4,7], 7));
console.log(minSumOfLengths([1,6,1], 7));
console.log(minSumOfLengths([4,3,2,6,2,3,4], 6));