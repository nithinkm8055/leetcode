package main

import (
	"DSA/Easy"
	"DSA/Medium"
	"fmt"
)

func main() {

	//test()

	//twoSumRun()
	//isPalindrome()
	//romanToInt()
	//longestCommonPrefix()
	//validParentheses()
	//mergeTwoLists()
	//addTwoNumbers()
	//removeDuplicates()
	//StrStr()
	//maxSubArray()
	//lengthOfLastWord()
	//plusOne()
	//addBinary()
	//lengthOfLongestSubstring()
	//merge()
	//climbingStairs()
	//fmt.Println(Hard.Trap([]int{4,2,0,3,2,5}))
	//[[12,13,1,12],[13,4,13,12],[13,8,10,12],[12,13,12,12],[13,13,13,13]]
	//fmt.Println(Hard.TrapRainWater([][]int{{12, 13, 1, 12}, {13, 4, 13, 12}, {13, 8, 10, 12}, {12, 13, 12, 12}, {13, 13, 13, 13}})) //14
	//fmt.Println(Hard.TrapRainWater([][]int{{5, 5, 5, 1}, {5, 1, 1, 5}, {5, 1, 5, 5}, {5, 2, 5, 8}})) //3
	//fmt.Println(Hard.TrapRainWater([][]int{{9, 9, 9, 9, 9, 9, 8, 9, 9, 9, 9}, {9, 0, 0, 0, 0, 0, 1, 0, 0, 0, 9}, {9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9}, {9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9}, {9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}})) //215
	//fmt.Println(Hard.TrapRainWater([][]int{{78,16,94,36},{87,93,50,22},{63,28,91,60},{64,27,41,27},{73,37,12,69},{68,30,83,31},{63,24,68,36}})) //44
	//SortArrayByParity()
	fmt.Println(Medium.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func test() {

	//s := "abc"
	//freq := make(map[rune]int)
	//freq[rune(s[0])] = 1
	//freq['b'] = freq['b']
	//
	//fmt.Println(freq['b'])
}

func SortArrayByParity() {
	Easy.SortArrayByParity([]int{1, 0, 3, 5})
}

func lengthOfLongestSubstring() {
	s := "abcaadb"
	fmt.Println(Medium.LengthOfLongestSubstring(s))
}

func addBinary() {
	num1 := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	num2 := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	//num1 := "1010000010010011011001000001010111101101100110"
	//num2 := "11010100101110111000111110011000101010000110101110"
	fmt.Println(Easy.AddBinary(num1, num2))
}

func plusOne() {
	nums := []int{0}
	fmt.Println(Easy.PlusOne(nums))
}

func lengthOfLastWord() {
	s := "Hello World 32"
	fmt.Println(Easy.LengthOfLastWord(s))
}

func maxSubArray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(Easy.MaxSubArray(nums))
}

func StrStr() {
	fmt.Println(Easy.StrStr("mississippi", "issippi"))
	//fmt.Println(Easy.StrStr("hello", "oo"))
}

func removeDuplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(Easy.RemoveDuplicates(nums))
}

func addTwoNumbers() {
	//{2,4,3} {5,6,4}  {9,9,9,9,9,9,9} {9,9,9,9}
	vals1 := []int{2, 1}
	vals2 := []int{2}

	l1 := &Medium.ListNode{}
	l2 := &Medium.ListNode{}

	param1 := l1
	param2 := l2

	for i := 0; i < len(vals1); i++ {
		l1.Val = vals1[i]

		if i != len(vals1)-1 {
			l1.Next = &Medium.ListNode{}
			l1 = l1.Next
		}
	}

	for i := 0; i < len(vals2); i++ {
		l2.Val = vals2[i]

		if i != len(vals2)-1 {
			l2.Next = &Medium.ListNode{}
			l2 = l2.Next
		}
	}

	output := Medium.AddTwoNumbers(param1, param2)

	for output != nil {
		fmt.Println(output.Val)
		output = output.Next
	}
}

func mergeTwoLists() {
	l3 := Easy.ListNode{
		Val:  4,
		Next: nil,
	}

	l2 := Easy.ListNode{
		Val:  2,
		Next: &l3,
	}

	l1 := Easy.ListNode{
		Val:  1,
		Next: &l2,
	}

	l6 := Easy.ListNode{
		Val:  4,
		Next: nil,
	}

	l5 := Easy.ListNode{
		Val:  3,
		Next: &l6,
	}

	l4 := Easy.ListNode{
		Val:  1,
		Next: &l5,
	}

	l7 := Easy.MergeTwoLists(&l1, &l4)

	for l7.Next != nil {
		fmt.Println(l7.Val)
		l7 = l7.Next
		if l7.Next == nil {
			fmt.Println(l7.Val)
		}
	}
}

func validParentheses() {
	s := "()[]{}"
	//s := "{[]}"
	fmt.Println(Easy.IsValid(s))
}

func longestCommonPrefix() {
	strs := []string{"flower", "flow", "flight", "dog"}
	//strs := []string{"dog","car","race"}
	fmt.Println(Easy.LongestCommonPrefix(strs))
}

func romanToInt() {
	s := "XXVIII"
	fmt.Println(Easy.RomanToInt(s))
}

func isPalindrome() {
	num := 1223221
	fmt.Println(Easy.IsPalindrome2(num))
}

func twoSumRun() {
	nums := []int{3, 3, 3}
	target := 6
	//sum := Easy.TwoSum(nums, target) // bruteforce
	sum := Easy.TwoSum2(nums, target) // optimized approach
	fmt.Println(sum)
}
