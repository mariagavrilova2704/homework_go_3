package structs

import "fmt"

type User struct {
	ID     int
	Name   string
	Active bool
}

type Product struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Amount int
	Paid   bool
}

// 01. NewUser создаёт пользователя с ID и именем.
// TODO: Active должен получить zero value.
func NewUser(id int, name string) User {
	return User{
		ID:     id,
		Name:   name,
		Active: false,
	}
}

// 02. RenameUser возвращает копию пользователя с новым именем.
// TODO: ID и Active должны сохраниться.
func RenameUser(user User, name string) User {
	user.Name = name
	return user
}

// 03. ActivateUser возвращает активную копию пользователя.
// TODO: остальные поля не изменяйте.
func ActivateUser(user User) User {
	user.Active = true
	return user
}

// 04. DeactivateUser возвращает неактивную копию пользователя.
// TODO: остальные поля не изменяйте.
func DeactivateUser(user User) User {
	user.Active = false
	return user
}

// 05. IsActive возвращает состояние пользователя.
// TODO: прочитайте соответствующее поле структуры.
func IsActive(user User) bool {
	return user.Active
}

// 06. EmptyUser возвращает zero value структуры User.
// TODO: не заполняйте поля вручную значениями, отличными от zero value.
func EmptyUser() User {
	return User{
		ID:     0,
		Name:   "",
		Active: false,
	}
}

// 07. SameUserID сравнивает пользователей только по ID.
// TODO: имя и Active не должны влиять на результат.
func SameUserID(a, b User) bool {
	if a.ID == b.ID {
		return true
	}
	return false
}

// 08. UserLabel формирует строку "<ID>:<Name>".
// TODO: между двоеточием и именем пробел не нужен.
func UserLabel(user User) string {
	return fmt.Sprintf("%d:%s", user.ID, user.Name)
}

// 09. NewProduct создаёт товар с названием и ценой.
// TODO: перенесите оба аргумента в соответствующие поля.
func NewProduct(name string, price int) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}

// 10. ChangePrice возвращает копию товара с новой ценой.
// TODO: название товара должно сохраниться.
func ChangePrice(product Product, price int) Product {
	product.Price = price
	return product
}

// 11. ProductTotal считает стоимость count единиц товара.
// TODO: функция не должна изменять product.
func ProductTotal(product Product, count int) int {
	return count * product.Price
}

// 12. ApplyDiscount возвращает товар с уменьшенной ценой.
// TODO: percent находится в диапазоне от 0 до 100. Результат вычисляется
// целочисленной арифметикой, название товара сохраняется.
func ApplyDiscount(product Product, percent int) Product {
	//product.Price = product.Price - (product.Price * percent / 100) -формула, которую я помню
	product.Price = product.Price * (100 - percent) / 100
	return product
}

// 13. NewOrder создаёт заказ с ID и суммой.
// TODO: Paid должен получить zero value.
func NewOrder(id, amount int) Order {
	return Order{
		ID:     id,
		Amount: amount,
		Paid:   false,
	}
}

// 14. MarkPaid возвращает оплаченную копию заказа.
// TODO: ID и Amount должны сохраниться.
func MarkPaid(order Order) Order {
	order.Paid = true
	return order
}

// 15. OrderStatus возвращает "paid" для оплаченного заказа
// и "pending" для неоплаченного.
// TODO: выберите строку по полю Paid.
func OrderStatus(order Order) string {
	if order.Paid {
		return "paid"
	}
	return "pending"
}
