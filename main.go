// Copyright 2021 The timecrystal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/itsubaki/q"
)

func main() {
	rand.Seed(1)

	qsim := q.New()

	q0 := qsim.Zero()
	q1 := qsim.One()
	q2 := qsim.Zero()
	q3 := qsim.One()

	// https://arxiv.org/abs/2105.06632
	period := func() {
		e := .5
		qsim.RY(math.Pi-e, q0)
		qsim.RY(math.Pi-e, q1)
		qsim.RY(math.Pi-e, q2)
		qsim.RY(math.Pi-e, q3)

		qsim.CNOT(q0, q1)
		qsim.CNOT(q2, q3)

		j := func() float64 {
			return math.Pi * (1.0 + 3*rand.Float64()) / 8
		}
		qsim.RZ(-2*j(), q1)
		qsim.RZ(-2*j(), q3)

		qsim.CNOT(q0, q1)
		qsim.CNOT(q2, q3)
		qsim.CNOT(q1, q2)

		qsim.RZ(-2*j(), q2)

		qsim.CNOT(q1, q2)
	}
	for i := 0; i < 8; i++ {
		period()
		max, binary := 0.0, []string{}
		for _, state := range qsim.State() {
			if state.Probability > max {
				max, binary = state.Probability, state.BinaryString
			}
		}
		fmt.Println(binary, max)
	}
}
