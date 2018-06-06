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
type PassQueue struct {
	CarNum int
	front *node
	rear   *node
}

var p    carPacking
var pc   pCarPacking
var pass PassQueue
var a car
var b int

func A_Car(a car) {
	if p.top < 3 {
		p.top += 1
		p.park[p.top-1].num= a.num
		p.park[p.top-1].stu = a.stu
		p.park[p.top-1].time = a.time
		fmt.Println("停车成功")
	}else {
		fmt.Println("停车场已满")
		var t *node
		t.data = a.num
		t.next = nil
		pass.rear.next = t
		pass.rear = t
		pass.CarNum += 1
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

	if pass.CarNum == 0 {
		fmt.Println("通道无车")
		p.top -= 1
		return
	}

	fmt.Println("通道有车")
	a := pass.rear.next
	pass.rear.next = a.next
	p.park[0].num = a.data
	p.park[0].time = t
	pass.CarNum -= 1
}

func main() {

	fmt.Println("1:停车   2:移车   3:显示所有信息     4:退出")

	for {
		fmt.Scan(&b)
		switch b {
		case 1:
			fmt.Scan(&a.num ,&a.time, &a.stu)
			A_Car(a)

		case 2:
			fmt.Scan(&a.num ,&a.time, &a.stu)
			D_car(a.num, a.time)
		case 3:
			fmt.Println(p.top,pass.CarNum)
		case 4: return
		}
	}
}
