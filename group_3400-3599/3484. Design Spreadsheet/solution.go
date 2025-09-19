type Spreadsheet struct {
    row int
    sheet [26]map[int]int
}


func Constructor(rows int) Spreadsheet {
    var sheet [26]map[int]int
    for i := 0; i < 26; i++ {
        sheet[i] = make(map[int]int)
    }
    return Spreadsheet{
        row:   rows,
        sheet: sheet,
    }
}


func (this *Spreadsheet) SetCell(cell string, value int)  {
    column := int(cell[0] - 'A')
    r, _ := strconv.Atoi(cell[1:])
    rows := r - 1

    this.sheet[column][rows] = value 
}


func (this *Spreadsheet) ResetCell(cell string)  {
    column := int(cell[0] - 'A')
    r, _ := strconv.Atoi(cell[1:])
    rows := r - 1

    delete(this.sheet[column], rows)
}


func (this *Spreadsheet) GetValue(formula string) int {
    formula = formula[1:]
    parts := []string{}
    for i, c := range formula {
        if c == '+' {
            parts = append(parts, formula[:i], formula[i+1:])
            break
        }
    }

    return this.resolve(parts[0]) + this.resolve(parts[1])
}

func (this *Spreadsheet) resolve(token string) int {
    if token[0] >= 'A' && token[0] <= 'Z' {
        col := int(token[0] - 'A')
        r, _ := strconv.Atoi(token[1:])
        row := r - 1
        return this.sheet[col][row]
    }
    val, _ := strconv.Atoi(token)
    return val
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */