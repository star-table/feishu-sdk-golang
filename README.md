## Introduce
FeiShu ([https://feishu.cn](https://feishu.cn)) Server Api Sdk For Go.

## Doc
FeiShu Server Api Doc ([https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN](https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN))

## Feature
- [x] 授权
  - [x] 获取 app_access_token（企业自建应用）
  - [x] 获取 app_access_token（应用商店应用）
  - [x] 获取 tenant_access_token（企业自建应用）
  - [x] 获取 tenant_access_token（应用商店应用）
  - [x] 重新推送 app_ticket
  - [x] [查询租户授权状态](https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#dCNL6V)
  - [x] [申请授权](https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#kHHiAa)
- [ ] 身份验证
  - [ ] 请求身份验证
  - [x] 获取登录用户身份
  - [x] 刷新access_token
  - [x] 获取用户信息
- [ ] 通讯录
  - [x] 获取通讯录授权范围
  - [x] 获取子部门列表
  - [ ] 获取子部门 ID 列表
  - [x] 获取部门详情
  - [ ] 批量获取部门详情
  - [x] 获取部门用户列表
  - [x] 获取部门用户详情
  - [ ] 获取企业自定义用户属性配置
  - [x] 批量获取用户信息
  - [ ] 新增用户
  - [ ] 更新用户信息
  - [ ] 删除用户
  - [ ] 新增部门
  - [ ] 更新部门信息
  - [ ] 删除部门
  - [ ] 批量新增用户
  - [ ] 批量新增部门
  - [ ] 查询批量任务执行状态
  - [ ] 获取应用管理员管理范围
  - [ ] 获取角色列表
  - [ ] 获取角色成员列表
- [ ] 用户信息
  - [ ] 使用手机号或邮箱获取用户 ID
- [ ] 应用信息
  - [ ] 获取应用管理权限
  - [ ] 获取应用在企业内的可用范围
  - [ ] 获取用户可用的应用
  - [ ] 获取企业安装的应用
  - [ ] 更新应用可用范围
  - [ ] 设置用户可用应用
- [ ] 应用商店计费信息
  - [ ] 查询用户是否在应用开通范围
  - [ ] 查询租户购买的付费方案
  - [ ] 查询订单详情
- [x] 群组
  - [x] 获取用户所在的群列表
  - [x] 获取群成员列表
  - [x] 搜索用户所在的群列表
- [ ] 机器人
  - [x] 批量发送消息
  - [x] 发送文本消息
  - [x] 发送图片消息
  - [x] 发送富文本消息
  - [ ] 发送群名片
  - [ ] 消息撤回
  - [x] 发送卡片通知
  - [x] 群信息和群管理
    - [x] 创建群
    - [x] 获取群列表
    - [x] 获取群信息
    - [x] 更新群信息
    - [x] 拉用户进群
    - [x] 移除用户出群
    - [x] 解散群
  - [x] 机器人信息与管理
    - [x] 拉机器人进群
- [x] 日历
  - [x] 获取日历
  - [x] 获取日历列表
  - [x] 创建日历
  - [x] 删除日历
  - [x] 更新日历
  - [x] 获取日程
  - [x] 创建日程
  - [x] 获取日程列表
  - [x] 删除日程
  - [x] 更新日程
  - [x] 邀请/移除日程参与者

## 参考
* 飞书官方文档 https://open.feishu.cn/document/uQjL04CN/ucDOz4yN4MjL3gzM
