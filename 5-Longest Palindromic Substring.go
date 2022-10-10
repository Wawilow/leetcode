package main

import (
	"fmt"
	"strings"
)

func is_palindromic(originalString string) bool {
	reversString := ""

	list := strings.SplitAfter(originalString, "")

	for i := len(list) - 1; i >= 0; i-- {
		reversString = reversString + list[i]
	}
	return originalString == reversString
}

func longestPalindrome(s string) string {
	for i := len(s); i >= 1; i-- {
		for j := 0; j <= len(s)-i; j++ {
			//fmt.Println(s[j : j+i])
			if is_palindromic(s[j : j+i]) {
				return s[j : j+i]
			}
		}
	}
	return ""
}

func main() {
	//fmt.Print(is_palindromic("aba"))
	fmt.Print(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
