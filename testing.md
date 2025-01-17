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

```go
//bit shifts
import "fmt"
var b []byte = make([]byte,36)
for i:= range b {
	b[i]=byte(i)
	//fmt.Printf("%d ",b[i])
}
fmt.Printf("\nlen = %d\n",len(b))
for i:= range b {
	fmt.Printf("(%d -> ",b[i])
	b[i]=(b[i])>>2
	fmt.Printf("%d) ",b[i])
	if i%10==0{
		fmt.Println("")
	}
}
fmt.Println(b[24])
fmt.Printf("\nlen = %d\n",len(b))

```

```go
import "fmt"
// neg mod always neg 
fmt.Printf("neg mod: %d\n",-10%7)

//neg bitshift works exactly the same except in negative direction. i.e. if x<<y=8, -x<<y=-8
//shifting by a neg number crashes (can't do 1<<-1)
fmt.Printf("Leftshift 1: %d\n",67>>3)
fmt.Printf("Rightshift 3:%d\n",-10<<3)
```

```go
//uhh xor
b:=byte(22)
b2:=b+2*2*2*2*2*2
_=b2
fmt.Println("^2:",^2,"|^8:",^8,"|^11:",^11)
```

```go
//Re-Reprsenting Right Shift 

import 	"fmt"
import  "math"

x:=12
y:=2

//these dumb notebooks dont work with the import but this is true yes
ypow:=math.Pow(float64(2),float64(y))
fmt.Println(x>>y==int(ypow))
```