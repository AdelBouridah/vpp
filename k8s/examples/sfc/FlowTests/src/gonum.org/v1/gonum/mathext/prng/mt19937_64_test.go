// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prng

import (
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

var _ rand.Source = (*MT19937_64)(nil)

// Random values in tests are produced by 40 iterations of the C code
// with or without an initial seed array.

func TestMT19937_64(t *testing.T) {
	want := []uint64{
		14514284786278117030, 4620546740167642908, 13109570281517897720, 17462938647148434322, 355488278567739596,
		7469126240319926998, 4635995468481642529, 418970542659199878, 9604170989252516556, 6358044926049913402,
		5058016125798318033, 10349215569089701407, 2583272014892537200, 10032373690199166667, 9627645531742285868,
		15810285301089087632, 9219209713614924562, 7736011505917826031, 13729552270962724157, 4596340717661012313,
		4413874586873285858, 5904155143473820934, 16795776195466785825, 3040631852046752166, 4529279813148173111,
		3658352497551999605, 13205889818278417278, 17853215078830450730, 14193508720503142180, 1488787817663097441,
		8484116316263611556, 4745643133208116498, 14333959900198994173, 10770733876927207790, 17529942701849009476,
		8081518017574486547, 5945178879512507902, 9821139136195250096, 4728986788662773602, 840062144447779464,
	}

	mt := NewMT19937_64()
	for i := range want {
		got := mt.Uint64()
		if got != want[i] {
			t.Errorf("unexpected random value at iteration %d: got:%d want:%d", i, got, want[i])
		}
	}
}

func TestMT19937_64SeedFromKeys(t *testing.T) {
	want := []uint64{
		7266447313870364031, 4946485549665804864, 16945909448695747420, 16394063075524226720, 4873882236456199058,
		14877448043947020171, 6740343660852211943, 13857871200353263164, 5249110015610582907, 10205081126064480383,
		1235879089597390050, 17320312680810499042, 16489141110565194782, 8942268601720066061, 13520575722002588570,
		14226945236717732373, 9383926873555417063, 15690281668532552105, 11510704754157191257, 15864264574919463609,
		6489677788245343319, 5112602299894754389, 10828930062652518694, 15942305434158995996, 15445717675088218264,
		4764500002345775851, 14673753115101942098, 236502320419669032, 13670483975188204088, 14931360615268175698,
		8904234204977263924, 12836915408046564963, 12120302420213647524, 15755110976537356441, 5405758943702519480,
		10951858968426898805, 17251681303478610375, 4144140664012008120, 18286145806977825275, 13075804672185204371,
	}

	mt := NewMT19937_64()
	mt.SeedFromKeys([]uint64{0x12345, 0x23456, 0x34567, 0x45678})
	for i := range want {
		got := mt.Uint64()
		if got != want[i] {
			t.Errorf("unexpected random value at iteration %d: got:%d want:%d", i, got, want[i])
		}
	}
}

func TestMT19937_64RoundTrip(t *testing.T) {
	var src MT19937_64
	src.Seed(uint64(time.Now().Unix()))

	src.Uint64() // Step PRNG once to makes sure states are mixed.

	buf, err := src.MarshalBinary()
	if err != nil {
		t.Errorf("unexpected error marshaling state: %v", err)
	}

	var dst MT19937_64
	// Get dst into a non-zero state.
	dst.Seed(1)
	for i := 0; i < 10; i++ {
		dst.Uint64()
	}

	err = dst.UnmarshalBinary(buf)
	if err != nil {
		t.Errorf("unexpected error unmarshaling state: %v", err)
	}

	if dst != src {
		t.Errorf("mismatch between generator states: got:%+v want:%+v", dst, src)
	}
}