#advShow
=======

##advShow 简介
> advShow(advise show)是广告显示服务器，客户端每隔一段时间向服务器发起请求。
服务器按照客户端的请求给客户端回复相应的广告。这这个工具的目的是为了让公司的动态
进度显示器能够很容易的维护。advShow服务器端代码是由golang完成。

##使用说明
编译后运行程序，现在只支持在main目录下面运行。后面会修改相关目录后可以在bin运行。
服务器运行以后用户只要修改`config`和`static`这两个文件夹就可以。
###config
-------server.xml                                **服务器配置**
```xml
	<server>
	    <monitorsrv>
	        <localserver>127.0.0.1</localserver> **服务器地址**
	        <port>8080</port>                    **服务器监听端口**
	    </monitorsrv>
	    <displayctl>
	        <delaytime>20</delaytime>            **广告更新时间单位秒**
	        <transtime>2</transtime>     
	        <trmaxno>3</trmaxno>
	        <recmaxno>3</recmaxno>
	    </displayctl>
	</server>
```

-------content.xml 								**广告内容配置**
```xml
<content>
	<dissplay>
		<name>pic</name>                       **广告名称便于记忆取的别名随便取**
		<type>static</type>                    **是本主机的静态网页还是外部网页。static表示是今天网页，link表示是外部网页。**
		<file>1.htm</file>                     **静态网页名称**
	</dissplay>
	<dissplay>
		<name>pic</name>
		<type>static</type>
		<file>2.htm</file>
	</dissplay>
	<dissplay>
		<name>baidu</name>
		<type>link</type>
		<file>http://www.baidu.com/</file>          **记得一定要写全，写网整的地址连接，如果不是完整的连接会被浏览器误认为是本地今天链接。其实也可以修改代码支持。**
	</dissplay>
</content>
```
广告显示的顺序按照`content.xml`文件里面的定义顺序
###static
放所有静态网页和网页相关资源的地方

###连接服务器显示广告
1.在客户端浏览器地址输入栏输入http://hostip:pot/show
2.允许弹出网页
3.全屏弹出网页



