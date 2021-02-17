# Deck of Cards - Go

https://play.golang.org/

## Language Types

Go is a **statically typed language**.

### **Static Types**

The interpreter cares which value we are assigning to a variable.

* C++
* Java
* Go

```go
...

// this would error in Go
var test string = 5 

...
```

#### Basic Go Types

* `bool` - true/false
* `string` - "hi"
* `int` - 0, -10000, 99999
* `float64` - 10.00001, 0.0008, -100.003

### **Dynamic Types**

The interpreter does not care which value we are assigning to a variable.

* Javascript
* Ruby
* Python


```javascript
...

// this would not error in Javascript
var test = 5 ;
test = 'string';

...
```

## Variable Declaration

The most basic way to declare a variable in Go can be seen below. This way of declaring a variable can be done inside or outside of a function

```go

var card string = "Ace of Spades"

```

The shorthand of this can only be done inside of a function:

```go

func main() {
  card := "Ace of Spades"
}

```

You cannot reassign a variable to a different type. The below example would error:

```go

func main() {
  card := "Ace of Spades"
  card = 5 // error
}

```

You have to declare what type a function returns

```go

func newCard() string {
	return "Five of Diamonds"
}

```

## Lists

Go has two type of lists:

* `array` - fixed length list of things
* `slice` - an array that can shrink or grow


### Slice

Every element in a `slice` must be of the same type

#### Declaring a Slice

You must declare the type of the elements within a slice.

```go

func main() {
  cards := []string{"Ace of Diamonds", "Six of Spades"}
  ...
}

```

#### Appending New Element

You can use the `append` method to add a new element to the end of a `slice`.

```go

func main() {
  cards := []string{"Ace of Diamonds", "Six of Spades"}
  cards = append(cards, "King of Hearts")
  ...
}

```

#### Loops

```go

func main() {
  cards := []string{"Ace of Diamonds", "Six of Spades"}
  cards = append(cards, "King of Hearts")
  for index, card := range cards {
    // do stuff with index and card
  }
}

```

#### Accessing Indexes

You can use syntax like `[0:2]` or `[1:4]` or `[:2]` or `[2:]` to access ranges of indexes. The first number is the starting position and is includes, the second number is the ending position and is not included.


```go
...

cards := []string{"Ace of Diamonds", "Six of Spades", "Queen of Hearts"}

hand := cards[:2] // would give me the cards indexes 0 through 1
hand2 := cards[1:3] // would give me the cards indexes 1 through 2

```

## OO vs Go Approach

### OO Approach

* Declare a new type and set it to whatever existing type you like (slice of strings in this case)
* add a method to the type `deck` called `print`, the receiver takes two variables: `d` which is similar to `this` in other languages and `deck` which is the type you're adding the method to. By convention the first variable is a 1 or 2 letter variable matching the second variable's starting letters.
* Init an instance of the type, setting it to a variable, then you can access the methods.

```go
...

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

...

func main() {
  cards := deck{"Ace of Diamonds", "Six of Spades"}
  cards.print()
}

```

## Functions

* Has a receiver of `d` with type **deck**
* Function takes one argument `handSize` with type **int**
* It returns two things, both with type **deck**

```go

type deck []string

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

...

func main() {
  cards := newDeck()
  hand, remainingDeck := deal(cards, 5)
}

```

## Packages

You can see available packages in Go here: https://golang.org/pkg/

### Write File

You can use the package `ioutil` to write files to your hard drive, specifically the WriteFile function https://golang.org/pkg/io/ioutil/#WriteFile

* It takes a filename with type **string**
* It takes data which is a **slice** of **bytes** (aka a "byte slice" in Go)
* It also takes a perm (short for permissions)

```go

func WriteFile(filename string, data []byte, perm os.FileMode) error

```

### Byte slice

A byte slice is a slice of numbers that coorispond to a string. For example the string `Hi there!` as a byte slice looks like this: `[72 105 32 116 104 101 114 101 33]`. This is just a computer fiendly way to write a string.

You can see at http://www.asciitable.com/ a table that you can look up bytes to letters.



### Type Conversion

You can change a type in Go 

`<type-you-want>(<value-you-have>)`

This example will covert "Hello there!" into a **byte slice**:

```go

greeting := "Hi There!"
greetingBytes := []byte(greeting)

```

### Testing In Go

Always end with `_test.go`
