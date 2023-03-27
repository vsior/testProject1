//package main

// import (
// 	"errors"
// 	"fmt"
// )

// var mySignalErr = errors.New("mySignalErr LOL!")

// type MyErr struct {
// 	err     error
// 	message string
// }

// func (m MyErr) Error() string {
// 	return fmt.Sprintf("%s:%s", m.message, m.err)
// }

// func (m MyErr) Unwrap() error {
// 	return m.err
// }

// func Is_DefaultErr() error {
// 	return MyErr{
// 		err:     mySignalErr,
// 		message: "mySignalErr detect Is",
// 	}
// }

// func As_DefaultErr() error {
// 	err := Is_DefaultErr()
// 	return MyErr{
// 		err:     fmt.Errorf("add in As_DefaultErr %w", err),
// 		message: "As_DefaultErr detect As",
// 	}
// }

// func main() {

// 	err := As_DefaultErr()
// 	if err != nil {
// 		fmt.Println(errors.Is(err, mySignalErr)) //true
// 		fmt.Println(errors.As(err, &MyErr{}))    //true
// 	}

// }
