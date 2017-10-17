# Week1实验报告

##实验要求：

用go语言开发linux命令行实验工具。

借助网站如下：https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html

---

## 总体思想：

根据网址的c语言实现，用go语言翻译即可。

> 使用的额外go包：
>
> - os包：用于处理命令行参数和使用Stdin，Stdout，Stderr
> - strconv包：字符串转int值
> - os/exec包：用于新建子进程，并执行shell指令
> - bufio包：用于缓冲区输入和缓冲区输出
> - io包：用于文件处理

（PS：由于不会单独处理input_file参数，未使用flag包。）

---

##基本功能实现：

由于是照着c语言翻译，因此c语言的功能全部实现。即：

- 命令行参数处理
- 两种输入格式
- 异常处理
- 文件代替Stdin输入
- Shell重定向
- 调用lp指令打印文件

可使用input.txt和out.txt测试。

---

## 实验感想：

感觉golang和c还是有很多共同之处的，有些不习惯的有三点：

1. :=可以定义变量
2. 返回值不仅有一个
3. 类型写在后面

还是要多多学习啊。。。