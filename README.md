# DnsPodAPI-go

##*******账户相关*******
Describe|API|parameter|请求方式
--|:--:|:--:|--:
获取用户信息|/getuserinfo|..|POST
获取用户日志|/getuserlog|..|POST

##*******域名相关*******
Describe|API|parameter|请求方式
--|:--:|:--:|--:
创建域名|/createdomain|..|POST
域名列表|/listdomain|..|POST
删除域名|/removedomain|..|POST
设置域名状态|/setdomainstatus|..|POST
获取域名信息|/getdomaininfo|..|POST
获取域名日志|/getdomainlog|..|POST

##*******记录相关*******
Describe|API|parameter|请求方式
--|:--:|:--:|--:
记录|/createrecord|..|POST
记录列表|/listrecord|..|POST
修改记录|/setuprecord|..|POST
删除记录|/removerecord|..|POST
更新DNS记录|/ddnsrecord|login_token=ID,TOKEN<br>form=json<br>domain_id=DOMAIN_ID<br>record_id=RECORD_ID<br>sub_domain=WWW<br>record_line=默认<br>value=172.0.0.1|POST
获取记录日志|/getrecordinfo|..|POST
