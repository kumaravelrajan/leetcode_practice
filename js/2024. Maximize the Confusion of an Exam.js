var maxConsecutiveAnswers = function(answerKey, k) {
    let max = 0;
    let l = 0, r = 0;  
    let t = 0, f = 0;

    for (; r < answerKey.length; r++){
        let currChar = answerKey[r];
        if (currChar === 'T'){
            t++;
        } else {
            f++;
        }

        let minCharCount = Math.min(t, f);

        if (minCharCount > k){
            if (answerKey[l] === 'T'){
                t--;
            } else {
                f--;
            }
            l++;
        }

        max = (max > (r-l+1)) ? max : r - l + 1;
    }

    return max;
};

console.log(maxConsecutiveAnswers("TFFT", 1));