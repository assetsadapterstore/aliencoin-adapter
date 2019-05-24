# aliencoin-adapter

aliencoin-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建ALC.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url
serverAPI = "http://127.0.0.1:8080"
# RPC Authentication Username
rpcUser = "alcuser"
# RPC Authentication Password
rpcPassword = "alcuser"
# Is network test?
isTestNet = false
# support segWit
supportSegWit = false
# minimum transaction fees
minFees = "0.00001"

```

## 资料介绍

### 区块浏览器

http://39.105.54.199:8080/

### github

https://github.com/waixingdun/aliencoin

### 适配资料

#### 地址编码的相关代码或代码链接

aliencoin/src/chainparams.cpp的CMainParams()

```C++
base58Prefixes[PUBKEY_ADDRESS] = {23};
base58Prefixes[SCRIPT_ADDRESS] = {48};
base58Prefixes[SECRET_KEY] =     {46};
base58Prefixes[EXT_PUBLIC_KEY] = {0x04, 0x86, 0xE6, 0x36};
base58Prefixes[EXT_SECRET_KEY] = {0x04, 0x89, 0xE9, 0x49};

```


#### rpc api文档的链接

https://en.bitcoin.it/wiki/API_reference_(JSON-RPC)
