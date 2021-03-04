// Copyright (c) 2021, Ricard Bejarano. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license, which can be found in the LICENSE file.

package schema

import (
	"errors"
	"time"
)

type Errors []struct {
	Message    string
	Path       []string
	Extensions struct {
		Timestamp time.Time
	}
}

// Slice returns a []error containing error structures holding the Message
// attribute of all errors in `E`.
func (E Errors) Slice() []error {
	var errs []error
	for _, e := range E {
		errs = append(errs, errors.New(e.Message))
	}
	return errs
}
