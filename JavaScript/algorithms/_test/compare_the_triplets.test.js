const compareTriplets = require("../compare_the_triplets.js").compareTriplets

test("compareTheTriplets() can compare stuff", () => {
    const aPoints = [5, 6, 7]
    const bPoints = [3, 6, 10]

    let result = compareTriplets(aPoints, bPoints)

    let expected = [1,1]

    expect(result).toEqual(expected)
})
