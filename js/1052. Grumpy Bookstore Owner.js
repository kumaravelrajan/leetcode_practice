/* Core idea: minutes needs to always be consecutive - so subarray. Already hinting at sliding window technique. 
    
We need to figure out using the consecutive minutes over which subarray makes the largest number of customers happy who would have been otherwise unhappy.

How do we do this?

Well, in a window of length minutes, we only add the customers who would normally be unsatisfied (corresponding to grumpy === i). This is the customer cost of grumpiness per window. The window which scores the highest in this metric is selected.

Apart from this all the usual customers who arrive at non grumpy times are accounted for.*/

var maxSatisfied = function(customers, grumpy, minutes) {

    let maxSum = -Infinity, maxWindowStart = null, maxWindowEnd = null, currSum = 0;

    for (let r = 0; r < customers.length; r++){

        if (grumpy[r] === 1){
            // Customer would go unsatisfied normally. With minutes applied they would go satisfied. Add these customers to the sum. 
            currSum += customers[r];
        }

        if (r > minutes - 1){
            // Window length exceeded.
            if (grumpy[r - minutes] === 1){
                currSum -= customers[r - minutes];
            }
        }

        if (currSum > maxSum){
            maxSum = currSum;
            maxWindowEnd = r;
        }
    }

    maxWindowStart = maxWindowEnd - (minutes - 1);
    let result = 0;
    for (let i = 0; i < customers.length; i++){
        if (grumpy[i] === 0 || (i >= maxWindowStart && i <= maxWindowEnd)){
            result += customers[i];
        }
    }

    return result;    
};

console.log(maxSatisfied([1,0,1,2,1,1,7,5], [0,1,0,1,0,1,0,1], 3));