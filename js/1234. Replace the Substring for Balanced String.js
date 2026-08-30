var balancedString = function(s) {
    let lookup = {"Q": 0, "W": 0, "E": 0, "R": 0};
    let limit = s.length / 4;

    for (let i = 0; i < s.length; i++) {
        lookup[s[i]]++;
    }

    // === DEVIATION 1 ===
    // Instead of splitting into excessChars/deficitChars objects and then
    // collapsing them into one combined "countToRemoveDeficit" threshold,
    // we track a per-character "need" — how many of THIS specific character
    // still have to appear in the window. Deficit chars just don't get an
    // entry here at all (we don't care how many of them are in the window).
    let need = {};
    for (const char of ["Q", "W", "E", "R"]) {
        if (lookup[char] > limit) {
            need[char] = lookup[char] - limit;
        }
    }

    // If nothing is over the limit, string is already balanced.
    if (Object.keys(need).length === 0) return 0;

    let l = 0, r = 0;
    let result = Infinity;

    // Helper: window is valid only when EVERY excess char's need is met,
    // not when some aggregate sum crosses a threshold.
    function windowIsValid() {
        return Object.values(need).every(v => v <= 0);
    }

    for (; r < s.length; r++) {
        // Only excess chars affect "need" — decrement if this char is one we're tracking.
        if (need[s[r]] !== undefined) {
            need[s[r]]--;
        }

        // === DEVIATION 2 ===
        // Your shrink loop was commented out, and even uncommented it was
        // checking the wrong condition (aggregate diff vs. threshold).
        // Here we shrink while the window is STILL valid (per-character),
        // which is how you actually minimize the window length —
        // shrinking a window that's no longer valid would break correctness.
        while (windowIsValid()) {
            result = Math.min(result, r - l + 1);

            // Removing s[l] from the window: if it was an excess char,
            // give back that unit of "need" since it's leaving the window.
            if (need[s[l]] !== undefined) {
                need[s[l]]++;
            }
            l++;
        }
    }

    return result;
};

console.log(balancedString("RQRQRWRWEEEE"));