package test

import (
	"project/proxy_pool/api"
	"testing"
)

func TestWorker(t *testing.T) {
	api.StartWorker()
}
