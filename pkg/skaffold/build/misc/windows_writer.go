/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package misc

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
)

type windowsToUTF8Writer struct {
	wrappedWriter io.Writer
}

func NewWindowsToUTF8Writer(w io.Writer) io.Writer {
	return &windowsToUTF8Writer{wrappedWriter: w}
}

func (w *windowsToUTF8Writer) Write(p []byte) (int, error) {
	decoder := charmap.CodePage850.NewDecoder()

	transformedWriter := transform.NewWriter(w.wrappedWriter, decoder)

	n, err := transformedWriter.Write(p)
	if err != nil {
		return n, err
	}

	err = transformedWriter.Close()
	return n, err
}
