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

package main

import (
	"github.com/lEx0/fx-example/app"
	"github.com/lEx0/fx-example/pkg/config"
	"github.com/lEx0/fx-example/pkg/weather"

	"github.com/lEx0/fx-http-module"
	"github.com/lEx0/fx-logger-module"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		// можем добавляем fx.NopLogger что бы не выводить отладочную информацию Fx
		// fx.NopLogger,

		// добавляем нужны модули
		config.Module,
		logger.Module,
		http.Module,
		weather.Module,

		// добавляем модуль самого приложения
		app.Module,
	).Run() // запускам приложение
}
