package task

// O(n2) Очень большое кол-во проходов по срезу
func PartsSums1(ls []uint64) []uint64 {
	sumList := make([]uint64, 0)

	if len(ls) == 0 {
		return []uint64{0}
	}

	for i := 0; i <= len(ls); i++ {
		var sumNum uint64 = 0
		sl := ls[i:]

		for _, i := range sl {
			sumNum += i
		}

		sumList = append(sumList, sumNum)
	}

	return sumList
}

// O(n) Два прохода по срезу
func PartsSums2(ls []uint64) []uint64 {
	sumL := make([]uint64, 0)
	sumN := uint64(0)

	// Если срез пустой, возвращаем срез с 0
	if len(ls) == 0 {
		return []uint64{0}
	}

	// Вычисляем общею сумму и добавляем ее в sumL
	for _, i := range ls {
		sumN += i
	}
	sumL = append(sumL, sumN)

	// Вычитаем из общей суммы и добавляем в срез
	for _, i := range ls {
		sumN -= i
		sumL = append(sumL, sumN)
	}

	return sumL
}

// Вариан ChatGpt с одним проходом
func PartsSums(ls []uint64) []uint64 {
	// Срез для результатов, в том числе для последнего элемента (0)
	sumL := make([]uint64, len(ls)+1)

	// Накопленная сумма (сначала 0, так как после последнего элемента должна быть 0)
	accumulatedSum := uint64(0)

	// Вычисляем общую сумму с конца
	for i := len(ls) - 1; i >= 0; i-- {
		accumulatedSum += ls[i]
		sumL[i] = accumulatedSum
	}

	return sumL
}
