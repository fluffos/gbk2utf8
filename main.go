package main

//go:generate go run github.com/flw-cn/go-smartConfig/gen-ver -a gbk2utf8

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/flw-cn/go-smartConfig"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// Config 是程序的配置数据结构。
// 通过 github.com/flw-cn/go-smartConfig 的支持，可以同时将配置文件的内容和
// 命令行参数的内容捕获至本数据结构。
//
type Config struct {
	From    string `flag:"f|GB18030|将要转换的文件的编码，可选值: GB2312/GBK/GB18030/BIG5/UTF8"`
	To      string `flag:"t|UTF8|想要转换成的目的文件编码，可选值: GB2312/GBK/GB18030/BIG5/UTF8"`
	Src     string `flag:"s||想要转换的文件或者目录。如果是目录，则将会转换此目录下所有的文件，包含子目录"`
	Dst     string `flag:"d||转换完的目标文件名或目录名。注意目标路径下的内容会被清空"`
	Pattern string `flag:"p|*.c|想要转换的文件名的模式。只有文件名符合模式的文件才会被转换"`
}

func main() {
	app := NewApp(AppName, Version)
	app.LoadConfig()

	begin := time.Now()

	app.Run()

	log.Printf("全部转换完成。共转换了 %d 个文件，耗时 %s。", app.files, time.Since(begin))
}

// App 用来封装应用程序，主要是为了在几个方法之间共享变量及状态。
//
type App struct {
	config Config

	name    string
	version string

	from *encoding.Decoder
	to   *encoding.Encoder

	bytes int64
	files int
}

// NewApp 用来创建一个 App，创建时必须指定 App 名称及版本号。
//
func NewApp(name string, version string) *App {
	return &App{
		name:    name,
		version: version,
	}
}

// LoadConfig 方法用来从配置文件和命令行参数中捕获程序运行时配置。
//
func (app *App) LoadConfig() {
	smartConfig.VersionDetail = func() string {
		info := fmt.Sprintf("版本信息如下：\n程序版本: %s\n编译时间: %s\n编译环境: %s\n编译设备: %s\n其中：\n",
			Version, BuildTime, BuildGoVersion, BuildHost)
		for _, contributor := range Contributors {
			info += fmt.Sprintf("    %s 贡献了 %d 行代码\n", contributor.Name, contributor.Lines)
		}

		return info
	}()

	smartConfig.LoadConfig(app.name, app.version, &app.config)

	app.from = resolveEncoding(app.config.From).NewDecoder()
	app.to = resolveEncoding(app.config.To).NewEncoder()
}

// Run 方法遍历目录，为每个文件执行编码转换操作。
//
func (app *App) Run() {
	err := filepath.Walk(app.config.Src, app.walk)
	if err != nil {
		return
	}
}

func (app *App) walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	baseName := filepath.Base(path)

	matched, err := filepath.Match(app.config.Pattern, baseName)
	if err != nil {
		return err
	}

	if !matched {
		return nil
	}

	output, err := filepath.Rel(app.config.Src, path)
	if err != nil {
		return err
	}

	output = filepath.Join(app.config.Dst, output)
	outDir := filepath.Dir(output)
	_ = os.MkdirAll(outDir, 0755)

	return app.transform(path, output)
}

func (app *App) transform(input string, output string) error {
	in, err := os.Open(input)
	if err != nil {
		return err
	}

	out, err := os.Create(output)
	if err != nil {
		return err
	}

	dst := transform.NewWriter(out, app.to)
	src := transform.NewReader(in, app.from)

	n, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("文件转换 %s => %s 失败: %s", input, output, err)
	} else {
		log.Printf("文件转换 %s => %s 成功。", input, output)
	}

	app.bytes += n
	app.files++

	return nil
}

func resolveEncoding(e string) encoding.Encoding {
	e = strings.ToUpper(e)
	switch e {
	case "GB2312", "HZ-GB-2312", "HZGB2312", "EUC-CN", "EUCCN":
		return simplifiedchinese.HZGB2312
	case "GBK", "CP936":
		return simplifiedchinese.GBK
	case "GB18030":
		return simplifiedchinese.GB18030
	case "BIG5", "BIG-5", "BIG-FIVE":
		return traditionalchinese.Big5
	case "UTF8", "UTF-8":
		return encoding.Nop
	}

	log.Fatalf("invalid encoding name: %s", e)

	return encoding.Nop
}
