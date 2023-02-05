/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
    if(s.length<=1){
        return s.length
    }
    let pArr = s.split('')
    let Max = 0
    let str = ''
    console.log(pArr)
    pArr.forEach((_, index) => {
        if (!str.includes(_)) {
            str += _
            if(index + 1 === pArr.length){
                Max = str.length > Max ? str.length : Max
            }
            return
        }
         Max = str.length > Max ? str.length : Max
         str = str.slice(str.indexOf(_) + 1) + _
    })
    return Max
};
const str1 = "abcabcbb"
lengthOfLongestSubstring(str1)



