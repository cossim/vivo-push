package vivopush

type ResultItem struct {
	Result int    `json:"result"`
	Desc   string `json:"desc"`
}

type SendResult struct {
	ResultItem
	TaskId string `json:"taskId"`
}

type BatchStatusResult struct {
	ResultItem
	Statistics []TaskData `json:"statistics"`
}

type TaskData struct {
	TaskId         string `json:"taskId"`
	Send           int    `json:"send"`
	Receive        int    `json:"receive"`
	Display        int    `json:"display"`
	Click          int    `json:"click"`
	TargetInvalid  int    `json:"targetInvalid"`
	TargetUnSub    int    `json:"targetUnSub"`
	TargetInActive int    `json:"targetInActive"`
	Covered        int    `json:"covered"`
	Controlled     int    `json:"controlled"`
	TargetOffline  int    `json:"targetOffline"`
}
