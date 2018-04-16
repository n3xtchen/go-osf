package register

// 服务注册的抽象方法
type Register struct {
}

// 注册配置列表，抽象方法（需要子类实现）
func (this *Register) ListRegistryConfig() []RegistryConf {
	panic("Not Implemented")
}

// 注册服务, 同时订阅或者登记数据
func (this *Register) Register(data string) {
	for _, registryConf := range this.ListRegistryConfig() {
		this.CreateNode(registryConf, data)
	}
}

// 创建服务结点
func (this *Register) CreateNode(registryConf RegistryConf, data string) {
	// zkClient := registryConf.zkClient

	// pathArr := strings.Split(this.GetPath(), "/")

	panic("xx")
	// 递归创建结点
	// CreateIfNotExists := func(pathSeg []string, path string) {
	//     if len(pathSeg) == 0 {
	//         return
	//     } else {
	//         newPath := path+pathSeg[0]
	//
	//         path, _, err := zkClient.Exists(newPath)
	//         if err == nil  {
	//             flags := int32(0)
	//             acl := zk.WorldACL(zk.PermAll)
	//             zkClient.Create(newPath, []byte(pathSeg[0]), flags, acl)
	//         }
	//         CreateIfNotExists(pathSeg[1:], newPath)
	//     }
	// }
	// CreateIfNotExists(pathArr[1:], pathArr[0])
}

func (this *Register) UnRegister() {
	for _, registryConf := range this.ListRegistryConfig() {
		registryConf.zkClient.Delete(this.GetPath(), -1)
	}
}

// 获取服务注册路径
func (this *Register) GetPath() string {
	panic("Not Implemented")
}

func (this *Register) Close() {
	for _, registryConf := range this.ListRegistryConfig() {
		registryConf.zkClient.Close()
	}
}
