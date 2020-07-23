package test

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func TestJsonWrite001(t *testing.T) {
	info := []Website{
		{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}},
		{"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}},
	}
	// 创建文件
	filePtr, err := os.Create("sources//info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func TestJsonRead001(t *testing.T) {
	filePtr, err := os.Open("sources//info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	var info []Website
	// 创建Json编码器
	encoder := json.NewDecoder(filePtr)
	err = encoder.Decode(&info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
		fmt.Println("info：", info)
	}
}

func TestXmlWrite001(t *testing.T) {
	info := Website{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}
	// 创建文件
	filePtr, err := os.Create("sources//info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := xml.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func TestXmlRead001(t *testing.T) {
	filePtr, err := os.Open("sources//info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	var info Website
	// 创建Json编码器
	encoder := xml.NewDecoder(filePtr)
	err = encoder.Decode(&info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
		fmt.Println("info：", info)
	}
}

func TestGobWrite001(t *testing.T) {
	info := map[string]string{
		"name": "Joy",
		"sex":  "nan",
	}

	name := "demo.gob"
	File, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("OpenFile ERR:", err)
		return
	}
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println("Encode ERR:", err)
	}
}

func TestGobRead001(t *testing.T) {
	var m map[string]string
	File, err := os.Open("demo.gob")
	if err != nil {
		fmt.Println("OpenFile ERR:", err)
		return
	}
	d := gob.NewDecoder(File)
	d.Decode(&m)
	fmt.Println("result:", m)
}

func TestTextWrite001(t *testing.T) {
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open File ERR:", err)
		return
	}
	defer file.Close()
	str := "this is a test data\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func TestTextRead001(t *testing.T) {
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
}

func TestCusBinWrite001(t *testing.T) {
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println("Open File ERR:", err)
		return
	}
	defer file.Close()
	type Website struct {
		Url int32
	}
	for i := 1; i <= 10; i++ {
		info := Website{int32(i)}

		var binBuf bytes.Buffer
		binary.Write(&binBuf, binary.LittleEndian, info)
		b := binBuf.Bytes()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}

func TestCusBinRead001(t *testing.T) {
	file, err := os.Open("output.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	type Website struct {
		Url int32
	}
	m := Website{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}

func TestZipWrite001(t *testing.T) {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)
	// 创建一个压缩文档
	w := zip.NewWriter(buf)
	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}
	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}
	// 将压缩文档内容写入文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)
}

func TestZipRead001(t *testing.T) {
	// 打开一个zip格式文件
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer r.Close()
	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		fmt.Printf("文件名: %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Printf(err.Error())
		}
		_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
		if err != nil {
			fmt.Printf(err.Error())
		}
		rc.Close()
	}
}

func TestTarWrite001(t *testing.T) {
	f, err := os.Create("./output.tar") //创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()
	fileinfo, err := os.Stat("./output.txt") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}
	err = tw.WriteHeader(hdr) //写入头文件信息
	if err != nil {
		fmt.Println(err)
	}
	f1, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将main.exe文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

func TestTarRead001(t *testing.T) {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("123" + fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}
