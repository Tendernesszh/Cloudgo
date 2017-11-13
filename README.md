# Cloudgo

我的项目实现的web功能：<br>
1.在访问“/”时显示“Hello world！”<br>
2.访问“/user/:name”时显示“Hello name！”其中name是访问时用户自定义的。<br>

对于web框架，我选择了beego，我认为begoo有很多优点<br>
比如：<br>
1.beego比较全面，可以实现的功能很多<br>
2.beego的教程很丰富，不仅有文字教程，还有视频教程发布在网上，让人学习起来感到轻松一些<br>
3.bee工具包使得项目的框架创建变得十分快捷<br>

#### 安装beego，bee的方法以及项目的快速创建讲解：
先在#GOPATH/src路径下于命令行中输入go get github.com/astaxie/beego从而安装beego<br>
再在#GOPATH/src路径下于命令行中输入go get github.com/beego/bee从而安装开发工具bee<br>
之后在#GOPATH/src路径下于命令行中输入bee new webgo就可创建有完整框架的名为webgo的项目<br>
#### 使用curl来测试 curl -v http://loacalhost:9090/  和  curl -v http://loacalhost:9090/user/zhang
结果如下：<br>
![image](https://github.com/Tendernesszh/Cloudgo/blob/master/webgo/curl%E6%B5%8B%E8%AF%95.png)
#### ab测试（指令为：ab -n 1000 -c 100 http://localhost:9090/user/zhang）
其中-n代表执行的请求数量，-c代表并发请求个数。<br>

结果如下：<br>
![image](https://github.com/Tendernesszh/Cloudgo/blob/master/webgo/%E5%8E%8B%E5%8A%9B%E6%B5%8B%E8%AF%95.png)

可以看到在56ms内所有的请求被处理完成，且没有失败请求。
