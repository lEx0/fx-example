// Copyright (c) 2021 Amangeldy Kadyl
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package usecase

import (
	"fmt"
	"github.com/lEx0/fx-example/pkg/weather"
)

// GreetUser приветствуем пользователя
func GreetUser(name string, city string, weather *weather.Client) (string, error) {
	var additional string

	current, err := weather.Current(city)
	if err != nil {
		return "", err
	}

	if len(current.Weather) > 0 {
		additional = ", " + current.Weather[0].Description
	}
	if name != "" {
		name = " " + name
	}

	return fmt.Sprintf(
		"Привествую вас%s.\nВ городе %s сейчас на улице %d градусов%s.",
		name, current.Name, int(current.Main.Temp), additional,
	), nil
}
