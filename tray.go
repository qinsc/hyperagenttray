package main

import (
	"fmt"

	xcgui "github.com/codyguo/xcgui/xc"
	"github.com/lxn/walk"
)

var (
	hWindow    xcgui.HWINDOW
	notifyIcon *walk.NotifyIcon
)

func createTray() {
	hWindow = xcgui.XWnd_Create(0, 0, 0, 0, "HyperAgentTray", 0, xcgui.XC_WINDOW_STYLE_NOTHING)

	// We load our icon from a file.
	icon, err := walk.NewIconFromFile("tray.ico")
	if err != nil {
		fmt.Printf("NewIconFromFile error: %v", err)
		return
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon()
	if err != nil {
		fmt.Printf("NewNotifyIcon error: %v", err)
		return
	}
	defer ni.Dispose()

	notifyIcon = ni

	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		fmt.Printf("SetIcon error: %v", err)
	}
	if err := ni.SetToolTip("HyperAgent"); err != nil {
		fmt.Printf("SetToolTip error: %v", err)
	}

	// When the left mouse button is pressed, bring up our balloon.
	//	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
	//		if button != walk.LeftButton {
	//			return
	//		}

	//		if err := ni.ShowCustom(
	//			"自定义消息",
	//			"这是一个带图标的自定义消息."); err != nil {
	//			fmt.Printf("ShowCustom error: %v", err)
	//		}
	//	})

	// 菜单使用walk的，主程序为xcgui.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出"); err != nil {
		fmt.Printf("exitAction.SetText error: %v", err)
	}
	exitAction.Triggered().Attach(func() {
		ni.Dispose()
		walk.App().Exit(0)
		xcgui.XExitXCGUI()
	})

	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		fmt.Printf("ContextMenu().Actions().Add(exitAction) error: %v", err)
	}

	// 托盘图标默认为隐藏状态，需设置为显示。
	if err := ni.SetVisible(true); err != nil {
		fmt.Printf("SetVisible error: %v", err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("HyperAgent", "正在运行中."); err != nil {
		fmt.Printf("ShowInfo error: %v", err)
	}

	// Run the message loop.
	xcgui.XRunXCGUI()
}

func showMessageAll(msg string) {
	showMessage(msg, true)
}

func showMessage(msg string, showBox bool) {
	if err := notifyIcon.ShowInfo("信息", msg); err != nil {
		fmt.Printf("ShowInfo error: %v", err)
	}

	if showBox {
		fmt.Println("Do Show MessageBox ...")
		go func() {
			xcgui.MessageBox(xcgui.XWnd_GetHWND(hWindow), "提示信息", msg, xcgui.MB_ICONINFORMATION)
		}()
	}
}
