// #array

type StockSpanner struct {
    prices []int
}

func Constructor() StockSpanner {
    return StockSpanner {
        prices: make([]int, 0),
    }
}

func (this *StockSpanner) Next(price int) int {
    out := 1

    for i := len(this.prices) - 1; i >= 0; i-- {
        if this.prices[i] > price {
            break
        }

        out++
    }
    
    this.prices = append(this.prices, price)

    return out
}
