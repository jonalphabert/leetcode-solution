func candy(ratings []int) int {
    numberChild := len(ratings)
    candyArr := make([]int, numberChild)
    numberOfCandy := 0

    for i := 0; i<numberChild; i++ {
        candyArr[i] = 1
    }

    // Tambah permen untuk anak lebih tinggi rating yang ada di sebelah kanan
    for i := 1; i<numberChild; i++ {
        if ratings[i] > ratings[i-1] {
            candyArr[i] = candyArr[i-1] + 1
        }
    }

    // Tambah permen untuk anak lebih tinggi di sebelah kiri, ini perlu dicari nilai max dari nilai sebelumnya
    for i := numberChild-1 ; i > 0 ; i-- {
        if ratings[i] < ratings[i-1] {
            candyArr[i-1] = max(candyArr[i-1], candyArr[i]+1)
        }
    }

    // Hitung semua permen
    for _, candy := range candyArr {
        numberOfCandy += candy
    }

    return numberOfCandy
}