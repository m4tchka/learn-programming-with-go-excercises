package main

import (
	"fmt"
	"time"
)

/*
Write a program which checks if it’s time to play on your gaming PC.

You’re given an input time in the format {hour}:{minutes}:{seconds} AM/PM and you should check if the time is in one of the ranges when you can play on your PC:
07:00:00 PM - 09:00:00 PM
08:00:00 AM - 10:00:00 AM
11:00:00 PM - 06:00:00 AM

If the given time is in any of those ranges, print It’s gaming time!
If not, print It’s not the time for games yet…
*/
//cd ..;mkdir 08-gaming-time ; cd $_; go mod init $_ ; echo "package main" > main.go ; code main.go
func main() {
	now := time.Now()
	nowF := now.Format()
	fmt.Println(time.Now().Format(RFC))
}
