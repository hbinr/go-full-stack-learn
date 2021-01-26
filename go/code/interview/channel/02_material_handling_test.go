package channel

import (
	"fmt"
	"sync"
	"testing"
)

/*
	题目：材料处理

	某工厂有A、B、C三辆厨房工程车，A车上能清洗材料，B车上能加工材料，C车上能装载材料；

	三辆工程车能边行驶边清洗/加工/装载材料，每辆车上有3个工人；最初的原始材料有D1、D2、D3三种材料，
	每种材料的清洗耗时比例为6:3:3,每种材料的加工/装载耗时皆为1:1:1,每种材料的数量一致；

	材料的处理顺序为:清洗->加工->装载；车辆之间材料进行交互，需要保持比较近的相对距离

	要求：这三辆车需要将处理完的原材料，尽快送达商家手里。请问
	如何分配比较好？

	分析：
	数据流转：func A{} => func B{} => func C{}
*/

var (
	wg     sync.WaitGroup
	global sync.WaitGroup
	ch     = make(chan []int, 1)
)

// CarA A车功能函数,清洗材料
func CarA(elements []int) {
	// 存储切分后的任务
	var tasks = make([][]int, 3)

	// 3个工人
	for i := 0; i < 3; i++ {
		// 用来分配任务
		task := []int{}
		// 获取任务分割
		for _, value := range elements {
			task = append(task, value/3.0) // 每种材料的清洗耗时比例为6:3:3
		}
		wg.Add(1)
		// 开始洗车任务
		go func(task []int, i int) {
			// 循环变量i被func literalloopclosure捕获。
			tasks[i] = clean(task)
			wg.Done()
		}(task, i)
	}
	wg.Wait()

	// 将清洗后的各个材料合并，方便后续操作
	for idx, _ := range elements {
		// 清空原来的elements idx下标的数据
		elements[idx] = 0

		// 叠加每种材料的结果
		for _, task := range tasks {
			elements[idx] = task[idx]
		}
	}
	ch <- elements
	global.Done()

}

// CarB B车功能函数,加工材料
func CarB() {
	elements := []int{}
	for {
		select {
		case elements = <-ch:
			goto Label

		default:
			continue
		}
	}
	// 对每种材料进行加工
Label:
	for idx, element := range elements {
		wg.Add(1)
		go func(element, index int) {
			elements[index] = process(element)
			wg.Done()
		}(element, idx)
	}

	wg.Wait()
	global.Done()
}

// CarC C车功能函数,装载材料
func CarC() {
	elements := []int{}
	for {
		select {
		case elements = <-ch:
			goto Label
		default:
			continue
		}
	}
	// 对每种材料进行加工
Label:
	for idx, element := range elements {
		wg.Add(1)
		go func(element, index int) {
			elements[index] = load(element)
			wg.Done()
		}(element, idx)
	}

	wg.Wait()
	global.Done()
}

func clean(task []int) []int {
	fmt.Println("清洗材料")
	return task
}

func process(task int) int {
	fmt.Println("加工材料")
	return task
}

func load(task int) int {
	fmt.Println("装载材料")
	return task
}

// TestMaterialHandle 测试
func TestMaterialHandle(t *testing.T) {
	elements := []int{3.0, 3.0, 3.0}
	global.Add(1)
	go CarA(elements)

	global.Add(1)
	go CarB()

	global.Add(1)
	go CarC()

	global.Wait()
}
