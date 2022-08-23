* 后台前端布局和登录页从`https://github.com/flipped-aurora/gin-vue-admin` 精简而来
* 后台前端用户管理，角色管理，菜单管理，api管理，配置管理重写
* 后端全部重写

## 安装步骤
* 复制`server`下`config.yaml.example`为`config.yaml`并修改配置
* `server`下运行`go mod tidy`
* 复制`web`下`.env.example`为`.env`并修改配置
* `web`下运行`npm install`
* 导入`init.sql`
* 登录后台：`http://127.0.0.1:8081/admin/login` 帐号：`admin` 密码：`123456`