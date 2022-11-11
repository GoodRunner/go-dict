### 背景

作为程序员经常需要阅读一些英文文档，然而由于词汇量有限经常需要去网站上翻译不认识的词汇，为了节约一些时间，所以有了该工具，我们在命令行就可以完成翻译。

### 安装

1. 在有go环境的机器上，执行`go build dict.go` 生成可执行文件`dict`
2. 在`~/.bashrc`或者`~/.zshrc`中追加 `export PATH="$PATH:~/myCmd"`
3. 在`~`目录下创建`myCmd`目录，将生成的可执行文件放进来，`source ~/.zshrc`           

### 演示


![](https://user-images.githubusercontent.com/10717670/201341196-79ad3155-4e6e-4283-8f13-cf4db52a3991.png)
 
