package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 消息
type Message struct {
	gorm.Model
	UserId     int64  //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn          *websocket.Conn
	Addr          string
	HeartbeatTime uint64
	LoginTime     uint64
	DataQueue     chan []byte
	GroupSets     set.Interface //好友 / 群
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	//1.  获取参数 并 检验 token 等合法性
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	fmt.Println("[ws] 0、 chat >>>> userId", userId)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//	context := query.Get("context")
	isvalida := true //checkToke()  待.........
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//2.获取conn
	currentTime := uint64(time.Now().Unix())
	node := &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(), //客户端地址
		HeartbeatTime: currentTime,                //心跳时间
		LoginTime:     currentTime,                //登录时间
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}

	//3. 用户关系
	//4. userid 跟 node绑定 并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	//5.完成发送逻辑
	go sendProc(node)

	//6.完成接受逻辑
	go recvProc(node)

	//sendMsg(userId, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws]4、sendProc >>>> msg :", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		//心跳检测 msg.Media == -1 || msg.Type == 3
		//if msg.Type == 3 {
		//	currentTime := uint64(time.Now().Unix())
		//	node.Heartbeat(currentTime)
		//} else {
		fmt.Println("[ws] 1、recvProc <<<<< ", string(data))
		dispatch(data)
		//	broadMsg(data) //todo 将消息广播到局域网

		//}

	}
}

func dispatch(data []byte) {
	msg := Message{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		fmt.Println("[ws] 2、dispatch  data :", string(data))
		sendMsg(msg.TargetId, data)
	case 2: //群发
		//sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case 4: // 心跳
		// 	node.Heartbeat()
		//case 4:
		//
	}
}

func sendMsg(targetId int64, msg []byte) {

	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()

	if ok {
		fmt.Println("[ws] 3、sendMsg >>> targetId: ", targetId, "  msg:", string(msg))
		node.DataQueue <- msg
	}

	//jsonMsg := Message{}
	//json.Unmarshal(msg, &jsonMsg)
	//ctx := context.Background()
	//targetIdStr := strconv.Itoa(int(userId))
	//userIdStr := strconv.Itoa(int(jsonMsg.UserId))
	//jsonMsg.CreateTime = uint64(time.Now().Unix())
	//r, err := utils.Red.Get(ctx, "online_"+userIdStr).Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if r != "" {
	//	if ok {
	//		fmt.Println("sendMsg >>> userID: ", userId, "  msg:", string(msg))
	//		node.DataQueue <- msg
	//	}
	//}
	//var key string
	//if userId > jsonMsg.UserId {
	//	key = "msg_" + userIdStr + "_" + targetIdStr
	//} else {
	//	key = "msg_" + targetIdStr + "_" + userIdStr
	//}
	//res, err := utils.Red.ZRevRange(ctx, key, 0, -1).Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//score := float64(cap(res)) + 1
	//ress, e := utils.Red.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	////res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	//if e != nil {
	//	fmt.Println(e)
	//}
	//fmt.Println(ress)
}
