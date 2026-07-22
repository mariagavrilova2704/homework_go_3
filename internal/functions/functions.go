package functions

import (
	"errors"
	"fmt"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

// 01. SafeDivide выполняет целочисленное деление.
// TODO: при b == 0 верните 0 и ошибку "division by zero".
func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 02. FindUserByID ищет пользователя по ID.
// TODO: найденный пользователь возвращается с true; отсутствие — User{} и false.
func FindUserByID(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

// 03. FindActiveUser ищет пользователя с нужным ID только среди активных.
// TODO: неактивный пользователь считается ненайденным.
func FindActiveUser(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.Active == true && user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

// 04. SumAll складывает произвольное количество чисел.
// TODO: вызов без аргументов должен вернуть 0.
func SumAll(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

// 05. Apply вызывает переданную функцию op для a и b.
// TODO: верните результат op без изменения аргументов.
func Apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 06. ApplyIf применяет transform к value только при allowed=true.
// TODO: при false верните исходное value и не вызывайте transform.
func ApplyIf(value int, allowed bool, transform func(int) int) int {
	if allowed {
		return transform(value)
	}
	return value
}

// 07. NewCounter создаёт замыкание-счётчик.
// TODO: первый вызов возвращённой функции должен вернуть start+1,
// каждый следующий — ещё на единицу больше.
func NewCounter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// 08. NewAccumulator создаёт замыкание с накопленной суммой.
// TODO: каждый вызов добавляет аргумент к текущему состоянию и возвращает сумму.
func NewAccumulator(initial int) func(int) int {
	sum := initial
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

// 09. MakeMultiplier возвращает функцию умножения на factor.
// TODO: factor должен сохраняться внутри возвращённой функции.
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

// 10. DeferOrder возвращает строку, показывающую порядок defer.
// TODO: тело добавляет "body", затем зарегистрированы defer "first" и "second".
// Итоговая строка должна быть "body-second-first".
func DeferOrder() (result string) {
	result = "body"
	defer func() { result = result + "-first" }()
	defer func() { result = result + "-second" }()
	return result
}

// 11. CaptureDeferArgument демонстрирует вычисление аргумента defer сразу.
// TODO: сначала value="first", затем зарегистрируйте defer с value как параметром,
// после регистрации поменяйте value на "second". Верните захваченное значение.
func CaptureDeferArgument() (result string) {
	value := "first"
	defer func(v string) {
		result = v
	}(value)
	value = "second"
	return value
}

// 12. ReadDeferredVariable демонстрирует чтение переменной deferred-замыканием.
// TODO: deferred-функция без параметров должна прочитать value после того,
// как оно изменилось с "first" на "second".
func ReadDeferredVariable() (result string) {
	value := "first"
	defer func() {
		fmt.Println(value)
	}()
	value = "second"
	result = value
	return result
}

// 13. IncrementNamedResult возвращает value, увеличенный deferred-функцией на 1.
// TODO: используйте именованное возвращаемое значение.
func IncrementNamedResult(value int) (result int) {
	defer func() {
		result = value + 1
	}()
	result = value
	return result
}

// 14. RunWithCleanup выполняет action, а cleanup откладывает до выхода.
// TODO: верните две строки в порядке фактического выполнения: action, cleanup.
func RunWithCleanup(action, cleanup func() string) (events []string) {
	events = append(events, action())
	defer func() {
		events = append(events, cleanup())
	}()
	return events
}

// 15. ChooseOperation возвращает функцию для "add", "sub" или "mul".
// TODO: для неизвестного имени верните nil и false.
func ChooseOperation(name string) (func(int, int) int, bool) {
	if name == "add" {
		return func(i int, i2 int) int {
			return i + i2
		}, true
	}
	if name == "sub" {
		return func(i int, i2 int) int {
			return i - i2
		}, true
	}
	if name == "mul" {
		return func(i int, i2 int) int {
			return i * i2
		}, true
	}
	return nil, false
}
