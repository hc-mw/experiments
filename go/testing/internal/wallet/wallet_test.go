package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		w := Wallet{Bitcoin(10)}
		err := w.Withdraw(Bitcoin(15))

		assertError(t, err, ErrInsufficientBalance)
		assertBalance(t, w, Bitcoin(10))
	})
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()

	if err == nil {
		t.Errorf("wanted an error but didn't get one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
