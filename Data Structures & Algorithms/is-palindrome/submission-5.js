class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        let newS = s.replaceAll('?','')
        newS = newS.replaceAll(',','')
        newS = newS.replaceAll(`'`,'')
        newS = newS.replaceAll(`.`,'')
        newS = newS.replaceAll(`!`,'')
        newS = newS.replaceAll(`:`,'')
        newS = newS.toLowerCase().replaceAll(' ','')
        console.log()
        const len = newS.length
        for (let i=0; i < len/2; i++){
            if (newS[i] !== newS[len-1-i]) {return false}
        }
        return true
    }
}
