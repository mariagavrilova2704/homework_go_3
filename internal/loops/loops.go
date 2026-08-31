package loops

type User struct {
	Name   string
	Active bool
}

// 01. SumTo считает сумму целых чисел от 1 до n включительно.
// TODO: для n <= 0 верните 0.
func SumTo(n int) int {
	sum := 0
	if n > 0 {
		for i := 0; i <= n; i++ {
			sum += i
		}
		return sum
	}
	return 0
}

// 02. SumBetween считает сумму от start до end включительно.
// TODO: если start > end, верните 0.
func SumBetween(start, end int) int {
	sum := 0
	if start <= end {
		for i := start; i <= end; i++ {
			sum += i
		}
		return sum
	}
	return 0
}

// 03. CountDown возвращает числа от n до 1.
// TODO: для n <= 0 верните пустой ненулевой слайс.
func CountDown(n int) []int {
	if n <= 0 {
		return []int{}
	}
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = n - i
	}
	return count
}

// 04. Factorial вычисляет n! циклом.
// TODO: 0! и 1! равны 1; для отрицательного n верните 0.
func Factorial(n int) int {
	f := 1
	if n == 0 || n == 1 {
		return 1
	}
	if n < 0 {
		return 0
	}

	for i := 2; i <= n; i++ {
		f = f * i
	}
	return f
}

// 05. CountEven считает чётные числа в слайсе.
// TODO: ноль считается чётным.
func CountEven(items []int) int {
	count := 0
	for _, countNumber := range items {
		if countNumber%2 == 0 {
			count++
		}
	}
	return count
}

// 06. FirstNegative возвращает первое отрицательное значение и true.
// TODO: если отрицательных чисел нет, верните zero value и false.
func FirstNegative(items []int) (int, bool) {
	for _, value := range items {
		if value < 0 {
			return value, true
		}
	}
	return 0, false
}

// 07. SumWithoutZeros суммирует элементы, пропуская нули.
// TODO: используйте управляющую конструкцию, которая переходит
// к следующей итерации.
func SumWithoutZeros(items []int) int {
	sum := 0
	for _, item := range items {
		if item == 0 {
			continue
		}
		sum = sum + item
	}
	return sum
}

// 08. SumUntilLimit добавляет элементы по порядку, пока следующий элемент
// не сделал бы сумму больше limit.
// TODO: в этот момент остановите цикл. Для limit < 0 верните 0.
func SumUntilLimit(items []int, limit int) int {
	sum := 0
	if limit < 0 {
		return 0
	}
	for _, item := range items {
		if sum+item > limit {
			break
		}
		sum = sum + item
	}
	return sum
}

// 09. DoubleInPlace умножает каждый элемент слайса на 2.
// TODO: измените исходный слайс, а не копию value из range.
func DoubleInPlace(items []int) {
	for i, _ := range items {
		items[i] *= 2
	}
}

// 10. ReplaceNegativeInPlace заменяет отрицательные элементы на replacement.
// TODO: ноль и положительные значения должны сохраниться.
func ReplaceNegativeInPlace(items []int, replacement int) {
	for i, _ := range items {
		if items[i] < 0 {
			items[i] = replacement
		}
	}
}

// 11. CountActive считает активных пользователей.
// TODO: порядок пользователей не влияет на количество.
func CountActive(users []User) int {
	count := 0
	for _, user := range users {
		if user.Active {
			count++
		}
	}
	return count
}

// 12. ActiveNames возвращает имена только активных пользователей.
// TODO: сохраните исходный порядок.
func ActiveNames(users []User) []string {
	names := make([]string, 0, len(users))
	for _, user := range users {
		if user.Active {
			names = append(names, user.Name)
		}
	}
	return names
}

// 13. RuneCount считает Unicode-символы через range по строке.
// TODO: результат для кириллицы и emoji не должен равняться числу байт.
func RuneCount(text string) int {
	//runeArray := []rune(text)
	count := 0
	for range text {
		count++
	}
	return count
}

// 14. RuneByteIndexes возвращает байтовые индексы всех rune в строке.
// TODO: используйте индекс, который отдаёт range по string.
func RuneByteIndexes(text string) []int {
	byteArray := make([]int, 0, len(text))
	for i, _ := range text {
		byteArray = append(byteArray, i)
	}
	return byteArray
}

// 15. RepeatEach повторяет каждый элемент times раз подряд.
// TODO: для times <= 0 верните пустой ненулевой слайс.
func RepeatEach(items []int, times int) []int {
	if times <= 0 {
		return []int{}
	}
	totalSize := len(items) * times
	repeatArray := make([]int, 0, totalSize)

	for _, item := range items {
		for j := 1; j <= times; j++ {
			repeatArray = append(repeatArray, item)
		}
	}
	return repeatArray
}
