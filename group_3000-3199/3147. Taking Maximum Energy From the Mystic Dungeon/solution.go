func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    maxEnergy := -1 * 1 << 32
    
    for i := n-k-1; i>=0; i-- {
        energy[i] = energy[i] + energy[i+k]
    }


    for i := 0; i<n; i++ {
        maxEnergy = max(maxEnergy, energy[i])
    }

    return maxEnergy
}