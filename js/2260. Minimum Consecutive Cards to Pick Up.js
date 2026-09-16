/**
 * @param {number[]} cards
 * @return {number}
 */
var minimumCardPickup = function(cards) {
    // Monotonicity property holds. So dynamic sliding window can be used. 

    let l = 0, r = 0; 
    recordNum = {};
    result = Infinity;

    for (; r < cards.length; r++){
        let curr = cards[r];

        recordNum[curr] = (recordNum[curr] || 0) + 1;

        if (recordNum[curr] === 2){
            let temp = recordNum[curr];

            while(recordNum[curr] === 2){
                recordNum[cards[l]]--;
                l++;
            }

            l--;

            let cardDist = r - l + 1;

            if (cardDist % 2 !== 0) cardDist++;

            if (cardDist < result){
                result = cardDist;
            }

            l++;
        }
    }

    return result === Infinity ? -1 : result;
};

console.log(minimumCardPickup([3,4,1,2,3,5,6,5,99,98]));