package main

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

func main() {
	static.Install()
	setting.Install()
	atomic.Engine.Run()
}
