// Complete the compareTriplets function below.
function compareTriplets(a, b) {

    let aScore = 0
    let bScore = 0
    let i = 0
    let value

    for (value of a) {
        if (a[i] > b[i]) {
            aScore++
            i++
        } else if (b[i] > a[i]) {

            bScore++
            i++
        } else {
            i++

        }

    }

    let result = [aScore, bScore]
    return result

}

module.exports.compareTriplets = compareTriplets