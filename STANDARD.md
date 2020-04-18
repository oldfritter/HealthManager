# 分布式系统组件开发规范

### 一.主要组件
	1.Health  Manager (HM),系统健康监测组件
	2.API，对外的数据接口
	3.Push，数据推送服务
	4.Worker，异步处理系统
	5.Admin，管理员后台
	6.Schedule，定时任务系统

### 二.组件启动时需要立即完成的操作
	1.通过MQ向HM发送心跳消息（包含：组件名称，类型，订阅了哪些Queue，负载等），此后每10秒发送一次心跳消息
	2.通过MQ获取全系统配置（组件清单，Mysql，Redis等）
	3.订阅本类别组件的Queue
	4.创建本组件专属Queue并订阅，本组件关闭时，销毁此Queue

### 三.主要的Exchange
1.组件清单,每30秒广播一次（fanout）

Exchange: aha.components.list

	示例
		[
			{
				"name" : "api-35f5e6fa-804f-11ea-bc55-0242ac130003",
				"kind" : "api",
				"subscribed" : "aha.services.api",
				"details": {
					...
				}
			},
			{
				"name" : "api-6d364d3a-804f-11ea-bc55-0242ac130003",
				"kind" : "api",
				"subscribed" : "aha.services.api",
				"details": {
					...
				}
			},
			{
				"name" : "worker-b259964c-804f-11ea-bc55-0242ac130003",
				"kind" : "worker",
				"subscribed" : "worker.services.worker",
			},
			{
				"name" : "worker-383a9fb8-8050-11ea-bc55-0242ac130003",
				"kind" : "worker",
				"subscribed" : "worker.services.worker",
			},
			...
		]

2.服务配置,每30秒广播一次（fanout）

Exchange: aha.services.list

	示例
    {
      "mysql-main":
      {
        "Name":"main",
        "database":"serbia",
        "username":"root",
        "password":"",
        "host":"127.0.0.1",
        "port":"3306",
        "dbargs":"charset=utf8\u0026parseTime=True\u0026loc=Local",
        "pool":"5"
      },
      "rabbitmq-main":
      {
        "Name":"main",
        "Connect":
        {
          "Host":"127.0.0.1",
          "Port":"5672",
          "Username":"guest",
          "Password":"guest",
          "Vhost":"/main"
        }
      },
      "redis-data":
      {
        "Name":"data",
        "host":"127.0.0.1",
        "port":"6379",
        "db":"6",
        "pool":"5"
      },
      "redis-main":
      {
        "Name":"main",
        "host":"127.0.0.1",
        "port":"6379",
        "db":"0",
        "pool":"5"
      }
    }

3.数据更新，（fanout）

Exchange: aha.data.update

	示例, mysql数据库中configs表中id为123的数据发生改变
	{
		"service": "mysql-main"
		"model" : "configs",
		"id" : "123",
	}

4.发送心跳，（direct）

Exchange: aha.health.heartbeat

	routing key: aha.health.heartbeat
	示例
			{
				"kind" : "api",
				"name" : "api-35f5e6fa-804f-11ea-bc55-0242ac130003",	//  推荐使用UUID生成
				"subscribed" : "aha.services.api",
				"details": {
					... // 其它的一些信息
				}
			},


### 四.主要的Queue

### 五.订阅组件清单消息

	1.创建名为api-d1a773ca-807c-11ea-bc55-0242ac130003的Queue，此Queue的AutoDelete设置为true
	2.绑定Queue：api-d1a773ca-807c-11ea-bc55-0242ac130003与Exchange： aha.components.list
	3.订阅Queue：api-d1a773ca-807c-11ea-bc55-0242ac130003

### 六.订阅服务配置消息

	1.创建名为api-dd1ade3a-807d-11ea-bc55-0242ac130003的Queue，此Queue的AutoDelete设置为true
	2.绑定Queue：api-dd1ade3a-807d-11ea-bc55-0242ac130003与Exchange： aha.services.list
	3.订阅Queue：api-dd1ade3a-807d-11ea-bc55-0242ac130003

### 七.订阅数据更新消息

	1.创建名为api-19946c6e-807e-11ea-bc55-0242ac130003的Queue，此Queue的AutoDelete设置为true
	2.绑定Queue：api-19946c6e-807e-11ea-bc55-0242ac130003与Exchange： aha.data.update
	3.订阅Queue：api-19946c6e-807e-11ea-bc55-0242ac130003
