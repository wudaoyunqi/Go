package main

import(
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var i int =0

func main(){
	file1,err := os.Open("millionRandom.txt")
	if err != nil{
		panic(err)
	}
	defer file1.Close()

	file2,err := os.Create("TheResult.txt")
	if err != nil{
		panic(err)
	}
	defer file2.Close()

	reader := bufio.NewReader(file1)
	map1 := make(map[string]int)
	for{
		line,err := reader.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
			panic(err)
		}
		line=strings.Replace(line,"\r\n","",-1)
		map1[line]++
		i++
	}

	var keys []string
	for key := range map1{
		keys = append(keys,key)
	}
	sort.Strings(keys)

	for _,k :=range keys{
		v := strconv.Itoa(map1[k])
		file2.WriteString(k)
		file2.WriteString(" : ")
		file2.WriteString(v)
		file2.WriteString("\r\n")
	}

}


