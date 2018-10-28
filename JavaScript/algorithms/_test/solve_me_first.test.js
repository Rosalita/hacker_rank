const solveMeFirst = require("../solve_me_first.js").solveMeFirst


test('solveMeFirst can add two positive numbers', () => {
  expect(solveMeFirst(5,1)).toEqual(6)
})

test('solveMeFirst can add two negative numbers', () => {
  expect(solveMeFirst(-5,-1)).toEqual(-6)
})

test('solveMeFirst can add a positive and a negative number', () => {
  expect(solveMeFirst(5,-1)).toEqual(4)
})