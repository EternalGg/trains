package test

import "fmt"

func Test() {
	////	m := sync.Mutex{}
	//wg := sync.WaitGroup{}
	//
	//array, count := []int{}, 6
	//for i := 0; i < count; i++ {
	//	array = append(array, i)
	//}
	//fmt.Println(array)
	//for i := 0; i < count; i++ {
	//	array[i]++
	//}
	//fmt.Println(array)
	//for key, value := range array {
	//	value++
	//	array[key]++
	//	//m.Lock()
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		array[key]++
	//
	//	}()
	//	//m.Unlock()
	//}
	//wg.Wait()
	//fmt.Println(array)

}

func Mix() {
	A()
	B()
}

func A() {
	fmt.Println("A")
}
func B() {

	fmt.Println("B")
	C()
}
func C() {

	fmt.Println("C")
	A()
}
