package myfile

import "os"

func start() {

	file, err := os.OpenFile("Hello.log", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	count := 0
	for count < 100000 {
		file.WriteString("Reader实现了给一个io.Reader接口对象附加缓冲。\n")
		count++
	}
}
