/**
 * @param {number} n
 * @param {number[][]} logs
 * @param {number} x
 * @param {number[]} queries
 * @return {number[]}
 */
/**
 * @param {number} n
 * @param {number[][]} logs
 * @param {number} x
 * @param {number[]} queries
 * @return {number[]}
 */
var countServers = function(n, logs, x, queries) {
    // Sort logs by time
    logs.sort((a, b) => a[1] - b[1]);

    // Pair each query with its original index so we can restore order later
    const qWithIdx = queries.map((q, i) => [q, i]);
    qWithIdx.sort((a, b) => a[0] - b[0]);

    // freq[serverId] = how many times serverId appears in the current window
    const freq = new Array(n + 1).fill(0);
    let distinctActive = 0; // number of servers with freq > 0

    let left = 0;
    let right = 0;
    const ans = new Array(queries.length);

    for (const [q, originalIdx] of qWithIdx) {
        // Add all logs with time <= q
        while (right < logs.length && logs[right][1] <= q) {
            const server = logs[right][0];
            if (freq[server] === 0) distinctActive++;
            freq[server]++;
            right++;
        }

        // Remove all logs with time < q - x
        while (left < right && logs[left][1] < q - x) {
            const server = logs[left][0];
            freq[server]--;
            if (freq[server] === 0) distinctActive--;
            left++;
        }

        // Servers with zero requests = total servers - active servers
        ans[originalIdx] = n - distinctActive;
    }

    return ans;
};

console.log(countServers(3, [[2,4],[2,1],[1,2],[3,1]], 2, [3, 4]));