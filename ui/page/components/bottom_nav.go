package components

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/crypto-power/cryptopower/ui/cryptomaterial"
	"github.com/crypto-power/cryptopower/ui/load"
	"github.com/crypto-power/cryptopower/ui/values"
)

var (
	bottomNavigationBarHeight = values.MarginPadding100
)

type BottomNavigationBarHandler struct {
	Clickable     *cryptomaterial.Clickable
	Image         *cryptomaterial.Image
	ImageInactive *cryptomaterial.Image
	Title         string
	PageID        string
}

type BottomNavigationBar struct {
	*load.Load

	FloatingActionButton []BottomNavigationBarHandler
	BottomNaigationItems []BottomNavigationBarHandler
	CurrentPage          string

	axis        layout.Axis
	textSize    unit.Sp
	bottomInset unit.Dp
	height      unit.Dp
	alignment   layout.Alignment
	direction   layout.Direction
}

func (bottomNavigationbar *BottomNavigationBar) LayoutBottomNavigationBar(gtx layout.Context) layout.Dimensions {
	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return cryptomaterial.LinearLayout{
				Width:       cryptomaterial.WrapContent,
				Height:      cryptomaterial.WrapContent,
				Orientation: layout.Horizontal,
				Background:  bottomNavigationbar.Theme.Color.Surface,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					list := layout.List{Axis: layout.Horizontal}
					return list.Layout(gtx, len(bottomNavigationbar.BottomNaigationItems), func(gtx C, i int) D {
						return cryptomaterial.LinearLayout{
							Orientation: bottomNavigationbar.axis,
							Width:       bottomNavigationbar.Load.GetCurrentAppWidth() / len(bottomNavigationbar.BottomNaigationItems), // Divide each cell equally
							Height:      cryptomaterial.WrapContent,
							Padding:     layout.Inset{Bottom: values.MarginPadding10},
							Alignment:   bottomNavigationbar.alignment,
							Direction:   bottomNavigationbar.direction,
							Background:  bottomNavigationbar.Theme.Color.Surface,
							Clickable:   bottomNavigationbar.BottomNaigationItems[i].Clickable,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Stack{Alignment: layout.N}.Layout(gtx,
									layout.Stacked(func(gtx C) D {
										if bottomNavigationbar.BottomNaigationItems[i].Title == bottomNavigationbar.CurrentPage || bottomNavigationbar.BottomNaigationItems[i].PageID == bottomNavigationbar.CurrentPage {
											tabHeight := gtx.Dp(values.MarginPadding4)
											selectedTabDimsWidth := bottomNavigationbar.Load.GetCurrentAppWidth() / (len(bottomNavigationbar.BottomNaigationItems) * 8)
											tabRect := image.Rect(0, 0, 100+selectedTabDimsWidth, tabHeight)
											defer clip.RRect{Rect: tabRect, SE: 10, SW: 10, NW: 0, NE: 0}.Push(gtx.Ops).Pop()

											paint.FillShape(gtx.Ops, bottomNavigationbar.Theme.Color.Primary, clip.Rect(tabRect).Op())
											return layout.Dimensions{
												Size: image.Point{X: 100 + selectedTabDimsWidth, Y: tabHeight},
											}
										}

										return D{}
									}),
								)
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Inset{
									Top: values.MarginPadding8,
								}.Layout(gtx, func(gtx C) D {
									img := bottomNavigationbar.BottomNaigationItems[i].ImageInactive

									if bottomNavigationbar.BottomNaigationItems[i].Title == bottomNavigationbar.CurrentPage || bottomNavigationbar.BottomNaigationItems[i].PageID == bottomNavigationbar.CurrentPage {
										img = bottomNavigationbar.BottomNaigationItems[i].Image
									}

									return img.Layout24dp(gtx)
								})
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Inset{
									Bottom: bottomNavigationbar.bottomInset,
								}.Layout(gtx, func(gtx C) D {
									textColor := bottomNavigationbar.Theme.Color.GrayText1
									if bottomNavigationbar.BottomNaigationItems[i].Title == bottomNavigationbar.CurrentPage || bottomNavigationbar.BottomNaigationItems[i].PageID == bottomNavigationbar.CurrentPage {
										textColor = bottomNavigationbar.Theme.Color.DefaultThemeColors().Primary
									}
									txt := bottomNavigationbar.Theme.Label(bottomNavigationbar.textSize, bottomNavigationbar.BottomNaigationItems[i].Title)
									txt.Color = textColor
									return txt.Layout(gtx)
								})
							}),
						)
					})
				}),
			)
		}),
	)
}

func (bottomNavigationbar *BottomNavigationBar) LayoutSendReceive(gtx layout.Context) layout.Dimensions {
	gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
	if bottomNavigationbar.CurrentPage == "Overview" || bottomNavigationbar.CurrentPage == "Transactions" {
		return layout.S.Layout(gtx, func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					return layout.Center.Layout(gtx, func(gtx C) D {
						return cryptomaterial.LinearLayout{
							Width:       cryptomaterial.WrapContent,
							Height:      cryptomaterial.WrapContent,
							Orientation: layout.Horizontal,
							Background:  bottomNavigationbar.Theme.Color.DefaultThemeColors().Primary,
							Border:      cryptomaterial.Border{Radius: cryptomaterial.Radius(20)},
							Padding:     layout.UniformInset(values.MarginPadding8),
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return cryptomaterial.LinearLayout{
											Width:       cryptomaterial.WrapContent,
											Height:      cryptomaterial.WrapContent,
											Orientation: layout.Horizontal,
											Padding: layout.Inset{
												Right: values.MarginPadding16,
												Left:  values.MarginPadding16,
											},
											Clickable: bottomNavigationbar.FloatingActionButton[0].Clickable,
										}.Layout(gtx,
											layout.Rigid(func(gtx C) D {
												return layout.Inset{Right: values.MarginPadding8}.Layout(gtx,
													func(gtx C) D {
														return layout.Center.Layout(gtx, func(gtx C) D {
															return bottomNavigationbar.FloatingActionButton[0].Image.Layout24dp(gtx)
														})
													})
											}),
											layout.Rigid(func(gtx C) D {
												return layout.Inset{
													Left: values.MarginPadding0,
												}.Layout(gtx, func(gtx C) D {
													return layout.Center.Layout(gtx, func(gtx C) D {
														txt := bottomNavigationbar.Theme.Label(values.TextSize16, bottomNavigationbar.FloatingActionButton[0].Title)
														txt.Color = bottomNavigationbar.Theme.Color.DefaultThemeColors().White
														return txt.Layout(gtx)
													})
												})
											}),
										)
									}),
									layout.Rigid(func(gtx C) D {
										verticalSeparator := bottomNavigationbar.Theme.SeparatorVertical(50, 1)
										verticalSeparator.Color = bottomNavigationbar.Theme.Color.DefaultThemeColors().White
										return verticalSeparator.Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return cryptomaterial.LinearLayout{
											Width:       cryptomaterial.WrapContent,
											Height:      cryptomaterial.WrapContent,
											Orientation: layout.Horizontal,
											Padding: layout.Inset{
												Right: values.MarginPadding16,
												Left:  values.MarginPadding16,
											},
											Clickable: bottomNavigationbar.FloatingActionButton[1].Clickable,
										}.Layout(gtx,
											layout.Rigid(func(gtx C) D {
												return layout.Inset{Right: values.MarginPadding8}.Layout(gtx,
													func(gtx C) D {
														return layout.Center.Layout(gtx, func(gtx C) D {
															return bottomNavigationbar.FloatingActionButton[1].Image.Layout24dp(gtx)
														})
													})
											}),
											layout.Rigid(func(gtx C) D {
												return layout.Inset{
													Left: values.MarginPadding0,
												}.Layout(gtx, func(gtx C) D {
													return layout.Center.Layout(gtx, func(gtx C) D {
														txt := bottomNavigationbar.Theme.Label(values.TextSize16, bottomNavigationbar.FloatingActionButton[1].Title)
														txt.Color = bottomNavigationbar.Theme.Color.DefaultThemeColors().White
														return txt.Layout(gtx)
													})
												})
											}),
										)
									}),
								)
							}),
						)
					})
				}),
			)
		})
	}
	return D{}
}

func (bottomNavigationbar *BottomNavigationBar) OnViewCreated() {
	bottomNavigationbar.axis = layout.Vertical
	bottomNavigationbar.textSize = values.TextSize12
	bottomNavigationbar.bottomInset = values.MarginPadding0
	bottomNavigationbar.height = bottomNavigationBarHeight
	bottomNavigationbar.alignment = layout.Middle
	bottomNavigationbar.direction = layout.Center
}
