package tool

type Outcome struct {
	Output    string //节点输出结果 OR 日志
	Nextinput string //下一个节点输入
	Bitmap    byte   //节点状态
	/* 责任节点验证失败置为0,成功置为1;
	特别的如果需要跳转则 n 表示跳转一个n/2节点（取整）
	如果n为偶数则节点验证失败，反之则节点验证成功*/
	Goon bool //是否继续执行，true继续，false终止
}
