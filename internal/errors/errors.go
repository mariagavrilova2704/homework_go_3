package errorshomework

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNotFound = errors.New("not found")

type FieldError struct {
	Field  string
	Reason string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

// 01. IsNil проверяет, отсутствует ли ошибка.
//
// TODO: верните true только для nil-ошибки.
func IsNil(err error) bool {
	if err == nil {
		return true
	}
	return false
}

// 02. RequireText проверяет обязательное текстовое поле.
//
// TODO: если value пустая строка, верните *FieldError с указанным field
// и причиной "is required". Для непустого значения верните nil.
func RequireText(field, value string) error {
	if value == "" {
		return &FieldError{
			Field:  field,
			Reason: "is required",
		}
	}
	return nil
}

// 03. ValidatePositive проверяет числовое поле.
//
// TODO: положительное число считается корректным. Для нуля и отрицательных
// значений верните *FieldError с причиной "must be positive".
func ValidatePositive(field string, value int) error {
	if value <= 0 {
		return &FieldError{
			Field:  field,
			Reason: "must be positive",
		}
	}
	return nil
}

// 04. ParseAge преобразует текст в целое число.
//
// TODO: верните полученное число. Если строка не является целым числом,
// верните ошибку с контекстом "parse age", не потеряв исходную причину.
func ParseAge(raw string) (int, error) {
	num, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parse age: %w", err)
	}
	return num, nil
}

// 05. WrapOperation добавляет к ошибке название выполняемой операции.
//
// TODO: nil должен остаться nil. Для ненулевой ошибки добавьте operation
// к тексту так, чтобы исходную ошибку можно было найти в цепочке.
func WrapOperation(operation string, err error) error {
	if err != nil {
		return fmt.Errorf("%s:%w", operation, err)
	}
	return nil
}

// 06. NotFound создаёт ошибку отсутствующего ресурса.
//
// TODO: текст должен содержать название resource, а ErrNotFound должна
// оставаться доступной как причина ошибки.
func NotFound(resource string) error {
	return fmt.Errorf("%s not found: %w", resource, ErrNotFound)
}

// 07. IsNotFound определяет, есть ли ErrNotFound в цепочке ошибок.
//
// TODO: корректно обработайте nil, прямую и многократно обёрнутую ошибку.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrNotFound)
}

// 08. SameCause проверяет, относится ли err к указанной причине target.
//
// TODO: учитывайте цепочку обёрнутых ошибок, а не только текст.
func SameCause(err, target error) bool {
	return errors.Is(err, target)
}

// 09. FieldFrom пытается извлечь *FieldError из цепочки ошибок.
//
// TODO: верните найденную ошибку и true. Если подходящего типа нет,
// верните nil и false.
func FieldFrom(err error) (*FieldError, bool) {
	var myErr *FieldError
	if errors.As(err, &myErr) {
		return myErr, true
	}
	return nil, false
}

// 10. FieldName возвращает имя поля из *FieldError.
//
// TODO: ошибка может быть обёрнута несколько раз. Для другой ошибки
// или nil верните пустую строку.
func FieldName(err error) string {
	var myErr *FieldError
	if errors.As(err, &myErr) {
		return myErr.Field
	}
	return ""
}

// 11. FirstError возвращает первую ненулевую ошибку из списка.
//
// TODO: порядок ошибок должен сохраняться. Если ошибок нет, верните nil.
func FirstError(errs []error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// 12. CountErrors считает количество ненулевых ошибок.
//
// TODO: nil-элементы не должны попадать в результат.
func CountErrors(errs []error) int {
	count := 0
	for _, err := range errs {
		if err != nil {
			count++
		}
	}
	return count
}

// 13. ErrorText возвращает текст ошибки.
//
// TODO: для nil верните пустую строку.
func ErrorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// 14. ValidateUser проверяет поля пользователя в порядке: name, email, age.
//
// TODO: пустые name и email считаются ошибкой обязательного поля,
// age должен быть положительным. Верните только первую найденную ошибку.
func ValidateUser(name, email string, age int) error {
	err := RequireText("name", name)
	if err != nil {
		return err
	}
	err = RequireText("email", email)
	if err != nil {
		return err
	}
	err = ValidatePositive("age", age)
	if err != nil {
		return err
	}
	return nil
}

// 15. Classify возвращает короткую категорию ошибки.
//
// TODO: nil -> "none"; ErrNotFound в цепочке -> "not_found";
// *FieldError -> "field:<имя поля>"; любая другая ошибка -> "other".
func Classify(err error) string {
	if err == nil {
		return "none"
	}
	if errors.Is(err, ErrNotFound) {
		return "not_found"
	}
	var myErr *FieldError
	if errors.As(err, &myErr) {
		return "field:" + myErr.Field
	}
	return "other"
}
