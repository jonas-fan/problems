// #string

func allSame(nums []int, ignore int) bool {
    num := 0

    for i := 0; i < len(nums); i++ {
        if num = nums[i]; num != ignore {
            break
        }
    }

    for i := 0; i < len(nums); i++ {
        if nums[i] == ignore {
            continue
        } else if nums[i] != num {
            return false
        }
    }

    return true
}

func equalFrequency(word string) bool {
    freq := [26]int{}

    for _, letter := range word {
        freq[letter-'a']++
    }

    oldFreq := 0

    for i := 0; i < len(freq); i++ {
        if freq[i] == 0 {
            continue
        }

        oldFreq, freq[i] = freq[i], freq[i] - 1

        if allSame(freq[:], 0) {
            return true
        }

        freq[i] = oldFreq
    }

    return false
}
