// #search #binary-search

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left := n
    right := n

    if isBadVersion(n) {
        for isBadVersion(n) {
            right = n
            n = n >> 1
            left = n
        }
    } else {
        for !isBadVersion(n) {
            left = n
            n = n << 1
            right = n
        }
    }

    for left < right {
        mid := (left + right) >> 1

        if isBadVersion(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}
