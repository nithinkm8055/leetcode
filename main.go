package main

import (
	"DSA/Easy"
	"fmt"
)

func main(){

	test()

	twoSumRun()
	//isPalindrome()
	//romanToInt()
	//longestCommonPrefix()
	//validParentheses()
}


func test(){
	//s := "abcde"
	//resultSlice := make([]string, 0)
	//resultSlice = append(resultSlice, "1")
	//resultSlice = append(resultSlice, "2")
	//fmt.Println(len(resultSlice), cap(resultSlice), resultSlice)

}

func validParentheses(){
	s := "()[]{}"
	//s := "{[]}"
	fmt.Println(Easy.IsValid(s))
}

func longestCommonPrefix(){
	strs := []string{"flower","flow","flight"}
	//strs := []string{"dog","car","race"}
	fmt.Println(Easy.LongestCommonPrefix(strs))
}

func romanToInt() {
	s := "XXVIII"
	fmt.Println(Easy.RomanToInt(s))
}


func isPalindrome(){
	num := 1223221
	fmt.Println(Easy.IsPalindrome2(num))
}

func twoSumRun(){
	nums := []int{3,3,3}
	target := 6
	//sum := Easy.TwoSum(nums, target) // bruteforce
	sum := Easy.TwoSum2(nums, target) // optimized approach
	fmt.Println(sum)
}

