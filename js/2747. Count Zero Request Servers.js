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
    let serverIdFreqInWindow = {};
    let arr = new Array(queries.length).fill(-1);

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

    // Cleanup logs such that no record exists where log time < queries[0] - x (lowest possible value for inclusion.)

    let cleanupI = 0; 

    while(logs[cleanupI][1] < queries[0] - x){
        cleanupI++;
    }

    // Start the processing
    let l = 0, r = 0, queryI = 0, logsL = cleanupI, logsR = cleanupI; 
    for (r = leftExtreme; r <= rightExtreme; r++){
        while (logsR < logs.length && logs[logsR][1] === r){
            // The time of the current logs[logsR] matches r along the imaginary line. This means, the current log is valid for the current window. Record it in the hashmap.

            serverIdFreqInWindow[logs[logsR][0]] = (serverIdFreqInWindow[logs[logsR][0]] || 0) + 1;
            logsR++;
        }

        if (r === queries[queryI]){
            let temp = queries[queryI];
            
            while(queries[queryI] === temp){
                // Write answer to arr.

                arr[originalQueryIndex[queries[queryI]].front()] = n - Object.keys(serverIdFreqInWindow).length;
                originalQueryIndex[queries[queryI]].popFront();
                queryI++;
            }

            while (l < queries[queryI] - x){
                while (logs[logsL][1] < queries[queryI] - x){
                    serverIdFreqInWindow[logs[logsL][0]]--;

                    if (serverIdFreqInWindow[logs[logsL][0]] === 0){
                        delete serverIdFreqInWindow[logs[logsL][0]];
                    }

                    logsL++;
                }

                l++;
            }
        }
    }

    return arr;
};

console.log(countServers(3, [[1,3],[2,6],[1,5]], 5, [10, 11]));