# agenda 开发项目

## 一、需求分析

- 业务需求：具体见命令


- 功能需求： 设计一组命令完成 agenda 的管理，例如：

  - agenda help ：列出命令说明
  - agenda register -uUserName –password pass –email=a@xxx.com ：注册用户
  - agenda help register ：列出 register 命令的描述
  - agenda cm … : 创建一个会议
  - 原则上一个命令对应一个业务功能

  ​

- 持久化要求：

  - 使用 json 存储 User 和 Meeting 实体
  - 当前用户信息存储在 curUser.txt 中



## 二、命令与参数

**用户注册**

1. 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
2. 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

```bash
agenda register -uUserName –password pass -email=a@xxx.com -phone=phoneNum
```

**用户登录**

1. 用户使用用户名和密码登录 Agenda 系统。
2. 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

```bash
agenda login -uUserName –password pass
```

**用户登出**

1. 已登录的用户登出系统后，只能使用用户注册和用户登录功能。

```bash
agenda logout
```

**用户查询**

1. 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。

```bash
agenda -lUser
```

**用户删除**

1. 已登录的用户可以删除本用户账户（即销号）。
2. 操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。
3. 删除成功则退出系统登录状态。删除后，该用户账户不再存在。
4. 用户账户删除以后：
   - 以该用户为 **发起者** 的会议将被删除
   - 以该用户为 **参与者** 的会议将从 **参与者** 列表中移除该用户。若因此造成会议 **参与者** 人数为0，则会议也将被删除。

```bash
agenda delete
```

**创建会议**

1. 已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册 用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
   - 会议主题(title)（在会议列表中具有唯一性）
   - 会议参与者(participator)
   - 会议起始时间(start time)
   - 会议结束时间(end time)
2. 注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
3. 用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。

```bash
agenda mkmeeting –title Name –participator user1 [user2 ….] –stime start –etime end
```

**增删会议参与者**

1. 已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
2. 增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
3. 删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。

```bash
agenda meetingadd –participator user1 [user2 …]
agenda meetingdel –participator user1 [user2 …]
```

**查询会议**

1. 已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
2. 用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
3. 在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
4. 注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。

```bash
agenda querymeeting -stime start -etime end
```

**取消会议**

1. 已登录的用户可以取消 自己发起 的某一会议安排。
2. 取消会议时，需提供唯一标识：会议主题（title）。

```bash
agenda meetingdel –title Name
```

**退出会议**

1. 已登录的用户可以退出 自己参与 的某一会议安排。
2. 退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。

```bash
agenda meetingout –title Name
```

**清空会议**

1. 已登录的用户可以清空 自己发起 的所有会议安排。

```bash
agenda meetingclear
```

