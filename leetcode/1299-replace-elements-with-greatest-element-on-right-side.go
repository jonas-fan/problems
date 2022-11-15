// #array

func replaceElements(arr []int) []int {
    maximum := -1

    for i := len(arr) - 1; i >= 0; i-- {
        value := arr[i]

        arr[i] = maximum

        if value > maximum {
            maximum = value
        }
    }

    return arr
}
