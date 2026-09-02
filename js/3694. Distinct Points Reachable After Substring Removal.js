var distinctPoints = function(s, k) {
    let finalPoint = [0, 0]
    let reverse = {'U':'D', 'D':'U', 'L':'R', 'R':'L'};
    let uniquePoints = new Set();

    for (let i = 0; i < s.length; i++){
        movePoint(finalPoint, s[i]);
    }

    // finalPoint now contains the end point after all the computations in s. Now create a sliding window of length k and do the reverse of the operations on finalPoint as mentioned in the window. 

    let l = 0, r = 0; 
    let currWindow = [0, 0]

    for (; r < s.length; r++){
        movePoint(currWindow, s[r]);

        if (r - l + 1 >= k){
            // Window at limit. Record unique point and then shrink from left.
            let negCurrWindow = [-currWindow[0], -currWindow[1]];
            uniquePoints.add(JSON.stringify(addPoints(finalPoint, negCurrWindow)));
            movePoint(currWindow, reverse[s[l]]);
            l++;
        }
    }

    return uniquePoints.size;
};

var addPoints = function (point1, point2){
    return [point1[0] + point2[0], point1[1] + point2[1]];
}

var movePoint = function (point, direction){
    switch (direction){
        case 'U':
            point[1]++;
            break;
        case 'D':
            point[1]--;
            break;
        case 'L':
            point[0]--;
            break;
        case 'R':
            point[0]++;
            break;
    }
}

console.log(distinctPoints("LULDR", 2));