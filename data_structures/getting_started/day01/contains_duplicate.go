package day01

//import "fmt"
//
//func main()  {
//	newArray  := []int{1,2,4,3,1}
//	result := containsDuplicate(newArray)
//	fmt.Println(result)
//}

func containsDuplicate(nums []int) bool {
	var formatArray = make(map[int]int)
	for _, value := range nums {
		_, ok := formatArray[value]
		if ok {
			return true
		}
		formatArray[value] = value
	}
	return false
}
