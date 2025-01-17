package config

type ConfigDataFalcoMock struct {
}

func (c *ConfigDataFalcoMock) IsFalcoEbpfEngine() bool {
	return true
}

func (c *ConfigDataFalcoMock) GetFalcoSyscallFilter() []string {
	return []string{"open", "openat", "execve", "execveat"}
}

func (c *ConfigDataFalcoMock) GetFalcoKernelObjPath() string {
	return "./../../resources/ebpf/mock_falco_ebpf_engine/kernel_obj_mock.o"
}

func (c *ConfigDataFalcoMock) GetEbpfEngineLoaderPath() string {
	return "./../../resources/ebpf/mock_falco_ebpf_engine/userspace_app_mock"
}
