/**
 * Count substrings of s that contain every vowel at least once
 * and have exactly k consonants.
 *
 * Trick: exactly(k) = atMost(k) - atMost(k-1)
 * atMost(k) is monotonic (consonant count only grows as window grows),
 * so it CAN be solved with a normal two-pointer sliding window.
 */
function countOfSubstrings(s, k) {
  const atMost = (maxConsonants) => {
    if (maxConsonants < 0) return 0; // can't have negative consonants

    const VOWELS = new Set(['a', 'e', 'i', 'o', 'u']);
    // lastSeen[v] = most recent index where vowel v has occurred so far
    const lastSeen = { a: -1, e: -1, i: -1, o: -1, u: -1 };

    let left = 0;        // smallest start that keeps consonants <= maxConsonants
    let consonants = 0;  // consonant count in current window [left, right]
    let total = 0;

    for (let right = 0; right < s.length; right++) {
      const ch =  s[right];
      if (VOWELS.has(ch)) {
        lastSeen[ch] = right;
      } else {
        consonants++;
      }

      // Standard "at most k" shrink: only ever moves left forward.
      while (consonants > maxConsonants) {
        if (!VOWELS.has(s[left])) consonants--;
        left++;
      }

      // Furthest-right start that still keeps ALL 5 vowels in the window.
      const minLastSeen = Math.min(
        lastSeen.a, lastSeen.e, lastSeen.i, lastSeen.o, lastSeen.u
      );
      if (minLastSeen === -1) continue; // some vowel hasn't appeared yet at all

      // Valid starts are those in [left, minLastSeen].
      const validStarts = minLastSeen - left + 1;
      if (validStarts > 0) total += validStarts;
    }

    return total;
  };

  return atMost(k) - atMost(k - 1);
}

console.log(countOfSubstrings("ieaouqqieaouqq", 1));