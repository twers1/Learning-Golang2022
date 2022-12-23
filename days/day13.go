// Task 1 Задание 1 

type Dogstruct {
name string
age int
}
func main() {
x := &Dog {"Max", 8}
fmt.Println( x.age)
}

// Methods Методы 

/*
We can add functionality to our structures with methods!
Methods are simply functions with a special receiver argument.
The receiver is specified between the func keyword and the method name.
In the above example the receiver is the Contact structure.

Мы можем добавить функциональность нашим структурам с помощью методов!
Методы - это просто функции со специальным аргументом-получателем.
Получатель указывается между ключевым словом func и именем метода.
В приведенном выше примере приемником является структура Contact.
*/

func (x Contact) welcome() {
  fmt.Println(x.name)
  fmt.Println(x.age)
} 

// Example Пример

package main

import "fmt"

type Contact struct {
  name string
  age int
}

func (x Contact) welcome() {
  fmt.Println(x.name)
  fmt.Println(x.age)
}

func main() {
  x := Contact{"Джеймс", 42}
  x.welcome()
}

// Since methods are just functions with a receiver argument, we can achieve the same functionality by using a normal function that takes a structure as an argument:
// Поскольку методы - это просто функции с аргументом-получателем, мы можем достичь той же функциональности, используя обычную функцию, которая принимает структуру в качестве аргумента:

func welcome(x Contact) {
  fmt.Println(x.name)
  fmt.Println(x.age)
}
func main() {
  x := Contact{"Джеймс", 42}
  welcome(x)
}

// Task 2 Задание 2

type Test struct {
  a int
  b int
}
func (x Test) do() int {
  return x.a - x.b
}
func main() {
  t := Test{11, 2}
  fmt.Println(t.do())
} // 9

// If we need to change the structure data in a method, we can use pointers as recipients of the method.
// Если нам нужно изменить данные структуры в методе, мы можем использовать указатели в качестве получателей метода.

func (ptr *Contact) increase(val int) {
  ptr.age += val
}

// Now our increase() method uses a pointer as a recipient and can change the age field in the structure it is called on:
// Теперь наш метод increase() использует указатель в качестве получателя и может изменять поле age в структуре, на которую он вызван:

x := Contact{"Джеймс", 42}
x.increase(8)
fmt.Println(x.age) // выведет 50

// Go automatically dereferences pointers, so we just call the method with the name of the structure, just like with a simple recipient.
// Go автоматически разыменовывает указатели, поэтому мы просто вызываем метод с именем структуры, как и в случае с простым получателем.






