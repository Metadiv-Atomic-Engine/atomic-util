package main

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

func main() {
	static.Install()
	atomic.Engine.Run()
}
