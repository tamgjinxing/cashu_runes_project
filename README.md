用户创建多少钱包，后端就有多少个钱包文件
当用户创建钱包的时候，传入 pubkey，后端创建的时候，直接以 pubkey 命名钱包文件
创建钱包：
./ord wallet --name <传入的钱包名称> create

eg:
./ord wallet --name publicKey001 create
./ord wallet --name publicKey002 create
...

提供以下接口给前端调用 1.查看钱包是否创建 2.创建钱包 3.查询余额 4.获取接收地址 5.查询交易记录 6.恢复钱包 7.导出钱包描述符 8.发送符文 1.查看是否创建钱包
/api/checkWallet

入参：
pubkey

出参：
isExist true / false

2. 创建钱包
   /api/createdWallet
   入参：
   pubkey
   passphrase optional

出参：
mnemonic

3. 查询余额
   /api/getBalance

入参：
pubkey

出参：(按照 ord wallet balance 返回的结果返回)
{
"cardinal": 0,
"ordinal": 0,
"runes": {},
"runic": 0,
"total": 0
}

4. 查看接收地址
   /api/getAddress

入参：
pubkey

出参：
address

5. 查询交易记录
   /api/getTransactions

入参：
pubkey

出参：
List<TransactionDTO>  
 transactionDTO:
transctionId
status confirmed not-confirmed

6. 恢复钱包
   /api/restoreWallet

入参：
pubkey
mnemonic
restoreFrom: Restore wallet from <SOURCE> on stdin. [possible values: descriptor, mnemonic]
passphrase

出参：

7. 导出描述符
   /api/exportDescriptors

入参：
pubkey

出参：
可把这些内容保存到一个 json 文件并返回给到前端.
==========================================
= THIS STRING CONTAINS YOUR PRIVATE KEYS =
= DO NOT SHARE WITH ANYONE =
==========================================
{
"wallet*name": "ord",
"descriptors": [
{
"desc": "tr([551ac972/86'/1'/0']tprv8h4xBhrfZwX9o1XtUMmz92yNiGRYjF9B1vkvQ858aN1UQcACZNqN9nFzj3vrYPa4jdPMfw4ooMuNBfR4gcYm7LmhKZNTaF4etbN29Tj7UcH/0/*)#uxn94yt5",
"timestamp": 1296688602,
"active": true,
"internal": false,
"range": [
0,
999
],
"next": 0
},
{
"desc": "tr([551ac972/86'/1'/0']tprv8h4xBhrfZwX9o1XtUMmz92yNiGRYjF9B1vkvQ858aN1UQcACZNqN9nFzj3vrYPa4jdPMfw4ooMuNBfR4gcYm7LmhKZNTaF4etbN29Tj7UcH/1/\_)#djkyg3mv",
"timestamp": 1296688602,
"active": true,
"internal": true,
"range": [
0,
999
],
"next": 0
}
]
}

8. 发送 runes
   /api/sendRunes

入参：
pubkey
targetAddress / targetPubkey
feeRate
runesName
runesAmount

出参：
{
"txid": "83be38338385c13782733a30332194c9b723dae75d6a499ea7e34d7d00b79b60",
"psbt": "cHNidP8BALICAAAAAgeF+iLYxf10lGsYhHi5oVAPj3hFW4wfLXaUvuQOF73qAQAAAAD/////JX7Tll567myrQ/uoOzU64Ey/zLD0l5skf8yb7LjIO0cAAAAAAP3///8CECcAAAAAAAAiUSCylgQqw/6wIlb0Q2GOTUKYWIpETY8gZSSJVJiIfBCWZ/aPAAAAAAAAIlEgYRVRR+2+D0GvgWvjCEiFJASXEs4sDyqb9WzZJ80va+AAAAAAAAEBKyICAAAAAAAAIlEgqaTT7XNKCb7YYtMy085Ni5ZQDh/dCQ+PyMvBb7ny1ygBCEIBQK/9R8m1eDcEvAiie2IUSGHeWV1zWSsCMezjEebMP9PwVp2+Dy+jwpZrjQK/b08CHessTZpZL+JSPxGlYbb/rvsAAQErsLoAAAAAAAAiUSB8XgLvEN5LoI/TVr9WXn7Spo83jZzwLVN30ypM86FpjgEIQgFAPAWCjfmbIb9sPwrm9aGx8BO39kGC+mI4ymU1BBxCC9LOCW3c/IjuOKSBcLvSdg9ACnygAsQzWR9pvftObZMIUgAAAQUgy4Iq2LGbXdaL1SWX4hxk3366TC0sXnBL51GNATd+QzshB8uCKtixm13Wi9Ull+IcZN9+ukwtLF5wS+dRjQE3fkM7GQDzvaZxVgAAgAAAAIAAAACAAQAAAAQAAAAA",
"outgoing": "6:AMAZING•PYRAMID•NUMBER",
"fee": 1484
}
