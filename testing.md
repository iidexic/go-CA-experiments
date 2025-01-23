```go

import "fmt"

for g:= range 3{
fmt.Println(g)
}
```

```go
//just checking shit
import "fmt"
var b byte = 0
b--
//floor div value
fmt.Println("floor div summed max:",(255*3)/2)
fmt.Println("fdiv then sum:", (255/2)*3, "- precedence check:", 255/2*3)
var i int = -16
fmt.Println("int to uint wraps neg:",uint(i))
fmt.Println("byte to int:",int8(byte(32)))

```

```go
//test slice modifying
import "fmt"

g:= make([]int, 300)
for x:= range g{
	g[x] = x
}

fmt.Println(g[128:138])

//? Does index of a subslice change?
//** Answer: Yes, always starts @ 0.  
for i, val := range g[225:229]{
	fmt.Println("index:",i,"value:", val)
}




```

```go
import "fmt"
g:= make([]int, 300)
for x:= range g{
	g[x] = x
}
//? Does modifying val in a loop modify original array?
//** Answer: Yes, always starts @ 0.  
for ind, vind := range g[16:19]{
	*vind *= *vind
	fmt.Println("vind:",vind,"slice:", g[16+ind])
}
```

```go
import "fmt"
var i, s int = 0,0
for i= range 30 {	
	s+=i
}
fmt.Println(i)
```

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
fmt.Println(uint(12)>>40)
//uhh neg modulo:
fmt.Println(((2%16)-20)%16)
```

```go
//loopy with returns

for i:= range 20{
	if i%3==0{
		fmt.Println(i)
		next 
	}
}
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
//pretty sure bit shifts are waaay more efficient tho.

```

```go
//? What does it take to get a float to give weird rand numbers
// Apparrently none of this
gimbley:= func() float64{
	g:=(float64(1.1)-float64(1.101)+0.001)*10000000000000.0
	g-=0.0011015494072
	g*=100000000000.001
	fmt.Printf("%0.32f\n",g)
	return g
}

x:=gimbley()
for i:= range 12{
	g:= []byte{byte(i)}
	for j:= range byte(3){
			g=append(g,j)
	} 
}
y:=gimbley()

gimb2:=gimbley
x2:=gimb2()
x=gimb2()

var gimb3 func() float64
gimb3=gimb2
gimbP:=&gimb2
x3:=gimb3()
x3p:=(*gimbP)()
_=x
_=x2
_=y
_=x3
_=x3p


```