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
	inits := []string{
		"0000",
		"0001",
		"0010",
		"0011",
		"0100",
		"0101",
		"0110",
		"0111",
		"1000",
		"1001",
		"1010",
		"1011",
		"1100",
		"1101",
		"1110",
		"1111",
	}
	for _, init := range inits {
		fmt.Println(init, findPeriod(init))
	}
}

func findPeriod(init string) float64 {
	rand.Seed(1)

	qsim := q.New()

	q := []q.Qubit{}
	for _, state := range init {
		if state == '0' {
			q = append(q, qsim.Zero())
		} else {
			q = append(q, qsim.One())
		}
	}

	// https://arxiv.org/abs/2105.06632
	period := func() {
		e := .5
		qsim.RY(math.Pi-e, q[0])
		qsim.RY(math.Pi-e, q[1])
		qsim.RY(math.Pi-e, q[2])
		qsim.RY(math.Pi-e, q[3])

		qsim.CNOT(q[0], q[1])
		qsim.CNOT(q[2], q[3])

		j := func() float64 {
			return math.Pi * (1.0 + 2*rand.Float64()) / 8
		}
		qsim.RZ(-2*j(), q[1])
		qsim.RZ(-2*j(), q[3])

		qsim.CNOT(q[0], q[1])
		qsim.CNOT(q[2], q[3])
		qsim.CNOT(q[1], q[2])

		qsim.RZ(-2*j(), q[2])

		qsim.CNOT(q[1], q[2])
	}
	sum, last, count := 0.0, -1, 0.0
	for i := 0; i < 128; i++ {
		period()
		max, binary := 0.0, []string{}
		for _, state := range qsim.State() {
			if state.Probability > max {
				max, binary = state.Probability, state.BinaryString
			}
		}
		if binary[0] == init {
			sum += float64(i - last)
			last = i
			count++
		}
		//fmt.Println(binary, max)
	}
	return sum / count
}
