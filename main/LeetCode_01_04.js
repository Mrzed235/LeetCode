/**
 * @param {string} s
 * @return {boolean}
 */
var canPermutePalindrome = function(s) {
    if(s.length === 1){
        return true
    }
    let sToArr = s.split('').sort()
    console.log(sToArr)
    if([...new Set(sToArr)].length === 1){
        return true
    }
    let arr = []
    sToArr.forEach((_, index) => {
        const sameNums = sToArr.lastIndexOf(_) - sToArr.indexOf(_)
        const nums = sameNums === 0 ? 1 : sameNums + 1
        if(sToArr.indexOf(_) === index){
            arr.push(nums)
        }
        // console.log(sToArr.lastIndexOf(_), sToArr.indexOf(_))
        // sToArr.splice(sToArr.lastIndexOf(_))
        // console.log(sToArr)
    })

    const evenNums = arr.filter(_ => _%2).length
    console.log(arr)
    console.log(evenNums)
    return evenNums <=1
};