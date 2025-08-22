package task

/*
LeetCode Version 1
Решение через map и уникальный ключ
Runtime 6ms
*/
func SingleNumber1(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	mapNums := make(map[int]int, len(nums))

	// O(n)
	for _, v := range nums {
		if _, ok := mapNums[v]; !ok {
			mapNums[v] = 1
		} else {
			mapNums[v] += 1
		}
	}

	var oneNums int

	// O(n)
	// TODO: Возможно ли сделать O(1)?
	for k, v := range mapNums {
		if v == 1 {
			oneNums = k
		}
	}

	return oneNums
}

/*
Version 2
Решение через map и уникальный ключ
Runtime 6ms
*/
func SingleNumber2(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	mapUniqueNums := make(map[int]int)
	var uniqueNums int //

	// 4, 2, 1, 1, 2, 3, 3
	for _, n := range nums {
		if _, ok := mapUniqueNums[n]; !ok {
			mapUniqueNums[n] = 0
		} else {
			delete(mapUniqueNums, n)
		}
	}

	for k := range mapUniqueNums {
		uniqueNums = k
	}

	return uniqueNums
}

/*
Version 3
Решение через XOR подсказали
Runtime приблизительно 0.5ms

func init() {
    debug.SetMemoryLimit(2)
}

Она нужна для того, чтобы ограничить максимальное количество памяти,
которое может использовать программа (точнее — сборщик мусора, GC).
Пользоваться нужно аккуратно и только в том случае, если уверен что твоя
программа будет тратить мало памяти. В основном успользуется только в Docker
*/

func SingleNumber(nums []int) int {

	var uniqueNums int //

	for _, n := range nums {
		uniqueNums = uniqueNums ^ n
	}

	return uniqueNums
}
