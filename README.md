<p align="center">
    <h3 align="center">gbk2utf8</h3>
    <p align="center">提供 GB2312/GBK/GB18030 与 utf-8 编码之间的互相转换</p>
    <p align="center">
<a href="https://github.com/fluffos/gbk2utf8/releases/latest">
<img alt="最新版本" src="https://img.shields.io/github/v/release/fluffos/gbk2utf8.svg?logo=github&style=flat-square">
</a>
<a href="https://github.com/fluffos/gbk2utf8/actions?workflow=Release">
<img alt="Release workflow" src="https://github.com/fluffos/gbk2utf8/workflows/Release/badge.svg">
</a>
<a href="https://github.com/fluffos/gbk2utf8/actions?workflow=Build">
<img alt="Build workflow" src="https://github.com/fluffos/gbk2utf8/workflows/Build/badge.svg">
</a>
<a href="https://goreportcard.com/report/github.com/fluffos/gbk2utf8">
<img alt="Go Report" src="https://goreportcard.com/badge/github.com/fluffos/gbk2utf8">
</a>
<a href="https://golangci.com/r/github.com/fluffos/gbk2utf8">
<img alt="GolangCI Report" src="https://github.com/golangci/golangci-web/blob/master/src/assets/images/badge_a_plus_flat.svg">
</a>
    </p>
</p>

## 下载安装

### 预编译二进制

前往[发布页面](https://github.com/fluffos/gbk2utf8/releases)下载适合你的操作系统的版本。

### HomeBrew 安装

[HomeBrew](https://brew.sh) 用户可以使用 `brew` 命令快速安装：

```
brew tap mudclient/tap
brew install gbk2utf8
```

### 源码安装

```Shell
go get -u github.com/fluffos/gbk2utf8
```

## 用法举例

gbk2utf8 是一个命令行工具，通过命令行参数来控制其行为。举个例子：

### 转换单个文件

```
gbk2utf8 --from GB18030 --to UTF8 --src foo.c --dst foo-utf8.c
```

以上命令可以将文件 `foo.c` 从 `GB18030` 编码转换至 `UTF8` 编码，并保存为文件 `foo-utf8.c`。

老实说，上面命令所做的事情和 [iconv](https://www.gnu.org/savannah-checkouts/gnu/libiconv/documentation/libiconv-1.15/iconv.1.html) 没什么不同。甚至其支持的编码种类还不如 `iconv` 多。

但 `gbk2utf8` 的主要优点在于它能够一次性[转换整个目录树](#转换整个目录树)。

### 转换整个目录树

```
gbk2utf8 --from GB18030 --to UTF8 --src code --dst code-utf8 --pattern "*.c"
```

以上命令可以将 `code` 目录及其子目录下的所有 `*.c` 文件，
从 `GB18030` 编码转换至 `UTF8` 编码，并存放在 `code-utf8` 目录下。

另外，`gbk2utf8` 还有一些实用的命令行参数。你可以查阅[参数列表](#参数列表)。

## 参数列表

在命令行下输入 `gbk2utf8 --help` 可以看到帮助如下：

```Shell
$ gbk2utf8 --help
gbk2utf8(version v1.0.3)

Usage:
  gbk2utf8 [flags]

Flags:
  -c, --config FILENAME   config FILENAME, default to `config.yaml` or `config.json`
      --version           just print version number only
  -h, --help              show this message
      --gen-yaml          generate config.yaml
      --gen-json          generate config.json
  -f, --from string       将要转换的文件的编码，可选值: GB2312/GBK/GB18030/BIG5/UTF8 (default "GB18030")
  -t, --to string         想要转换成的目的文件编码，可选值: GB2312/GBK/GB18030/BIG5/UTF8 (default "UTF8")
  -s, --src string        想要转换的文件或者目录。如果是目录，则将会转换此目录下所有的文件，包含子目录
  -d, --dst string        转换完的目标文件名或目录名。注意目标路径下的内容会被清空
  -p, --pattern string    想要转换的文件名的模式。只有文件名符合模式的文件才会被转换 (default "*.c")
```

## 如何贡献

* 体验并向周围的人分享你的体验结果。
* 通过[提交 issue](https://github.com/fluffos/gbk2utf8/issues/new) 来反馈意见。
* 通过 PR 来贡献代码，贡献代码时请先阅读[贡献指南](CONTRIBUTING.md)。
