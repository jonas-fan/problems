// #array

import (
    "sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(lhs int, rhs int) bool {
        return boxTypes[lhs][1] > boxTypes[rhs][1]
    })

    units := 0

    for truckSize > 0 && len(boxTypes) > 0 {
        num := boxTypes[0][0]

        if num > truckSize {
            num = truckSize
        }

        units += num * boxTypes[0][1]
        truckSize -= num
        boxTypes = boxTypes[1:]
    }

    return units
}
