/**
 * @param {number[][]} tiles
 * @param {number} carpetLen
 * @return {number}
 */
var maximumWhiteTiles = function(tiles, carpetLen) {

    tiles.sort((a, b) =>{
        return a[0] - b[0];
    });

    let l = 0, r = 0; 
    let result = -Infinity;
    let tilesInWindow = 0;
    let blanksInWindow = 0;

    for (; r < tiles.length; r++){

        if (r > 0){
            blanksInWindow += tiles[r][0] - 1 - tiles[r-1][1];
        }

        let tilesInCurrentR = tiles[r][1] - tiles[r][0] + 1;

        if (tilesInCurrentR > carpetLen){
            return carpetLen;
        }

        if (tilesInWindow + blanksInWindow + (tilesInCurrentR) <= carpetLen){
            // Current interval is alright. It can be included in our calculations.
            tilesInWindow += tilesInCurrentR;

            if (tilesInWindow > result){
                result = tilesInWindow;
            }
        } else {
            // Including tilesInCurrentR would include tiles not coverable by carpet.
            
            let tilesCoverableInR = carpetLen - (tilesInWindow + blanksInWindow);

            if (tilesInWindow + tilesCoverableInR > result){
                result = tilesInWindow + tilesCoverableInR;
            }

            // Shrink window from left
            while(l < r && (tilesInWindow + blanksInWindow + (tilesInCurrentR)) > carpetLen){
                tilesInWindow -= tiles[l][1] - tiles[l][0] + 1;
                l++;
                blanksInWindow -= tiles[l][0] - 1 - tiles[l-1][1];

                tilesCoverableInR = carpetLen - (tilesInWindow + blanksInWindow);

                if (Math.min(tilesCoverableInR, tilesInCurrentR) + tilesInWindow > result){
                    result = Math.min(tilesCoverableInR, tilesInCurrentR) + tilesInWindow;
                }
            }

            tilesInWindow += tilesInCurrentR;

            if (tilesInWindow > result){
                result = tilesInWindow;
            }
        }
    }

    return result;

};

console.log(maximumWhiteTiles([[1, 100]], 10));
