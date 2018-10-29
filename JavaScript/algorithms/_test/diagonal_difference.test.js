const diagonalDifference = require("../diagonal_difference.js").diagonalDifference

test("compareTheTriplets() can compare stuff", () => {
    const matrix = [
       [11, 2, 4],
       [4, 5, 6],
       [10, 8, -12]
    ];

    let result = diagonalDifference(matrix)

    let expected = 15

    expect(result).toEqual(expected)
})
