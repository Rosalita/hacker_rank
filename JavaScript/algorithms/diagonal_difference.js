function diagonalDifference(arr) {

    const size = arr.length

    let primarySum = 0;
    let secondarySum = 0;
    for (let i = 0; i < size; i++) {
        primarySum += arr[i][i]
        secondarySum += arr[size - 1 - i][i]
    }

    let diff = Math.abs(primarySum - secondarySum)
    return diff
}

module.exports.diagonalDifference = diagonalDifference