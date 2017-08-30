package nanogui

import (
//	"github.com/gianpaolog/nanovgo"
)

type TabHeader struct {
	WidgetImplement
	tabButtons               []*TabButton
	activeTab                int
	callback                 func(index int)
	visibleStart, visibleEnd int
}

func NewTabHeader(parent Widget) *TabHeader {
	header := &TabHeader{}

	InitWidget(header, parent)
	return header
}

func (t *TabHeader) SetActiveTab(index int) {
	if index < 0 || index >= len(t.tabButtons) {
		return
	}
	t.activeTab = index

	if t.callback == nil {
		return
	}
	t.callback(index)
}

func (t *TabHeader) ActiveTab() int {
	return t.activeTab
}

func (t *TabHeader) isVisibleTab(index int) bool {
	return index >= t.visibleStart && index < t.visibleEnd
}

func (t *TabHeader) AddTab(index int, label string) {
	if index > len(t.tabButtons) {
		index = len(t.tabButtons)
	}

	tab := NewTabButton(t, label)
	t.tabButtons = append(t.tabButtons, nil)
	copy(t.tabButtons[index+1:], t.tabButtons[index:])
	t.tabButtons[index] = tab

	t.SetActiveTab(index)
}

func (t *TabHeader) RemoveTab(index int) {
	copy(t.tabButtons[index:], t.tabButtons[index+1:])
	t.tabButtons[len(t.tabButtons)-1] = nil
	t.tabButtons = t.tabButtons[:len(t.tabButtons)-1]
}

func (t *TabHeader) tabIndex(label string) (index int, ok bool) {
	for i := range t.tabButtons {
		if t.tabButtons[i].Label == label {
			return i, true
		}
	}
	return -1, false
}

func (t *TabHeader) TabLabelAt(index int) string {
	if index < 0 || index >= len(t.tabButtons) {
		return ""
	}
	return t.tabButtons[index].Label
}
