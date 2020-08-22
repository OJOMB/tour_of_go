package main

type ItemInterface interface {
	GetItemValue() string
}

type Item struct {
	ID int
}

type URLItem struct {
	Item
	URL string
}

type TextItem struct {
	Item
	Text string
}

func (ui URLItem) GetItemValue() {
	return ui.URL
}

func (ti TextItem) GetItemValue() {
	return ti.Text
}

func FindItem(ID int) ItemInterface {
	// ...
}
