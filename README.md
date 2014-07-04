#advShow
=======

##advShow 简介
> advShow(advise show)是广告显示服务器，客户端每隔一段时间向服务器发起请求。
服务器按照客户端的请求给客户端回复相应的广告。这这个工具的目的是为了让公司的动态
进度显示器能够很容易的维护。advShow服务器端代码是由golang完成。

##使用说明
编译后运行程序，现在只支持在main目录下面运行。后面会修改相关目录后可以在bin运行。
服务器运行以后用户只要修改`config`和`static`这两个文件夹就可以。
config
-------server.xml
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
