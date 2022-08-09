# notifier 
该工具会逐行读取标准输入，然后将每一行发送给企业微信群。比如：可以将自动化工具扫描到的漏洞发送给该工具，该工具将会自动将漏洞信息推送给企业微信群。

## Download
```
go install github.com/ShangRui-hash/notifier@latest
```

## Usage

step1. 创建企业微信，并在手机版企业微信群中添加群机器人，获取web_hook地址

step2. 将web_hook地址写入到环境变量中,例如:
```
vim $HOME/.zsh
```
```
export WX_WEB_HOOK=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=14ece0d0-4b96-49a1-99a3-5b41e4xxxxxx
```
```
source $HOME/.zsh 
```
step3. 使用样例 
```
echo "happy hacking" | notifier
```

```
cat vul.txt | notifier 
```