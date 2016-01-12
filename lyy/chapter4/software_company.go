package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"errors"
)

func main() {

	om := &OrderMarket{}
	om.Init()

	cm := &CoderMarket{}
	cm.Init()

	totalMoney := 100.0 //单位：w

    K:=3
	L:=false

	for {
		if totalMoney > 1000.0 || totalMoney < 0 {
			fmt.Println("创业结束. totalMoney:", totalMoney)
			return
		}

		fmt.Println("新一轮工作开始")
		order, err := om.GetOrder()
		if err != nil {
			fmt.Println("超过90天没有获取到订单,创业失败")
			return
		}

		coder, err := cm.GetCoder(K,L)
		if err != nil {
			fmt.Println("没有找到程序员，退回订单后继续")
			om.ReturnOrder(order)
			totalMoney -= order.Money * 2
			continue
		}

		usetime:=order.GetTime+rand.Intn(300)
		if usetime > order.LeadTime {
			fmt.Println("订单完成时间超出约定!")
			totalMoney+=order.Money*2
			moneyToCoder :=coder.Salary
			coder.GetMoney(moneyToCoder)
			totalMoney += moneyToCoder

		}else {
			moneyToCoder :=coder.Salary + order.Money/10
			coder.GetMoney(moneyToCoder)
			totalMoney += moneyToCoder

		}

		laoy:=rand.Intn(1)
		coderx:=rand.Intn(1)
		if laoy==1 && coderx==1 {
			L=false
			fmt.Println("公司同意",coder.Name,"辞职")
		}else {
			K=coder.nownum
			L=true
			fmt.Println("公司不同意",coder.Name,"辞职")
		}


		software, err := coder.WriteCode(order)

		if err != nil {
			fmt.Println("开发出错，公司倒闭")
			return
		}
		om.DeliverOrder(order, software)
		totalMoney += order.Money

		/*moneyToCoder := order.Money / 10.0
		coder.GetMoney(moneyToCoder)
		totalMoney -= moneyToCoder
*/

	}
}

type Order struct {
	Money float64 //w
	Name  string
	LeadTime int
	GetTime int
}

type OrderMarket struct {
	orderPool []Order
}

func (om *OrderMarket) Init() {

	rand.Seed(time.Now().Unix())
	var orderPool []Order
	for i := 0; i < 20; i++ {
		order := Order{
			Money: float64(rand.Int()%50 + 50),
			Name:  "order:" + strconv.Itoa(i),
			LeadTime: rand.Intn(185)+180,
			GetTime: rand.Intn(112),
		}
		orderPool = append(orderPool, order)
	}

	om.orderPool = orderPool
}

func (om *OrderMarket) GetOrder() (order Order, err error) {

		i := rand.Intn(19)
		order = om.orderPool[i]
	if order.GetTime>90 {
		err:=errors.New("TimeOut!")
		return order,err
	}else {
		om.orderPool = append(om.orderPool[0:19],om.orderPool[0:19]...)
		fmt.Printf("get an order:%v money:%vw leadtime:%v\n", order.Name, order.Money,order.LeadTime)
		return
	}
}

func (om *OrderMarket) ReturnOrder(order Order) {
	om.orderPool = append(om.orderPool, order)
	fmt.Println("ret an no complete order:", order.Name)
}

func (om *OrderMarket) DeliverOrder(order Order, sf SoftWare) {
	fmt.Println("deliver order success. software info:", sf.Name)
}

type CoderMarket struct {
	coderPool []GoCoder
}

func (cm *CoderMarket) Init() {

	gc := GoCoder{
		Name: "xiaomin",
		Salary:rand.Float64()*50+50,
	}
	cm.coderPool = append(cm.coderPool, gc)

	gc = GoCoder{
		Name: "laoA",
		Salary:rand.Float64()*80+100,
	}
	cm.coderPool = append(cm.coderPool, gc)

}

func (cm *CoderMarket) GetCoder(i int,flag bool) (coder GoCoder, err error) {
	index := rand.Int() % len(cm.coderPool)
	coder.nownum = index
	if flag {
		coder = cm.coderPool[i]
		coder.nownum = i
		fmt.Println("success get a order.")
		return
	}else {

		coder = cm.coderPool[index]
		fmt.Println("success get a order.")
		return
	}
}

type SoftWare struct {
	Name string
}

type Coder interface {
	WriteCode(order Order) (SoftWare, error)
	GetMoney(money float64)
}

type GoCoder struct {
	Name string
	Salary float64
	nownum int
}

func (gc *GoCoder) WriteCode(order Order) (sf SoftWare, err error) {

	fmt.Printf("begin to write code for %v\n", order.Name)
	time.Sleep(time.Second * 3)
	fmt.Println("write code ok.")

	sf.Name = fmt.Sprintf("hello, world! by %v", gc.Name)
	return
}

func (gc *GoCoder) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", gc.Name, money)
}

