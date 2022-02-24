package rwfile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) *bufio.Scanner {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		//defer file.Close()  -- 这里不能defer close  close可能返回一个空
		scanner := bufio.NewScanner(file)
		return scanner
	}
}

func ReadFileAll(filename string) (text string) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	return string(content)

}
func WriteFile(filename string, context string) bool {
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	fd.Write([]byte(context))

	//var zipBuffer *bytes.Buffer = new(bytes.Buffer)
	//var zipWriter *zip.Writer = zip.NewWriter(zipBuffer)
	//var zipEntry io.Writer
	//var err error
	//zipEntry, err = zipWriter.Create(filename)
	//fmt.Println("@@@@@@@@@@@@",filename)
	//if err != nil{
	//	fmt.Println("***************************",err)
	//}
	//fmt.Println("@@@@@@@@@@@@\n\n\n\n",context)
	//_, err = zipEntry.Write([]byte(context))
	//if err != nil{
	//	fmt.Println("***************************",err)
	//}
	//zipWriter.Close()

	return true
}
