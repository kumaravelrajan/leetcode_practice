var numberOfSubstrings = function(s) {

    let left = 0, right = 0; 
    let count = {'a': 0, 'b': 0, 'c': 0};
    let sum = 0; 

    for (; right < s.length; right++){

        count[s[right]]++

        while (count["a"] > 0 && count["b"] > 0 && count["c"] > 0) {
            // We have a complete set. Shrink left till we have an incomplete set.
            // Then, 0...left-1 is a valid count.

            count[s[left]]--
            left++
        }

        sum += left
    }

    return sum;
    
};

console.log(numberOfSubstrings("abcabc"));
console.log(numberOfSubstrings("aaacb"));