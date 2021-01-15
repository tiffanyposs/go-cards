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