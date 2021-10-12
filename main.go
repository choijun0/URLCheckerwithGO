package main

import (
  "fmt"
  "time"
)


func main() {
  channel := make(chan string); //type (if type is chan what is the datatype you bring to channel)
  members := [5]string{"jun", "young", "choi", "chang", "hoon"}
  for _, member := range members{
    go count(member, 3, channel);
  }
  //watch receive , 현재실행중인 goroutine에서 값을 받기를 기다리고있는 상태(값이 전달되기 전에 main함수는 종료되지 않는다. blcoking operation)
  //goroutine이 몇개인지 알고있다 따라서 channel이받아야할 값은 goroutine의 개수를 넘지 못한다.
  //Activate receivers
  for i:=0; i<len(members); i++{
    fmt.Println(<- channel);  //blocking operation == waiting
    //값을 받을때가지 진행x first in first
  }
}

//func
func count(name string, howmany int, channel chan string){
  for i:=0;i<howmany;i++{
    fmt.Println(name, i);
    time.Sleep(time.Second);
  }
  channel <- name + " is Finished!";
}