package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want int) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw inssuficient", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)

		if err == nil {
			t.Errorf("Invalid withdrawn operation: had balance of 20, withdrew 100 and resulted in %d", wallet.Balance())
		}
	})
}
