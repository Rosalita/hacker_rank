const simpleArraySum = require("../simple_array_sum.js").simpleArraySum

test('simpleArraySum can add an array of two positive numbers', () => {
    const nums = [1, 2]
    expect(simpleArraySum(nums)).toEqual(3)
  })

  test('simpleArraySum if given strings will add them to 0', () => {
    const nums = ["fish", "chips"]
    expect(simpleArraySum(nums)).toEqual("0fishchips")
  })

  test('simpleArraySum can add lots of numbers', () => {
    const nums = [1, 2, 5, -4, 0, 1]
    expect(simpleArraySum(nums)).toEqual(5)
  })
