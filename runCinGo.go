package main

/* // #include <stdio.h>
// #include <sys/ioctl.h>
// #include <stdlib.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <fcntl.h>
// #include <unistd.h>
// void statusCashDrawer(void) {
//     unsigned int arr = 12;
//     int ioctl_code;

//     FILE *fp = fopen("/dev/aipdcs", "rw");
//     ioctl_code = ioctl(fileno(fp), 0xc00c7710, &arr);

//     if (ioctl_code != 0) {
//         printf("ioctl error code %d\n", ioctl_code);
//         exit(1);
//     }
//     printf("%p\n", fp);
// }
import "C"
import "fmt"

func main() {
	fmt.Println("Go run main")
	C.statusCashDrawer()
} */
