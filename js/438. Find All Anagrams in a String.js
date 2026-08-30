var findAnagrams = function(s, p) {

    if (p.length > s.length){
        return []; 
    }

    let targetChars = {};
    let missing = 0;
    let result = []; 
    let left = 0, right = 0; 

    // Setup the object for quick lookup
    for (let i = 0; i < p.length; i++){
        targetChars[p[i]] = (targetChars[p[i]] || 0 ) + 1;
        missing++;
    }

    // Initial setting up of window of length p.length
    for (; right < p.length; right++){

        if (targetChars[s[right]] !== undefined){
            // Char of s found in p.

            if (targetChars[s[right]] > 0){
                missing--;
                
            } else {
                missing++;
            }

            targetChars[s[right]]--;

            if (missing === 0){
                result.push(left);
            }

        }
    }

    // Now starting from index p.length in s go all the way to the end, all the while maintaining the window length to be p.length
    for (; right < s.length; right++){

        // Move left by 1.
        if (targetChars[s[left]] !== undefined){
            if (targetChars[s[left]] < 0){
                missing--;
            } else {
                missing++;
            }

            targetChars[s[left]]++;
        }

        

        left++;

        // Assess right

        if (targetChars[s[right]] !== undefined){
            // relevant character found

            if (targetChars[s[right]] > 0){
                missing--;
            } else {
                missing++;
            }

            targetChars[s[right]]--;

            if (missing === 0){
                // Anagram found
                result.push(left);
            }
        }
    }

    return result; 
};

console.log(findAnagrams("bpaa", "aa"));