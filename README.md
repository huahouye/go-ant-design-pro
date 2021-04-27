
## environment

```
go >= 1.16
nodejs >= v14.2.0
npm >= 6.14.4
yarn >=1.22.0
npx >= 6.14.4
```

## create webui

```
## speed up js package download for people in mainland China
npm config get registry
npm config set registry https://registry.npm.taobao.org/
## rollback
npm config set registry https://registry.npmjs.org/

yarn config get registry
yarn config set registry https://registry.npm.taobao.org
## rollback
yarn config set registry https://registry.yarnpkg.com

npx create-umi ./webui
'''
chose ant-design-pro
chose Pro V4
chose TypeScript
chose complete
'''

cd webui && yarn install

yarn start

## open http://localhost:8000

## if you want to enable and develop webui under debug mode for convenient, install @umijs/preset-ui and restart webui again, there will be a bubble at the bottom right

yarn add @umijs/preset-ui -D
yarn start

## build for release
yarn build
```

## create and run GO server

```
go mod init go-ant-design-pro

## run
go run main.go

## open http://localhost:8080 it will open webui by default
```
