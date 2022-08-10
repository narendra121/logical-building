/*
user want to drew the lottery for each group.

for eg:- user will decide the lottery number and will ask the people in group to choose one number.
after people choose one number each we need to let the user to know howmany peoples prediction is correct.

if none of them predicted correct number we need to find the number of people who predictd the number nearest to the lottery number.
eg :- in input we give 3 params lotteryNumber,NumberOfPeopleInGroup,PeoplesPredictedNumbers

input:-
1 4  3,4,5,6
response:-
1

input:-
1 4  3,5,5,6
response:-
1

input:-
4 4 2,3,3,5


response:-
3


go run other/lottery.go 4 4 2,3,3,5
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pickedNum := make([]int, 0)
	if len(os.Args) < 4 {
		fmt.Println("invalid input")
		return
	}
	lotterNumStr := os.Args[1]
	noOfPeopleStr := os.Args[2]
	pickedNStr := os.Args[3]
	lotterNum, _ := strconv.Atoi(lotterNumStr)
	noOfPeople, _ := strconv.Atoi(noOfPeopleStr)
	for _, val := range strings.Split(pickedNStr, ",") {
		temp, _ := strconv.Atoi(val)
		pickedNum = append(pickedNum, temp)
	}

	fmt.Println(getTheWinnerCount(lotterNum, noOfPeople, pickedNum))
}

func getTheWinnerCount(lotterNum, noOfPeople int, pickedNum []int) int {
	normal := make([]int, 0)
	var foundCount int
	for _, val := range pickedNum {
		temp := 0
		if val > lotterNum {
			temp = val - lotterNum
		} else if val < lotterNum {
			temp = lotterNum - val
		} else if val == lotterNum {
			foundCount += 1
		}
		normal = append(normal, temp)
	}

	if foundCount > 0 {
		return foundCount
	}
	sort.Ints(normal)
	first := normal[0]
	for i := 1; i < len(normal); i++ {
		if first < normal[i] {
			foundCount += 1
			return foundCount
		} else if first == normal[i] {
			foundCount += 1
		}
	}
	return foundCount
}
