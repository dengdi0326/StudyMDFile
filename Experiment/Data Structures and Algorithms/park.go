package main

import (
	"fmt"
)

type car struct {
	pCar
	stu   string
}

type carPacking struct {
	top  int
	park [3]car
}

type pCar struct {
	num  int
	time int
}

// 临时倒车
type pCarPacking struct {
	park [3]pCar
	top  int
}

type node struct {
	data int
	next *node
}

// 通道
//type PassQueue struct {
//	CarNum int
//	front *node
//	rear   *node
//}

type PassQueue []node

var p    carPacking
var pc   pCarPacking
var pass PassQueue
var a car

func A_Car(a car) {
	if p.top < 3 {
		p.top += 1
		p.park[p.top-1].num= a.num
		p.park[p.top-1].stu = a.stu
		p.park[p.top-1].time = a.time
		fmt.Println("停车成功")
	}else {
		fmt.Println("停车场已满")
		var t node
		t.data = a.num
		t.next = nil
		//pass.rear.next = &t
		//pass.rear = &t
		pass = append(pass, t)
	}
}

func D_car(n, t int){
	// 停车场的车出去
	var num int
	for i:=0; i<3; i++ {
		if n == p.park[i].num{
			num = i
			fmt.Println("停车时间：" , (t - p.park[i].time) , "车号：" , p.park[i].num)
			break
		}
		pc.top += 1
		pc.park[i] = p.park[i].pCar
	}
	for k:=num; k > 0; k-- {
		p.park[k].pCar = pc.park[k]
		pc.park[k].time = 0
		pc.park[k].num = 0
	}
	pc.top = 0

	if len(pass) == 0 {
		fmt.Println("通道无车")
		p.top -= 1
		return
	}

	fmt.Println("通道有车")
	a := pass[0]
	pass = pass[1:]
	p.park[0].num = a.data
	p.park[0].time = t
}

func main() {

	fmt.Println("A:停车   D:移车   P:显示停车场数目  W:显示等候停车数目    E:退出")

	for {
		fmt.Scanln(&a.stu, &a.num, &a.time)
		switch a.stu{
		case "A":
			A_Car(a)

		case "D":
			D_car(a.num, a.time)
		case "P":
			fmt.Println(p.top)
		case "W":
			fmt.Println(len(pass))
		case "E":
			return
		default:
			fmt.Println("error")
		}
	}
}
