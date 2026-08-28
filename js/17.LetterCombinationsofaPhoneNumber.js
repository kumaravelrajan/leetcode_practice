var letterCombinations = function(digits) {
    let combos = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5":"jkl",
        "6":"mno",
        "7":"pqrs",
        "8":"tuv",
        "9":"wxyz"
    }

    function recurse (curDigits) {
        if (curDigits === ""){
            return []
        }

        let curr = curDigits[0];
        let result = recurse(curDigits.slice(1));
        let newResult = []

        if (result.length == 0){
            for (let i = 0; i < combos[curr].length; i++){
                newResult.push(combos[curr][i]);           
            }       
        } else {
            for (let i = 0; i < combos[curr].length; i++){
                for (let j = 0; j < result.length; j++){
                    newResult.push(combos[curr][i] + result[j])
                }
            }
        }

        return newResult
    }

    return recurse(digits);

    
};

console.log(letterCombinations("23"));