// https://github.com/fyne-io/demo/blob/main/theme_test.go
package main

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

func TestForceVariantTheme_Color(t *testing.T) {
	test.NewApp()
	dark := &forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark}
	light := &forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight}
	unspecified := fyne.ThemeVariant(99)

	assert.Equal(t, theme.DefaultTheme().Color(theme.ColorNameForeground, theme.VariantDark), dark.Color(theme.ColorNameForeground, unspecified))
	assert.Equal(t, theme.DefaultTheme().Color(theme.ColorNameForeground, theme.VariantLight), light.Color(theme.ColorNameForeground, unspecified))

}
