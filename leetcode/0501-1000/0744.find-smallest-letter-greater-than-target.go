// #binary-search

func nextGreatestLetter(letters []byte, target byte) byte {
    left := 0
    right := len(letters)

    for left < right {
        mid := left + (right - left) >> 1

        if letters[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    if left == len(letters) {
        return letters[0]
    }

    return letters[left]
}
