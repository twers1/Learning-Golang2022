// Structures Структуры

/*
Go does not support classes. Instead, it has structures.
Structures are a type of data that are collections of fields that can be used to group different data together.

Go не поддерживает классы. Вместо этого в нем есть структуры.
Структуры - это тип данных, который представляет собой коллекции полей, с помощью которых можно группировать различные данные вместе.
*/

type Contact struct {
  name string
  age  int
} 

/*
Our Contact structure has two fields, a string and an integer.
Now we can create a new instance of our Contact structure using the following syntax:
x is now a structure that is initialized with the data that is specified in curly brackets.

Наша структура Contact имеет два поля, строку и целое число.
Теперь мы можем создать новый экземпляр нашей структуры Contact, используя следующий синтаксис:
x теперь является структурой, которая инициализируется с данными, которые указаны в фигурных скобках.
*/

x := Contact{"Андрей", 42}

// Task 1 Задание 1 (You need to create a Car structure, which should have three fields: color and brand of lowercase type, and year of integer type.)Вам нужно создать структуру Car, которая должна иметь три поля: color и brand строчного типа, а также year целочисленного.

type Car struct {
  color string
  brand string
  year  int
} 

// We can access the fields of a structure using its name and a point: 
//Мы можем получить доступ к полям структуры, используя её имя и точку:

x := Contact{"Джеймс", 42}

x.age = 8
fmt.Println(x.age)  // выведет 8
fmt.Println(x.name) // выведет Джеймс

// Task 2 Задание 2 Which will take out что выведет

type Coord struct {
  x int
  y int
}

func main() {
  a := Coord{7, 5}
  fmt.Println(a.x - a.y)
} // 2

// Pointers to structures Указатели на структуры

// Like simple pointers, we can also create pointers to structures using the & operator:
// Подобно простым указателям, мы также можем создавать указатели на структуры с помощью оператора &:

x := Contact{"Джеймс", 42}
p := &x

// Pointers to structures are automatically dereferenced, which means that we can access field values simply by using a point:
// Указатели на структуры автоматически разыменовываются, что означает, что мы можем получить доступ к значениям полей, просто используя точку:

x := Contact{"Джеймс", 42}
p := &x
fmt.Println(p.age)

// We can also use pointers when creating a new structure: Now p is a pointer to the newly created instance of the Contact structure.
// Мы также можем использовать указатели при создании новой структуры: Теперь p - это указатель на только что созданный экземпляр структуры Contact.

p := &Contact{"Джон", 15}



