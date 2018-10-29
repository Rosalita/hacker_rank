const aVeryBigSum = require("../a_very_big_sum.js").aVeryBigSum

test("aVeryBigSum() can add all numbers", () => {
    const nums = [1000000001, 1000000002, 1000000003, 1000000004, 1000000005]

    let result = aVeryBigSum(nums)

    let expected = 5000000015

    expect(result).toEqual(expected)
})
