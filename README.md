# IsWadmin [![Build Status](https://travis-ci.org/dlion/IsWadmin.svg?branch=master)](https://travis-ci.org/dlion/IsWadmin)
Check if the process is running as administrator on Windows

## How to use it

```go
package main

import (
	"fmt"
	. "github.com/dlion/IsWadmin"
)

func main() {
	if IsWadmin() {
		fmt.Println("This process is running as administrator on Windows")
	}
}
```

## Test

```
go test
```

## Author

Domenico Luciani

https://domenicoluciani.com

## License

MIT