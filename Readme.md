# Go Essentials

- understanding the **Key Components** of a Go program
- Working with **Values & Types**
- Creating & Executing **Funcitons**
- Controlling Execution with **Control Structures**

"" or `` dont ''

menjalakan file:

```sh
go run pointers.go
```

```sh
go mod init example.com/first-app


go build

 ./first-app

```

```sh
go run .

go run nama-file

```

- harus di bungkus dengan funciton

```
Null Values
All Go value types come with a so-called "null value" which is the value stored in a variable if no other value is explicitly set.

For example, the following int variable would have a default value of 0 (because 0 is the null value of int, int32 etc):

var age int // age is 0
Here's a list of the null values for the different types:

int => 0

float64 => 0.0

string => "" (i.e., an empty string)

bool => false
```

```
package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100, float64(years))
	fmt.Print(futureValue)
}


```

```
package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount,years float64 = 1000, 10
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, float64(years))
	fmt.Print(futureValue)
}

```

# Formated Print

# Function

## untuk install third party library

```sh
go get

```
