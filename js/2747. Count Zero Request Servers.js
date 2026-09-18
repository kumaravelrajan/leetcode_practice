import { Deque } from "@datastructures-js/deque";

/**
 * @param {number} n
 * @param {number[][]} logs
 * @param {number} x
 * @param {number[]} queries
 * @return {number[]}
 */
var countServers = function(n, logs, x, queries) {
    let originalQueryIndex = {};

    // queries can have duplicates. No guarantee that they are all unique values. Hence, key: queries[i], value : list of i - [i, j, k] etc. Then at the end when results are available and we want to insert in arr[i], we can simply do 
    // arr[originalQueryIndex[number][0]] = (no. of zero request servers) and then popfront from originalQueryIndex[number].
    for (let i = 0; i < queries.length; i++){
        if (originalQueryIndex[queries[i]]){
            originalQueryIndex[queries[i]].pushBack(i);
        } else {
            originalQueryIndex[queries[i]] = new Deque([i]);
        }
    }

    // Sort queries so that we can easily run window.
    queries.sort((a, b) => {
        return a - b; 
    });

    // Extremes between which the window runs.
    let leftExtreme = -1, rightExtreme = -1;

    leftExtreme = queries[0] - x;
    rightExtreme = queries[queries.length - 1];

    // Sort logs according to time (index 1). If times are equal sort according to server id. 
    logs.sort((a, b) => {
        if (a[1] !== b[1]){
            return a[1] - b[1];
        } else {
            return a[0] - b[0];
        }
    });

    return -1;
};

console.log(countServers(3, [[2,4],[2,1],[1,2],[3,1]], 2, [3, 4]));