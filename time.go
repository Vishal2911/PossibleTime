package main

import "fmt"


func possibleTimes(digits []int) int {
	//this map is to check if that combination is already created
	times_exist := make(map[string]bool)

	//this time_count is used to store number of unique count possible
	time_count := 0
	for i := 0 ; i < len(digits) ;i++ {

		if digits[i] > 2 {
			continue
		}
		for j := 0 ; j < len(digits) ;j++{
			if (digits[i] == 2 && digits[j] > 4 ) || (i == j){
				continue
			}
		   for k := 0 ; k < len(digits) ;k++{
			   if (digits[k] > 6) || (j == k || k == i){
				   continue
			   }
			   for l := 0 ; l < len(digits) ;l++  {
				if (digits[k] == 6  && digits[l] > 0 ) || (l == k || l == j || l == i) || ( digits[i] == 2 && digits[j] == 4 &&(digits[k] > 0 || digits[l] > 0) ) {
					continue
				}
				possible_time := fmt.Sprintf("%d%d%d%d",digits[i], digits[j], digits[k],digits[l])
				_ , ok := times_exist[possible_time]
				if ok {
					continue
				}else {
					times_exist[possible_time] = true
					time_count +=1
				}
	
			   } 
			  
		   }
		}
	}
	return time_count
}

func main() {
	
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
}