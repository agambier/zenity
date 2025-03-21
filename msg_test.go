package zenity_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/agambier/zenity/v23"
	"go.uber.org/goleak"
)

func ExampleError() {
	zenity.Error("An error has occurred.",
		zenity.Title("Error"),
		zenity.ErrorIcon)
}

func ExampleInfo() {
	zenity.Info("All updates are complete.",
		zenity.Title("Information"),
		zenity.InfoIcon)
}

func ExampleWarning() {
	zenity.Warning("Are you sure you want to proceed?",
		zenity.Title("Warning"),
		zenity.WarningIcon)
}

func ExampleQuestion() {
	zenity.Question("Are you sure you want to proceed?",
		zenity.Title("Question"),
		zenity.QuestionIcon)
}

var msgFuncs = []struct {
	name string
	fn   func(string, ...zenity.Option) error
}{
	{"Error", zenity.Error},
	{"Info", zenity.Info},
	{"Warning", zenity.Warning},
	{"Question", zenity.Question},
}

func TestMessage_timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	for _, tt := range msgFuncs {
		t.Run(tt.name, func(t *testing.T) {
			defer goleak.VerifyNone(t)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second/5)
			defer cancel()

			err := tt.fn("text", zenity.Context(ctx))
			if skip, err := skip(err); skip {
				t.Skip("skipping:", err)
			}
			if !os.IsTimeout(err) {
				t.Error("did not timeout:", err)
			}
		})
	}
}

func TestMessage_cancel(t *testing.T) {
	for _, tt := range msgFuncs {
		t.Run(tt.name, func(t *testing.T) {
			defer goleak.VerifyNone(t)
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			err := tt.fn("text", zenity.Context(ctx))
			if skip, err := skip(err); skip {
				t.Skip("skipping:", err)
			}
			if !errors.Is(err, context.Canceled) {
				t.Error("was not canceled:", err)
			}
		})
	}
}

func TestMessage_script(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	for _, tt := range msgFuncs {
		t.Run(tt.name+"OK", func(t *testing.T) {
			err := tt.fn("Please, press OK.", zenity.OKLabel("OK"))
			if skip, err := skip(err); skip {
				t.Skip("skipping:", err)
			}
			if err != nil {
				t.Errorf("%s() = %v; want nil", tt.name, err)
			}
		})
		t.Run(tt.name+"Extra", func(t *testing.T) {
			err := tt.fn("Please, press Extra.", zenity.ExtraButton("Extra"))
			if skip, err := skip(err); skip {
				t.Skip("skipping:", err)
			}
			if err != zenity.ErrExtraButton {
				t.Errorf("%s() = %v; want %v", tt.name, err, zenity.ErrExtraButton)
			}
		})
	}

	tests := []struct {
		name string
		call string
		err  error
	}{
		{name: "QuestionYes", call: "press Yes"},
		{name: "QuestionNo", call: "press No", err: zenity.ErrExtraButton},
		{name: "QuestionCancel", call: "cancel", err: zenity.ErrCanceled},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := zenity.Question(fmt.Sprintf("Please, %s.", tt.call),
				zenity.OKLabel("Yes"),
				zenity.ExtraButton("No"),
				zenity.CancelLabel("Cancel"))
			if skip, err := skip(err); skip {
				t.Skip("skipping:", err)
			}
			if err != tt.err {
				t.Errorf("Question() = %v; want %v", err, tt.err)
			}
		})
	}
}

func TestMessage_icon(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	err := zenity.Question("Does this dialog have an error icon?",
		zenity.OKLabel("Yes"),
		zenity.CancelLabel("No"),
		zenity.ErrorIcon)
	if skip, err := skip(err); skip {
		t.Skip("skipping:", err)
	}
	if err != nil {
		t.Errorf("Question() = %v; want nil", err)
	}
}

func TestMessage_customIcon(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	err := zenity.Question("Does this dialog have a custom icon?",
		zenity.Title(""),
		zenity.OKLabel("Yes"),
		zenity.CancelLabel("No"),
		zenity.Icon("testdata/icon.png"))
	if skip, err := skip(err); skip {
		t.Skip("skipping:", err)
	}
	if err != nil && (runtime.GOOS == "windows" || runtime.GOOS == "darwin") {
		t.Errorf("Question() = %v; want nil", err)
	}
}

func TestMessage_callbacks(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	for i := 0; i < 2000; i++ {
		zenity.Error("text", zenity.Context(ctx))
	}
}
