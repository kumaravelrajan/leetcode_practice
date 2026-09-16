/**
 * @param {string} word
 * @param {number} k
 * @return {number}
 */
var countOfSubstrings = function(word, k) {
    function atMostK(maxConsonants){
        if (maxConsonants < 0) return 0;

        let l = 0, r = 0; 
        let result = 0; 
        let vowelLastSeenAt = {'a': -1, 'e':-1, 'i': -1, 'o': -1, 'u': -1};
        let consCount = 0;

        for (; r < word.length; r++){
            let ch = word[r];

            if (checkIfVowel(ch)){
                vowelLastSeenAt[ch] = r;
            } else {
                consCount++;
            }

            let minVowelIndex = Math.min(vowelLastSeenAt['a'], vowelLastSeenAt['e'], vowelLastSeenAt['i'], vowelLastSeenAt['o'], vowelLastSeenAt['u']);

            if (minVowelIndex === -1){
                // Some vowel has not yet appeared. 
                continue; 
            }

            while (consCount > maxConsonants){
                if (!checkIfVowel(word[l])){
                    consCount--;
                }
                l++;
            }

            let countRelevantSubarrays = minVowelIndex - l + 1; 

            if (countRelevantSubarrays > 0){
                result += countRelevantSubarrays;
            }
        }        

        return result; 
    }

    return atMostK(k) - atMostK(k-1);
}

function checkIfVowel(c){
    let vowelSet = new Set(['a','e','i','o','u',]);

    if (vowelSet.has(c)){
        return true;
    } else {
        return false;
    }
}

console.log(countOfSubstrings("ieaouqqieaouqq", 1));