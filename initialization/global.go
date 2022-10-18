package initialization

func Global() {
	Config()  // 配置文件
	Mysql()   // Mysql初始化
	Elastic() // Es初始化
}
