2016/01/04

[学习总结]

1、上午熟悉OS操作系统，搭建golang开发环境。
   熟悉开发过程中使用到的IDE工具
   
2、下午何李石师兄组织了golang编程基础培训
   编写了一个简单的乘法实例及测试。
   
3、学习使用git管理代码，练习git命令行提交代码

2016/01/05
[学习总结]
今天的主题是类型系统
1、上午学习了类型系统的主要内容：
包括切片，指针，结构体（重点）
常亮的定义、变量定义、字符串、数据类型、数组、map
其中对方法继承和方法覆盖的理解不是很透彻，还需要复习，并且写代码熟悉下。
2、下午做练习题。感觉后面的题还是挺有难度的，通过不断的查资料，对今天讲的大部分内容都有了一定的练习。

2016/01/06
[学习总结]
今天是张乐讲师讲解的控制语句和错误处理相关的内容
今天的内容,信息量感觉有点多,但是讲解的速度非常快,具体总结下今天学习的内容:
1,goto关键字   可以调整执行位置 支持函数内部goto跳转
2,defer关键字  他的执行顺序是"先进后出",并且通过习题可以验证执行顺序是在return之前的
3,对文件的相关操作(复制,修改)还不是很明白
4,下午进行了昨天习题的review,跟何老师请教了了一下昨天没明白的地方,感觉对切片和map的理解又加深了一点点

2016/01/07
[学习总结]
今天是由刘坚军讲师讲解的"go的接口和面向对象是编程"
对"指针"的理杰不是恨到位,在定义struct之后,成员方法一般都是定义为指针类型.下午用了很长一段时间去消化代码,然后跟同事讨论,
对"指针"和"地址"的概念有了些许理解.希望在实战中能够透彻理解.

2016/01/08
[学习总结]
今天的主题是并发编程
1,学习了goroutine的机制
goroutine是被调度的一个最小单位,并且goroutine自身并不知道自己何时被调度.goroutine在占用cpu时,不会让出cpu
2,学习了channel
go里面的消息同步机制，往channel里丢消息和取消息是一个原子操作（消息的通道），
同时处理多个消息是不会发生类似覆盖之类的错误的，类似于消息队列，消息间是相对独立安全的
并且今天对于取地址运算符也有了新的理解(&)

2016/01/08
[学习总结]
今天做的主要是针对上周学习后,课后习题的review,和修改原先的学习计划,对之后的学习内容进行调整.

2016/01/12
[学习总结]
RESTful API  部分理解:
资源＋动作＋参数 ＝ 一个请求
REST的核心原则是将API拆分为逻辑上的资源。这些资源通过http被操作（GET,POST,PUT,DELETE）。
从API用户的角度来看，“资源”应该是个名词。API设计的时候，不需要把它们一对一的都暴露出来，这里的关键是，隐藏内部资源，暴露必须的外部资源。
一旦定义好了要暴露的资源，就可以定义资源上允许的操作，以及这些操作和API的对应关系：
GET /tickets # 获取ticket列表
GET /tickets/12 # 查看某个具体的ticket
POST /tickets # 新建一个ticket
PUT /tickets/12 # 更新ticket 12.
DELETE /tickets/12 #删除ticekt 12
可以看出使用REST的好处在于可以充分利用http的强大，实现对资源的CURD功能。
［CURD： 创建create、更新update、读取read、删除delete］
而这里只需要一个endpoint：  ／tickets
关于endpoint的单数复数：
一个可以遵从的规则是：虽然看起来使用复数来描述某一个资源实例看起来别扭，但是统一所有的endpoint，是用复数使得URL更加规整。这让API的使用者更加容易理解，对于开发者来说，也更容易实现。
关于如何处理关联：
假如关联和资源独立，那么在资源的输出表示中保存相应资源的endpoint，然后API的使用者就可以通过点击链接找到相关的资源。如果关联和资源联系紧密，资源的输出表示就应该直接保存相应资源信息。
GET /tickets/12/messages- Retrieves list of messages for ticket #12
GET /tickets/12/messages/5- Retrieves message #5 for ticket #12
POST /tickets/12/messages- Creates a new message in ticket #12
PUT /tickets/12/messages/5- Updates message #5 for ticket #12
PATCH /tickets/12/messages/5- Partially updates message #5 for ticket #12
DELETE /tickets/12/messages/5- Deletes message #5 for ticket #12
例如，这里如果message资源是独立存在的，那么上面GET／tickets/12/messages就返回相应的message的链接；相反的如果message不独立存在，他和ticket依附存在，则上面的API调用返回直接返回message信息


