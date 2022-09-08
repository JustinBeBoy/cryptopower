package staking

import (
	"context"
	"strconv"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"

	"gitlab.com/raedah/cryptopower/ui/decredmaterial"
	"gitlab.com/raedah/cryptopower/ui/load"
	"gitlab.com/raedah/cryptopower/ui/page/components"
	"gitlab.com/raedah/cryptopower/ui/values"
	"gitlab.com/raedah/libwallet"
)

type ticketBuyerModal struct {
	*load.Load
	*decredmaterial.Modal

	ctx       context.Context // page context
	ctxCancel context.CancelFunc

	settingsSaved func()
	onCancel      func()

	cancel          decredmaterial.Button
	saveSettingsBtn decredmaterial.Button

	balToMaintainEditor decredmaterial.Editor

	accountSelector *components.AccountSelector
	vspSelector     *components.VSPSelector
}

func newTicketBuyerModal(l *load.Load) *ticketBuyerModal {
	tb := &ticketBuyerModal{
		Load:  l,
		Modal: l.Theme.ModalFloatTitle("staking_modal"),

		cancel:          l.Theme.OutlineButton(values.String(values.StrCancel)),
		saveSettingsBtn: l.Theme.Button(values.String(values.StrSave)),
		vspSelector:     components.NewVSPSelector(l).Title(values.String(values.StrSelectVSP)),
	}

	tb.balToMaintainEditor = l.Theme.Editor(new(widget.Editor), values.String(values.StrBalToMaintain))
	tb.balToMaintainEditor.Editor.SingleLine = true

	tb.saveSettingsBtn.SetEnabled(false)

	return tb
}

func (tb *ticketBuyerModal) OnSettingsSaved(settingsSaved func()) *ticketBuyerModal {
	tb.settingsSaved = settingsSaved
	return tb
}

func (tb *ticketBuyerModal) OnCancel(cancel func()) *ticketBuyerModal {
	tb.onCancel = cancel
	return tb
}

func (tb *ticketBuyerModal) OnResume() {
	tb.initializeAccountSelector()
	tb.ctx, tb.ctxCancel = context.WithCancel(context.TODO())
	tb.accountSelector.ListenForTxNotifications(tb.ctx, tb.ParentWindow())

	if len(tb.WL.MultiWallet.KnownVSPs()) == 0 {
		// TODO: Does this modal need this list?
		go tb.WL.MultiWallet.ReloadVSPList(context.TODO())
	}

	// loop through all available wallets and select the one with ticket buyer config.
	// if non, set the selected wallet to the first.
	// temporary work around for only one wallet.
	if tb.WL.SelectedWallet.Wallet.TicketBuyerConfigIsSet() {
		tbConfig := tb.WL.SelectedWallet.Wallet.AutoTicketsBuyerConfig()
		acct, err := tb.WL.SelectedWallet.Wallet.GetAccount(tbConfig.PurchaseAccount)
		if err != nil {
			tb.Toast.NotifyError(err.Error())
		}

		if tb.WL.SelectedWallet.Wallet.ReadBoolConfigValueForKey(libwallet.AccountMixerConfigSet, false) &&
			!tb.WL.SelectedWallet.Wallet.ReadBoolConfigValueForKey(load.SpendUnmixedFundsKey, false) &&
			(tbConfig.PurchaseAccount == tb.WL.SelectedWallet.Wallet.MixedAccountNumber()) {
			tb.accountSelector.SetSelectedAccount(acct)
		} else {
			err := tb.accountSelector.SelectFirstWalletValidAccount()
			if err != nil {
				tb.Toast.NotifyError(err.Error())
			}
		}

		tb.vspSelector.SelectVSP(tbConfig.VspHost)
		tb.balToMaintainEditor.Editor.SetText(strconv.FormatFloat(libwallet.AmountCoin(tbConfig.BalanceToMaintain), 'f', 0, 64))
	}

	if tb.accountSelector.SelectedAccount() == nil {
		err := tb.accountSelector.SelectFirstWalletValidAccount()
		if err != nil {
			tb.Toast.NotifyError(err.Error())
		}
	}
}

func (tb *ticketBuyerModal) Layout(gtx layout.Context) layout.Dimensions {
	l := []layout.Widget{
		func(gtx C) D {
			t := tb.Theme.H6(values.String(values.StrAutoTicketPurchase))
			t.Font.Weight = text.SemiBold
			return t.Layout(gtx)
		},
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{
						Top:    values.MarginPadding8,
						Bottom: values.MarginPadding16,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return tb.accountSelector.Layout(tb.ParentWindow(), gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return tb.balToMaintainEditor.Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{
						Top:    values.MarginPadding16,
						Bottom: values.MarginPadding16,
					}.Layout(gtx, func(gtx C) D {
						return tb.vspSelector.Layout(tb.ParentWindow(), gtx)
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.E.Layout(gtx, func(gtx C) D {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return layout.Inset{
							Right: values.MarginPadding4,
						}.Layout(gtx, tb.cancel.Layout)
					}),
					layout.Rigid(func(gtx C) D {
						return tb.saveSettingsBtn.Layout(gtx)
					}),
				)
			})
		},
	}

	return tb.Modal.Layout(gtx, l)
}

func (tb *ticketBuyerModal) canSave() bool {
	if tb.vspSelector.SelectedVSP() == nil {
		return false
	}

	if tb.balToMaintainEditor.Editor.Text() == "" {
		return false
	}

	return true
}

func (tb *ticketBuyerModal) initializeAccountSelector() {
	tb.accountSelector = components.NewAccountSelector(tb.Load).
		Title(values.String(values.StrPurchasingAcct)).
		AccountSelected(func(selectedAccount *libwallet.Account) {}).
		AccountValidator(func(account *libwallet.Account) bool {
			// Imported and watch only wallet accounts are invalid for sending
			accountIsValid := account.Number != libwallet.ImportedAccountNumber && !tb.WL.SelectedWallet.Wallet.IsWatchingOnlyWallet()

			if tb.WL.SelectedWallet.Wallet.ReadBoolConfigValueForKey(libwallet.AccountMixerConfigSet, false) &&
				!tb.WL.SelectedWallet.Wallet.ReadBoolConfigValueForKey(load.SpendUnmixedFundsKey, false) {
				// Spending from unmixed accounts is disabled for the selected wallet
				accountIsValid = account.Number == tb.WL.SelectedWallet.Wallet.MixedAccountNumber()
			}

			return accountIsValid
		})
}

func (tb *ticketBuyerModal) OnDismiss() {
	tb.ctxCancel()
}

func (tb *ticketBuyerModal) Handle() {
	tb.saveSettingsBtn.SetEnabled(tb.canSave())

	if tb.cancel.Clicked() || tb.Modal.BackdropClicked(true) {
		tb.onCancel()
		tb.Dismiss()
	}

	if tb.saveSettingsBtn.Clicked() {
		vspHost := tb.vspSelector.SelectedVSP().Host
		amount, err := strconv.ParseFloat(tb.balToMaintainEditor.Editor.Text(), 64)
		if err != nil {
			tb.Toast.NotifyError(err.Error())
			return
		}

		balToMaintain := libwallet.AmountAtom(amount)
		account := tb.accountSelector.SelectedAccount()

		tb.WL.SelectedWallet.Wallet.SetAutoTicketsBuyerConfig(vspHost, account.Number, balToMaintain)
		tb.settingsSaved()
		tb.Dismiss()
	}
}
