# Deck of Cards - Go

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