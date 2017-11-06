package main

//文件上传服务，导入fmt打印输出，io流操作。os磁盘操作，template模板操作，http请求操作,log日志输出
import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

//文件上传监听函数

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //请求上传路径
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil) //输出页面
	} else {
		//将文件写入到缓存中，指定内存大小
		r.ParseMultipartForm(32 << 20)
		//读取文件
		file, handler, err := r.FormFile("filename")
		if err != nil {
			fmt.Println(err)
			return
		}
		//打开文件缓存
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		//不要忘记关闭
		defer f.Close()
		//拷贝内容
		io.Copy(f, file)

		fmt.Fprintf(w, "%v", handler.Header)
		fmt.Fprintf(w, "上传成功！")
	}
}

func main() {
	// 注册监听
	http.HandleFunc("/upload", upload)
	//启动服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("启动失败", err)
	}
}
