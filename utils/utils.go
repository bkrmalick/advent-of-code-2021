package utils

import (
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"sync"
)


type Stack []interface{}

func (s *Stack) Push(item interface{}){
	*s = append(*s, item)
}
func (s *Stack) Pop() interface{}{
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}
func (s *Stack) IsEmpty() bool{
	return len(*s) == 0
}

func SetBasePathToCurrentDir(){
	_, filename, _, _ := runtime.Caller(1) // skip one function in the stack to ensure we get dir of the calling file and not the utils
	err := os.Chdir(path.Dir(filename))
	HandleError(err, "updating current working dir")
}

func HandleError(err error, action string) {
	if err!= nil {
		log.Fatalf("Error while %s: %s", action, err)
	}
}

func Binary2Int(binary string) int64 {
	n, err := strconv.ParseInt(binary, 2, 64)
	HandleError(err, "trying to convert binary to int")
	return n
}

func String2Int(s string) int {
	n, err := strconv.Atoi(s)
	HandleError(err, "trying to convert string to int")
	return n
}

func Int2String(i int) string {
	s := strconv.Itoa(i)
	return s
}

func Min(ls... int) int {
	min := ls[0]
	for _,x := range ls{
		if x<min{
			min = x
		}
	}
	return min
}

func Max(ls... int) int {
	max := ls[0]
	for _,x := range ls{
		if x>max{
			max = x
		}
	}
	return max
}

func InListInt(x int, ls []int) bool {
	for _,v := range ls{
		if v==x{
			return true
		}
	}
	return false
}


// InList generic in list because golang does not support type parameters
func InList(x interface{}, ls interface{}) bool {
	list := reflect.ValueOf(ls)
	for i:=0;i<list.Len();i++{
		if list.Index(i).Interface()==x{
			return true
		}
	}
	return false
}

func IndexOf(x interface{}, ls interface{}) int {
	list := reflect.ValueOf(ls)
	for i:=0;i<list.Len();i++{
		if list.Index(i).Interface()==x{
			return i
		}
	}
	return -1
}

// StringIntersect intersection using hash map which allows it to operate in O(n) time
// e.g utils.StringIntersect("abcd", "xyzad") returns "ad"
// TODO generalise to interface{} if needed
func StringIntersect(s1 string, s2 string) string {
	m := make(map[string]bool) // <character, bool_indicating_if_it_exists>
	intersection := ""

	for _, c := range s1 {
		char := string(c)
		m[char] = true
	}

	for _, c := range s2 {
		char := string(c)
		if m[char] {
			intersection+=char
		}
	}

	return intersection
}

func StringUnion(s1 string, s2 string) string {
	m := make(map[string]bool) // <character, bool_indicating_if_it_exists>
	union := ""

	for _, c := range s1 {
		char := string(c)
		m[char] = true
	}

	for _, c := range s2 {
		char := string(c)
		m[char] = true
	}

	for k,v := range m {
		if v {
			union+=k
		}
	}

	return union
}


// treats strings as sets, checks if all characters in substr are present in s
func IsSubset(s string, substr string) bool {
	m := make(map[string]bool)
	for _, char := range s {
		m[string(char)] = true
	}
	for _, char := range substr {
		if !m[string(char)] {
			return false
		}
	}
	return true
}


func StringSliceToCSV(ls []string) string{
	s:=""
	for _,v := range ls {
		s+=","+v
	}
	return s[1:]
}

// CloseChannelWhenDone waits for a wg then closes the channel c
// example use:
// ````
//  go utils.CloseChannelWhenDone(wg, solutions)
// ````
func CloseChannelWhenDone(wg *sync.WaitGroup, c interface{}){
	wg.Wait()
	reflect.ValueOf(c).Close() // would usually be close(c) but we want to keep this function generic for all channel types
}
