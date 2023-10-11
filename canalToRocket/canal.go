package canalToRocket

import (
	"fmt"
	"github.com/withlin/canal-go/client"
	"github.com/withlin/canal-go/protocol/entry"
	"time"
)

func init() {
	connector := client.NewSimpleCanalConnector("你的Canal服务端地址", 11111, "你的Canal服务端用户名", "你的Canal服务端密码", "你的Canal服务端destination", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		fmt.Println(err)
	}

	//mysql 数据解析关注的表，Perl正则表达式. 这里过滤你需要获取数据的表
	err = connector.Subscribe("你的数据库\\.你的表")
	if err != nil {
		fmt.Println(err)
	}

	for {
		message, err := connector.Get(100, nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		messageId := message.Id
		if messageId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		//处理信息
		doMessage(message.Entries)
	}
}

// 处理收到的推送
func doMessage(entrys []entry.Entry) {
	for _, entry := range entrys {
		fmt.Printf("entry %v\n", entry)
	}
	//for _, entry := range entrys {
	//	if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
	//		continue
	//	}
	//	rowChange := new(protocol.RowChange)
	//
	//	err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	if rowChange != nil {
	//		eventType := rowChange.GetEventType()
	//		header := entry.GetHeader()
	//
	//		if header.ExecuteTime < time.Now().UnixNano()/1e6 {
	//			//变动时间早于当前时间 直接抛弃这条数据
	//			continue
	//		}
	//		for _, rowData := range rowChange.GetRowDatas() {
	//			if eventType == protocol.EventType_DELETE {
	//				//todo 收到数据库删除数据推送，执行对应处理，rowData.GetBeforeColumns() 可以获取到删除的数据
	//				//rowData.GetBeforeColumns()
	//			} else if eventType == protocol.EventType_INSERT {
	//				//todo 收到数据库新增数据推送，执行对应处理，rowData.GetAfterColumns() 可以获取到新增的数据
	//				//rowData.GetAfterColumns()
	//			} else {
	//				//todo 收到数据库更新数据推送，执行对应处理，rowData.GetAfterColumns() 可以获取到更新后的数据，rowData.GetBeforeColumns() 可以获取到更新前的数据
	//				//rowData.GetBeforeColumns()
	//				//rowData.GetAfterColumns()
	//			}
	//		}
	//	}
	//}
}
