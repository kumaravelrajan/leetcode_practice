var numberOfSubstrings = function(s, k) {
    const n = s.length;
    const freq = new Array(26).fill(0);
    let left = 0;
    let result = 0;
    let good = 0; // counts how many characters appear at least k times

    for (let right = 0; right < n; right++) {
        const idx = s.charCodeAt(right) - 97;
        freq[idx]++;
        if (freq[idx] === k) {
            good++;
        }

        // Once we have a valid substring starting at `left`,
        // all longer substrings (with same left) are also valid.
        while (good > 0) {
            result += n - right;
            const leftIdx = s.charCodeAt(left) - 97;
            freq[leftIdx]--;
            if (freq[leftIdx] === k - 1) good--;
            left++;
        }
    }

    return result;
};

console.log(numberOfSubstrings("abacddb", 2));