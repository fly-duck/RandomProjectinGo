package main

import(
	"fmt"
	"time"
	"math/rand"
)

func f(c int ){
	for i :=0;i<20;i++{
		fmt.Println(c ,":",i )
		amt:=time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond*amt)

	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
	  c <- "ping"
	}
  }
  
  func printer(c chan string) {
	for {
	  msg := <- c
	  fmt.Println(msg)
	  time.Sleep(time.Second * 1)
	}
  }


  func ponger(c chan string) {
	for i := 0; ; i++ {
	  c <- "pong"
	}
  }
func main(){
	// go f(1)
	// var input string 
	// for i:=0; i<10 ; i++{
	// 	go f(i);
	
	// 	// fmt.Println(f(),"in coroutine",i)
	// }
	// fmt.Scanln(&input)
	// fmt.Println(input);


	// var c chan string = make(chan string)
	
	// for i:=0; i<10; i++{
	// 	go ponger(c);

	// }
	// go pinger(c)
	// go printer(c)
  
	// var input string
	// fmt.Scanln(&input)



	c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 20)
    }
  }()

  go func() {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second * 20)
    }
  }()

  go func() {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
		fmt.Println(msg2)
	case <- time.After(time.Second):
		fmt.Println("timeout")
	  default:
		// fmt.Println("Not ready yet")
      }
    }
  }()

  var input string
  fmt.Scanln(&input)
}