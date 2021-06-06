
## 环境

```
go >= 1.16
nodejs >= v14.2.0
npm >= 6.14.4
yarn >=1.22.0
npx >= 6.14.4
```

## 创建 webui

```
## 修改配置，加速下载速度
npm config get registry
npm config set registry https://registry.npm.taobao.org/
## 还原
npm config set registry https://registry.npmjs.org/

yarn config get registry
yarn config set registry https://registry.npm.taobao.org
## 还原
yarn config set registry https://registry.yarnpkg.com

npx create-umi ./webui
'''
选择 ant-design-pro
选择 Pro V4
选择 TypeScript
选择 complete
'''

cd webui && yarn install

yarn start

## 浏览器打开 http://localhost:8000

// 打开开发模式下页面右下角的小气泡，方便添加区块和模版等 pro 资产

yarn add @umijs/preset-ui -D
yarn start

## 构建 release 版本
yarn build
```

**遇到问题**:

如果有遇到使用问题，可以尝试:

```
rm -fr src/.umi
BABEL_CACHE=none yarn start
```

## 创建运行 GO 服务

```
## 设置 golang proxy 代理
go env -w GOPROXY=https://goproxy.cn,direct

go mod init go-ant-design-pro

## 运行，注意：运行 go 程序之前要先编译 webui
go run main.go

## 打开 http://localhost:8080 默认就会展示 webui（ant-design-pro）页面，注意不是上面的 8000 端口。
```
