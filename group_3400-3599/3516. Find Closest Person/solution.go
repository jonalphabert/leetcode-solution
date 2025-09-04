func findClosest(x int, y int, z int) int {
    diffX := int(math.Abs(float64(z-x)))
    diffY := int(math.Abs(float64(z-y)))

    if diffX == diffY {
        return 0
    } else if diffX > diffY {
        return 2
    } else {
        return 1
    }
}