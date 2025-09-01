/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left := 1
    right := n
    mid := left + (right - left) /2
    guessRes := guess(mid)

    for guessRes != 0 {
        if guessRes == -1 {
            right = mid
        } else if guessRes == 1 {
            left = mid + 1
        }
        mid = left + (right - left) /2
        guessRes = guess(mid)
    }

    return mid
}