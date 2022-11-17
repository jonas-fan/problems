// #binary-search

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

    for left < right {
        mid := (left + right) >> 1

        switch guess(mid) {
        case 1:
            left = mid + 1
        case -1:
            right = mid
        default:
            return mid
        }
    }

    return left
}
