package list

import (
	"errors"
	"fmt"
)

//定义两个常量一个用来存默认容量，一个用来存初始下标
const defaultCapacity = 10
const defaultSize = -1

//动态数组结构
//size	当前最后一个元素下标位置
//capacity	数组最大容量
//data	数组
type Array struct {
	size     int
	capacity int
	data     []interface{}
}

//创建一个新的动态数组
func NewArray(capacity int) *Array {
	data := make([]interface{}, capacity)
	return &Array{defaultSize, capacity, data}
}

//在不输入最大容量时创建
func NewArrayWithoutNoCap() *Array {
	data := make([]interface{}, defaultCapacity)
	return &Array{defaultSize, defaultCapacity, data}
}

//扩展数组
func (array *Array) grow(num int) (bool, []interface{}) {
	data := make([]interface{}, num)
	//将原数组中的值赋给新扩充容量的数组
	for i, elem := range array.data {
		data[i] = elem
	}
	return true, data
}

//检查数组容量是否足够
func (array *Array) check(num int) bool {
	//若存储大小加新数组的大小大于最大容量则进行扩充
	//这里只对数组进行扩充
	if array.size+num+1 > array.capacity {
		//定义新数组最大容量，这里是原数组长度加新元素长度再加一个元素的最大容量
		newCap := array.size + num + 2
		ok, array1 := array.grow(newCap)
		if ok == true {
			array.data = array1
			array.capacity = array.size + num + 2
			return true
		}
	}
	return false
}

//打印数组
func (array Array) Print() {
	for i := 0; i <= array.size; i++ {
		fmt.Print(array.data[i])
		fmt.Print("\t")
	}
}

//向list末尾加入元素，先检查容量是否足够再进行添加
func (array *Array) Add(obj ...interface{}) error {
	array.check(len(obj))
	for _, elem := range obj {
		array.size++
		array.data[array.size] = elem
	}
	return errors.New("no error")
}

//向指定位置加入元素,要判断容量是否足够且目标位置存在再进行添加
func (array *Array) Insert(location int, obj interface{}) error {
	if location <= 0 || location > array.size+1 {
		return errors.New("下标超出")
	}
	array.check(1)
	for i := array.size; i >= location-1; i-- {
		array.data[i+1] = array.data[i]
	}
	array.data[location-1] = obj
	return errors.New("no error")
}

//向指定位置修改元素，要判断目标位置是否存在
func (array *Array) Set(location int, obj interface{}) error {
	if location <= 0 || location > array.size+1 {
		return errors.New("下标超出")
	}
	array.data[location-1] = obj
	return errors.New("success")
}

//是否存在某元素，进入循环进行遍历存在返回true，否则返回false
func (array Array) Contain(obj interface{}) bool {
	for _, i := range array.data {
		if i == obj {
			return true
		}
	}
	return false
}

//是否为空，判断size如果为-1就说明数组中没有元素返回true,否则返回false
func (array Array) IsEmpty() bool {
	if array.size == -1 {
		return true
	}
	return false
}

//查看某一位置上的元素,判断目标位置是否在范围内
func (array *Array) Get(location int) (interface{}, error) {
	if location <= 0 || location > array.size+1 {
		return nil, errors.New("下标超出")
	}
	return array.data[location], errors.New("no error")
}

//判断是否相等,一个元素一个元素进行比较，其中有不同的就返回false,否则返回true,其中只有最大下标相同才会进行比较，最大下标不同直接返回false
func (array *Array) Equals(array1 *Array) bool {
	if array.size == array1.size {
		for i, elem := range array.data {
			if array1.data[i] != elem {
				return false
			}
		}
		return true
	}
	return false
}

//转换为Slice类型，data就是slice所以直接赋值过去
func (array Array) ToSlice() []interface{} {
	return array.data
}

//输出当前list的长度
func (array Array) Size() int {
	return array.size + 1
}
