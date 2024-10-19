# go-workshop

### To run this program, use the below command
`go run main.go`

## To use custom package
1. Copy all your functions to the seperate folder and new go file. Example: basics/basics.go
2.  Rename the package as per the folder name that you created
Example:
```
package basics

func RunBasics(){
    // implementation
}
```
3. Run the below command to create a new module
`go mod init go-workshop`
4. Import this package in main function. Use the functions with `package_name.FunctionName()`
```
package main

import (
    "go-workshop/basics"
)

func main(){
    basics.RunBasics()
}
```
5. Run your program
`go run main.go`