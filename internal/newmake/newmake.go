package newmake

type User struct {
	Name   string
	Scores []int
}

// 01. NewInt создаёт указатель на int с заданным значением.
// TODO: возвращаемый указатель не должен быть nil.
func NewInt(value int) *int {
	var newValue *int
	newValue = &value
	return newValue
}

// 02. NewString создаёт указатель на string с заданным значением.
// TODO: возвращаемый указатель не должен быть nil.
func NewString(value string) *string {
	var newValue *string
	newValue = &value
	return newValue
}

// 03. NewZeroUser создаёт указатель на zero value User.
// TODO: указатель должен быть ненулевым, поля — иметь zero value.
func NewZeroUser() *User {
	var newUser *User
	newUser = &User{
		Name:   "",
		Scores: nil,
	}
	return newUser
}

// 04. NewNamedUser создаёт указатель на User с заполненным Name.
// TODO: Scores остаётся в zero value.
func NewNamedUser(name string) *User {
	var newUser *User
	newUser = &User{
		Name:   name,
		Scores: nil,
	}
	return newUser
}

// 05. MakeInts создаёт ненулевой слайс int заданной длины.
// TODO: элементы должны получить zero value.
func MakeInts(length int) []int {
	sl := make([]int, length)
	return sl
}

// 06. MakeIntsWithCapacity создаёт слайс с заданными len и cap.
// TODO: в тестах capacity всегда не меньше length.
func MakeIntsWithCapacity(length, capacity int) []int {
	sl := make([]int, length, capacity)
	return sl
}

// 07. MakeAndAppend создаёт слайс с начальной capacity и добавляет values.
// TODO: итоговый порядок значений должен совпадать с порядком аргументов.
func MakeAndAppend(capacity int, values ...int) []int {
	sl := make([]int, 0, capacity)
	sl = append(sl, values...)
	return sl
}

// 08. MakeEmptyStrings создаёт пустой, но не nil слайс строк.
// TODO: len должен быть 0.
func MakeEmptyStrings() []string {
	sl := make([]string, 0)
	return sl
}

// 09. SliceState возвращает len, cap и результат сравнения с nil.
// TODO: функция не должна изменять входной слайс.
func SliceState(items []int) (length, capacity int, isNil bool) {
	return len(items), cap(items), items == nil
}

// 10. EmptySliceState создаёт пустой слайс через make и возвращает его состояние.
// TODO: верните len, cap и признак nil.
func EmptySliceState() (length, capacity int, isNil bool) {
	sl := make([]int, 0)
	return len(sl), cap(sl), sl == nil
}

// 11. MakeUserScores создаёт указатель на User и ненулевой слайс Scores
// заданной длины.
// TODO: Name и длина Scores должны соответствовать аргументам.
func MakeUserScores(name string, count int) *User {
	var newUser *User
	newUser = &User{
		Name:   name,
		Scores: make([]int, count),
	}
	return newUser
}

// 12. NewSlicePointer возвращает результат new для типа []int.
// TODO: сам указатель не nil, но слайс по нему остаётся nil.
func NewSlicePointer() *[]int {
	return new([]int)
}

// 13. MakeSlicePointer создаёт пустой ненулевой слайс с заданной capacity
// и возвращает указатель на него.
// TODO: len итогового слайса равен 0.
func MakeSlicePointer(capacity int) *[]int {
	newSlice := make([]int, 0, capacity)
	var slicePointer *[]int
	slicePointer = &newSlice
	return slicePointer
}

// 14. MakeStringIntMap создаёт пустую рабочую map[string]int.
// TODO: map должна принимать новые значения без panic.
func MakeStringIntMap() map[string]int {
	newMap := make(map[string]int)
	return newMap
}

// 15. MakeBoolChannel создаёт chan bool с указанным размером буфера.
// TODO: buffer=0 означает небуферизованный канал.
func MakeBoolChannel(buffer int) chan bool {
	if buffer < 0 {
		buffer = 0
	}
	newChan := make(chan bool, buffer)
	return newChan
}
