```go
import "fmt"
//NOTE -  (LEARNING) WHY IS THIS BIT SHIFTING DONE IN STDLIB COLOR FUNCTION(S)
func 

for i:=-2;i<30;i+=3{
	x:=uint32(0)
	y:=uint32(i)
	x=y<<8
	fmt.Printf("%d\n",x<<4)
}
```

```go
//NOTE: (LEARNING) INTS WRAP AT LIMIT
import "fmt"

for i:=250;i<300;i++{
	u:= uint8(i)
	fmt.Printf("[%d->%d]",i,u)
	if i%16==0{
		fmt.Printf("\n")
	}
}
```

```go
// Testing division
import "fmt"
fmt.Printf("7/3 = %d\n",7 / 3)
fmt.Printf("7%%3 = %d\n",7 % 3)
b1:=0b00010110
b2:=0b01000100
b3:=b2/b1
fmt.Printf("%08b (%d) / %08b (%d) = %08b (%d)",b2,int(b2),b1,int(b1),b3,int(b3))
```

```go
// Testing Byte signing
import "fmt"

x := 0b00_01_01_10
y := 0b00_00_10_01
z := x - (0b00000011*y)
fmt.Printf("%08b-%08b = %08b\n",x,y,x-y)
fmt.Printf("%d-%d=%d\n",x,y,x-y)
fmt.Printf("%08b-(0b11*%08b)=%08b\n",x,y,z)
fmt.Printf("|%08b-(0b11*%08b)|=%08b\n",x,y)
fmt.Printf("%d-(3*%d)=u8 %d\n",x,y,uint8(z))
```

```go
// looking at math rand v2
import "math/rand/v2"
import "fmt"
u:= rand.uint()
fmt.Printf("rand val: %d\n",u)
```