/**
 * @param {string} s
 * @return {string}
 * 给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 * 
 */
var longestPalindrome = function(s) {
    if(s.length<=1){
        return s
    }
    let sToArr = s.split('')
    let Max = ''
    sToArr.forEach((_, index) => {
        if(index === sToArr.length -1){
            return
        }
        let targetStr = _
        let evenTag = index
        Max = Max || _
        while(evenTag>0){
            evenTag--
            if(sToArr[evenTag] !== sToArr[index * 2 - evenTag]){
                break
            }
            targetStr = sToArr[evenTag] + targetStr + sToArr[index * 2 - evenTag]
            Max = targetStr.length >= Max.length ? targetStr : Max
        }
        let oddTag = index
        let oddStr = ''
        while(oddTag>=0){
            if(sToArr[oddTag] !== sToArr[index * 2 + 1 - oddTag]){
             //    Max = sToArr[oddTag].length > Max.length ? sToArr[oddTag] : Max
                break
            }
            oddStr = sToArr[oddTag] + oddStr +  sToArr[index * 2 + 1 - oddTag]
            Max = oddStr.length >= Max.length ? oddStr : Max
            oddTag--
        }
    })
    return Max
 };