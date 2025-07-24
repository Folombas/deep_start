package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	// Инициализация библиотеки
	if err := ui.Init(); err != nil {
		log.Fatalf("Ошибка инициализации termui: %v", err)
	}
	defer ui.Close()

	// 1. Виджет для вывода текста
	header := widgets.NewParagraph()
	header.Text = "ДЕМОНСТРАЦИЯ TERMUI\nУправление: ← → - переключать виджеты, q - выход"
	header.Border = true
	header.Title = "Управление"
	header.SetRect(0, 0, 80, 5)
	header.TextStyle.Fg = ui.ColorYellow

	// 2. Линейный график
	plot := widgets.NewPlot()
	plot.Title = "График функций: sin(x) и cos(x)"
	plot.Data = [][]float64{}
	plot.SetRect(0, 5, 80, 25)
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen
	plot.LineColors[1] = ui.ColorRed

	// Генерация данных для графика
	sinData := make([]float64, 100)
	cosData := make([]float64, 100)
	for i := 0; i < 100; i++ {
		x := float64(i) * 2 * math.Pi / 100
		sinData[i] = math.Sin(x)
		cosData[i] = math.Cos(x)
	}
	plot.Data = append(plot.Data, sinData, cosData)

	// 3. Круговая диаграмма
	gauge := widgets.NewGauge()
	gauge.Title = "Прогресс"
	gauge.Percent = rand.Intn(100)
	gauge.SetRect(0, 25, 40, 28)
	gauge.BarColor = ui.ColorBlue
	gauge.BorderStyle.Fg = ui.ColorCyan

	// 4. Список
	list := widgets.NewList()
	list.Title = "Список элементов"
	list.Rows = []string{
		"[1] Первый элемент",
		"[2] Второй элемент",
		"[3] Третий элемент",
		"[4] Четвертый элемент",
		"[5] Пятый элемент",
	}
	list.SetRect(40, 25, 80, 32)
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)

	// 5. Таблица
	table := widgets.NewTable()
	table.Title = "Таблица данных"
	table.Rows = [][]string{
		{"ID", "Товар", "Цена"},
		{"1", "Ноутбук", "1200$"},
		{"2", "Смартфон", "800$"},
		{"3", "Планшет", "600$"},
	}
	table.SetRect(0, 28, 80, 35)
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.RowSeparator = true
	table.BorderStyle = ui.NewStyle(ui.ColorMagenta)

	// Сборка интерфейса
	grid := ui.NewGrid()
	grid.SetRect(0, 0, 80, 35)
	grid.Set(
		ui.NewRow(1.0/7,
			ui.NewCol(1.0, header),
		),
		ui.NewRow(4.0/7,
			ui.NewCol(1.0, plot),
		),
		ui.NewRow(1.0/7,
			ui.NewCol(0.5, gauge),
			ui.NewCol(0.5, list),
		),
		ui.NewRow(1.0/7,
			ui.NewCol(1.0, table),
		),
	)

	// Обработка событий клавиатуры
	activeWidget := 0
	widgets := []ui.Drawable{grid}
	ui.Render(widgets...)

	eventLoop := func() {
		uiEvents := ui.PollEvents()
		ticker := time.NewTicker(time.Second).C
		update := time.NewTicker(500 * time.Millisecond).C

		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "<Resize>":
					payload := e.Payload.(ui.Resize)
					grid.SetRect(0, 0, payload.Width, payload.Height)
					ui.Clear()
					ui.Render(widgets...)
				case "<Left>":
					activeWidget = (activeWidget + len(widgets) - 1) % len(widgets)
				case "<Right>":
					activeWidget = (activeWidget + 1) % len(widgets)
				}
			case <-ticker:
				header.Text = fmt.Sprintf("ДЕМОНСТРАЦИЯ TERMUI | Время: %s", time.Now().Format("15:04:05"))
				ui.Render(widgets...)
			case <-update:
				gauge.Percent = (gauge.Percent + 5) % 100
				ui.Render(widgets...)
			}

			ui.Render(widgets...)
		}
	}

	eventLoop()
}
