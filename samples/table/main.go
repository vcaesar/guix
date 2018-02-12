package main

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/drivers/gl"
	"github.com/vcaesar/guix/samples/flags"
)

func appMain(driver guix.Driver) {
	theme := flags.CreateTheme(driver)

	label1 := theme.CreateLabel()
	label1.SetColor(guix.White)
	label1.SetText("1x1")

	cell1x1 := theme.CreateLinearLayout()
	cell1x1.SetBackgroundBrush(guix.CreateBrush(guix.Blue40))
	cell1x1.SetHorizontalAlignment(guix.AlignCenter)
	cell1x1.AddChild(label1)

	label2 := theme.CreateLabel()
	label2.SetColor(guix.White)
	label2.SetText("2x1")

	cell2x1 := theme.CreateLinearLayout()
	cell2x1.SetBackgroundBrush(guix.CreateBrush(guix.Green40))
	cell2x1.SetHorizontalAlignment(guix.AlignCenter)
	cell2x1.AddChild(label2)

	label3 := theme.CreateLabel()
	label3.SetColor(guix.White)
	label3.SetText("1x2")

	cell1x2 := theme.CreateLinearLayout()
	cell1x2.SetBackgroundBrush(guix.CreateBrush(guix.Red40))
	cell1x2.SetHorizontalAlignment(guix.AlignCenter)
	cell1x2.AddChild(label3)

	table := theme.CreateTableLayout()
	table.SetGrid(3, 2) // columns, rows

	// row, column, horizontal span, vertical span
	table.SetChildAt(0, 0, 1, 1, cell1x1)
	table.SetChildAt(0, 1, 2, 1, cell2x1)
	table.SetChildAt(2, 0, 1, 2, cell1x2)

	window := theme.CreateWindow(800, 600, "Table")
	window.AddChild(table)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
