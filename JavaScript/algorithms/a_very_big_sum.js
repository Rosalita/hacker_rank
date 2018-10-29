function aVeryBigSum(ar) {
    
    let sum = ar.reduce(add, 0); 
    // The reduce method reduces an array to a single value
    // array.reduce(function(total, currentValue, currentIndex, arr), initialValue)
    // reduce takes in a function and an initial value

    function add(a, b){
        return a + b
    }

    return sum

}

module.exports.aVeryBigSum = aVeryBigSum