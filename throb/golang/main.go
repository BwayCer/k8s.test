// 心跳

// 外部依賴：
//   /bin/stty

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
	checkTerminalSize()
}

func main() {
	// TODO
	// 仍有機會在 <Ctrl+C> 退出時出現以下訊息
	// 推測是 <Ctrl+C> 按下時已讓終端機結束，使得 `stty size` 取值錯誤
	// panic: signal: interrupt
	// goroutine 1 [running]:
	// main.terminalSize(0xc0004562d4, 0xc000124870)
	//    /main.go:140 +0x24f
	// main.throb(0x77)
	//    /main.go:92 +0x246
	// main.main()
	//    /main.go:45 +0xdc
	// exit status 2
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-osSignalChan
		fmt.Println()
		// Golang 只要使用 <Ctrl+C> 都會以代碼 1 退出
		os.Exit(0)
	}()

	var loop int = 0

	insCmdFlag := handleFlag()
	if insCmdFlag.Arrhythmia {
		throb_arrhythmia()
	}

	for {
		throb(loop)
		loop++
		time.Sleep(time.Duration(monitorRefreshPeriod) * time.Millisecond)
	}

	// for _ = range time.Tick(time.Duration(monitorRefreshPeriod) * time.Millisecond) {
	// 	throb(loop)
	// 	loop++
	// }
}

type CmdFlag struct {
	Arrhythmia bool
}

var ynCanUseStty = true

func handleFlag() CmdFlag {
	insCmdFlag := CmdFlag{}
	flag.BoolVar(&insCmdFlag.Arrhythmia, "a", false, "to make arrhythmia.")
	flag.BoolVar(&insCmdFlag.Arrhythmia, "arrhythmia", false, "same as -a. to make arrhythmia.")
	flag.Parse()
	return insCmdFlag
}

// monitorClear() {
// 	fmt.Printf("\033[H\033[2J")
// }

var throbRateCode = []byte{0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1}
var throbRateCodeLength = len(throbRateCode)
var throbSymbol = []rune("⠤⣄⣀⣠⠤⠖⠒⠋⠉⠙⠒⠲")
var throbSymbolLength = len(throbSymbol)
var monitorRefreshPeriod = 16
var monitorGraph = make([]rune, 0, 300)
var arrhythmiaExtent = 99

func throb(loop int) {
	rateIdx := (loop / throbSymbolLength) % throbRateCodeLength

	if throbRateCode[rateIdx] == 0 {
		monitorGraph = append([]rune{throbSymbol[0]}, monitorGraph...)
	} else {
		symbolIdx := throbSymbolLength - 1 - (loop % throbSymbolLength)
		monitorGraph = append([]rune{throbSymbol[symbolIdx]}, monitorGraph...)
	}

	var columns, cutLength int
	if ynCanUseStty {
		_, columns = terminalSize()
	} else {
		columns = 32
	}

	if columns >= 64 {
		cutLength = 58
	} else {
		cutLength = columns - 6
	}
	if len(monitorGraph) > cutLength {
		monitorGraph = monitorGraph[0:cutLength]
	}

	fmt.Printf("\r\033[K%s", string(monitorGraph))
}

func throb_arrhythmia() {
	newThrobRateCode := make([]byte, 99)

	for idx, leng := 0, arrhythmiaExtent; idx < leng; idx++ {
		newThrobRateCode[idx] = byte(rand.Intn(2))
	}

	throbRateCode = newThrobRateCode
	throbRateCodeLength = len(throbRateCode)
}

func checkTerminalSize() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		ynCanUseStty = false
	}
	outputStr := strings.Trim(string(output), " \n")
	ynMatch, _ := regexp.MatchString("^[0-9]+ [0-9]+$", outputStr)
	if !ynMatch {
		ynCanUseStty = false
	}
}

func terminalSize() (lines, columns int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	cutList := strings.Split(strings.Trim(string(output), " \n"), " ")
	lines, err = strconv.Atoi(cutList[0])
	if err != nil {
		panic(err)
	}
	columns, err = strconv.Atoi(cutList[1])
	if err != nil {
		panic(err)
	}
	return
}
