/**
 * @param {number} n
 * @return {string}
 */
var countAndSay = function(n) {
    let casRec = "1";
    n--;

    while (n != 0){
        casRec = findRLE(casRec);
        n--;
    }

    return casRec;
};

let findRLE = function(str){
    let curr = "-1", count = 0, result = ""

    for (let i = 0; i < str.length; i++){
        if (curr < 0) {
            curr = str[i]
            count++
        } else if (curr !== str[i]){
            // The previous sequence is over.
            result += (`${count}${curr}`);
            curr = str[i]
            count = 1
        } else if (curr == str[i]){
            count++
        }
    }

    result += (`${count}${curr}`);

    return result;
};

console.log(countAndSay(4));