func subtractProductAndSum(n int) int {
    sum := 0
    product := 1

    for n > 0 {
        digit := n % 10

        sum += digit
        product *= digit        
        n /= 10
    }

    return product - sum
}
