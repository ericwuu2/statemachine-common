package statemachine_common

import (
	"fmt"
	"github.com/Gurpartap/statemachine-go"
)

type GameProcess struct {
	statemachine.Machine
	IsLogin      bool
	IsDisconnect bool
}

func (gameProcess *GameProcess) NotifyTriggers() {
	fmt.Println("NotifyTriggers")
}
func (gameProcess *GameProcess) GetIsLogin() bool {
	return gameProcess.IsLogin
}
