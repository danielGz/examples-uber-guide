package copy_slices

import "fmt"

var LuckyNumbers = []int{2, 3, 5, 7, 11, 13}

type Lottery struct {
	nums []int
}
func (a *Lottery) SetLuckyNumbers(nums []int){
	a.nums = nums
}

func (a *Lottery) SetLuckyNumbersProperly(nums []int){
	a.nums = make([]int,len(nums))
	copy(a.nums,nums)
}

func (a *Lottery) Shuffle(){
	temp := a.nums[len(a.nums)-1]
	a.nums[len(a.nums)-1] = a.nums[0]
	a.nums[0]=temp
}

func (a *Lottery) PrintLuckyNumbers(){
	fmt.Println(a.nums)
}
