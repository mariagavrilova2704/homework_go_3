package slices

import "slices"

// 01. First возвращает первый элемент и true.
// TODO: для nil и пустого слайса верните zero value и false.
func First(items []int) (int, bool) {
	if len(items) == 0 {
		return 0, false
	}
	return items[0], true
}

// 02. Last возвращает последнюю строку и true.
// TODO: для nil и пустого слайса верните пустую строку и false.
func Last(items []string) (string, bool) {
	if len(items) == 0 {
		return "", false
	}
	return items[len(items)-1], true
}

// 03. SliceInfo создаёт обычный срез items[low:high].
// TODO: верните сам срез, его len и cap. В тестах границы корректны.
func SliceInfo(items []int, low, high int) (part []int, length, capacity int) {
	part = items[low:high]
	length = high - low
	capacity = cap(items) - low
	return part, length, capacity
}

// 04. FullSliceInfo создаёт полный срез items[low:high:max].
// TODO: верните сам срез, его len и cap. В тестах границы корректны.
func FullSliceInfo(items []int, low, high, max int) (part []int, length, capacity int) {
	part = items[low:high:max]
	length = high - low
	capacity = max - low
	return part, length, capacity
}

// 05. ChangeFirst меняет первый элемент переданного слайса.
// TODO: пустой слайс оставьте без изменений.
func ChangeFirst(part []int, value int) {
	if len(part) != 0 {
		part[0] = value
	}
}

// 06. MutateWindowInFunction создаёт срез items[low:high], передаёт его
// в ChangeFirst и возвращает исходный items.
// TODO: изменение должно быть видно в исходном слайсе. Пустое окно не меняйте.
func MutateWindowInFunction(items []int, low, high, value int) []int {
	part := items[low:high]
	if len(part) != 0 {
		ChangeFirst(part, value)
	}
	return items
}

// 07. AppendWindowInFunction создаёт обычный срез, добавляет value
// и возвращает исходный слайс вместе с получившимся окном.
// TODO: поведение должно зависеть от доступной capacity: append может
// изменить исходный массив или перейти на новый.
func AppendWindowInFunction(items []int, low, high, value int) (source, part []int) {
	part = items[low:high]
	part = append(part, value)
	return items, part
}

// 08. AppendLimitedWindowInFunction создаёт окно с capacity, ограниченной high,
// затем добавляет value и возвращает исходный слайс и новое окно.
// TODO: добавление не должно перезаписывать элемент исходного слайса за high.
func AppendLimitedWindowInFunction(items []int, low, high, value int) (source, part []int) {
	part = items[low:high:high]
	part = append(part, value)
	return items, part
}

// 09. Clone возвращает независимую копию значений.
// TODO: изменение результата не должно менять items. Для nil верните nil.
func Clone(items []int) []int {
	cloneSlice := slices.Clone(items)
	return cloneSlice
}

// 10. ChangeClone создаёт независимую копию, меняет в ней элемент index
// и возвращает исходный слайс и копию.
// TODO: в тестах index корректен.
func ChangeClone(items []int, index, value int) (source, clone []int) {
	cloneSlice := slices.Clone(items)
	cloneSlice[index] = value
	return items, cloneSlice
}

// 11. AppendOne добавляет один элемент и возвращает результат append.
// TODO: сохраните исходный порядок элементов.
func AppendOne(items []int, value int) []int {
	newItems := append(items, value)
	return newItems
}

// 12. AppendMany добавляет все values в исходном порядке.
// TODO: корректно обработайте пустой список добавляемых значений.
func AppendMany(items []int, values ...int) []int {
	newItems := append(items, values...)
	return newItems
}

// 13. CanAppendWithoutGrow сообщает, хватает ли текущей capacity для extra
// новых элементов.
// TODO: отрицательное extra считается некорректным и даёт false.
func CanAppendWithoutGrow(items []int, extra int) bool {
	if extra < 0 {
		return false
	}
	if len(items)+extra <= cap(items) {
		return true
	}
	return false
}

// 14. SliceKind классифицирует слайс.
// TODO: nil -> "nil", ненулевой слайс с len=0 -> "empty",
// остальные значения -> "filled".
func SliceKind(items []int) string {
	if items == nil {
		return "nil"
	}
	if len(items) == 0 {
		return "empty"
	}
	return "filled"
}

// 15. AppendIndependent возвращает независимый результат items+values.
// TODO: исходный слайс и его базовый массив не должны измениться,
// даже если в исходной capacity есть свободное место.
func AppendIndependent(items []int, values ...int) []int {
	totalLen := len(items) + len(values)
	templateSlice := make([]int, totalLen)
	copy(templateSlice, items)
	copy(templateSlice[len(items):], values)

	return templateSlice
}
