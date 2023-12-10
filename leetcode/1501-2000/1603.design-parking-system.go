type ParkingSystem struct {
    count []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        count: []int{0, big, medium, small},
    }
}

func (s *ParkingSystem) AddCar(carType int) bool {
    if s.count[carType] <= 0 {
        return false
    }

    s.count[carType]--

    return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
