# go-binds
main.go为demo，
ret, err := service.GetBlackList()
if err != nil {
fmt.Println(err)
}
返回值说明：
ret为一个map，key是地址 ，value是一个结构体

结构体定义为
type GlobalAccessV1StorageUserStatus struct {
Status    uint8  
PauseFrom *big.Int
PausedTo  *big.Int
}
Status 是enum变量

enum Status{
SEND_AND_RECEIVE,
RECEIVE_ONLY,
SEND_ONLY,
NO_SEND_AND_RECEIVE
}

PauseFrom和PauseTo是unix时间，目前测试合约中写入的是2022-12-20 2024-12-20