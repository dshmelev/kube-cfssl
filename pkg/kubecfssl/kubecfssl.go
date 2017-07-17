package kubecfssl

import (
	"github.com/dshmelev/kube-cfssl/pkg/kubecfssl_const"
)

func New(version string) *KubeCfssl {
	return &KubeCfssl{
		version:   version,
		log:       makeLog(),
		stopCh:    make(chan struct{}),
		waitGroup: sync.WaitGroup{},
	}
}
