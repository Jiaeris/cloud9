package protodemo

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

//官方指南：https://developers.google.cn/protocol-buffers/docs/overview
//github：https://github.com/google/protobuf
//
//准备好protoc,在go中使用准备好protoc-gen-go否则报错
//编译脚本：java,go,cpp
//protoc --java_out=./ *.proto
//protoc --plugin=protoc-gen-go=E:\Go_WorkSpace\bin\protoc-gen-go.exe --go_out=./ *.proto
//protoc --cpp_out=./ *.proto

func run() {
	p1 := &Person{}
	p1.Name = "Yunga"
	p1.IdNum = "510121"
	p1.Sex = Sex_MAN
	p1.Bio = "ok ok ok "

	//编码
	buf, err := encode(p1)
	if err != nil {
		fmt.Println(err)
	}
	//解码
	p2 := &Person{}
	err = decode(buf, p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("p1:", p1)
	fmt.Println("p2.sex:", p2.Sex)

	//proto json
	jbuf, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json:", string(jbuf))
	p3 := &Person{}
	err = jsonpb.UnmarshalString(string(jbuf), p3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("p1:", p1)
	fmt.Println("p3.sex:", p3.Sex)
}

func encode(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

func decode(buf []byte, pb proto.Message) error {
	return proto.Unmarshal(buf, pb)
}
