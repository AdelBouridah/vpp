// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file

// Script for automatic code generation of the benchmark routines.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var gopath string

var copyrightnotice = []byte(`// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file`)

var autogen = []byte("// Code generated by \"go run $GOPATH/src/gonum.org/v1/gonum/blas/testblas/benchautogen/autogen_bench_level1double.go\"; DO NOT EDIT.\n")

var imports = []byte(`import(
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/blas"
)`)

var randomSliceFunction = []byte(`func randomSlice(l, idx int) ([]float64) {
	if idx < 0{
		idx = -idx
	}
	s := make([]float64, l * idx)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}`)

const (
	posInc1 = 5
	posInc2 = 3
	negInc1 = -3
	negInc2 = -4
)

var level1Sizes = []struct {
	lower string
	upper string
	camel string
	size  int
}{
	{
		lower: "small",
		upper: "SMALL_SLICE",
		camel: "Small",
		size:  10,
	},
	{
		lower: "medium",
		upper: "MEDIUM_SLICE",
		camel: "Medium",
		size:  1000,
	},
	{
		lower: "large",
		upper: "LARGE_SLICE",
		camel: "Large",
		size:  100000,
	},
	{
		lower: "huge",
		upper: "HUGE_SLICE",
		camel: "Huge",
		size:  10000000,
	},
}

type level1functionStruct struct {
	camel      string
	sig        string
	call       string
	extraSetup string
	oneInput   bool
	extraName  string // if have a couple different cases for the same function
}

var level1Functions = []level1functionStruct{
	{
		camel:    "Ddot",
		sig:      "n int, x []float64, incX int, y []float64, incY int",
		call:     "n, x, incX, y, incY",
		oneInput: false,
	},
	{
		camel:    "Dnrm2",
		sig:      "n int, x []float64, incX int",
		call:     "n, x, incX",
		oneInput: true,
	},
	{
		camel:    "Dasum",
		sig:      "n int, x []float64, incX int",
		call:     "n, x, incX",
		oneInput: true,
	},
	{
		camel:    "Idamax",
		sig:      "n int, x []float64, incX int",
		call:     "n, x, incX",
		oneInput: true,
	},
	{
		camel:    "Dswap",
		sig:      "n int, x []float64, incX int, y []float64, incY int",
		call:     "n, x, incX, y, incY",
		oneInput: false,
	},
	{
		camel:    "Dcopy",
		sig:      "n int, x []float64, incX int, y []float64, incY int",
		call:     "n, x, incX, y, incY",
		oneInput: false,
	},
	{
		camel:      "Daxpy",
		sig:        "n int, alpha float64, x []float64, incX int, y []float64, incY int",
		call:       "n, alpha, x, incX, y, incY",
		extraSetup: "alpha := 2.4",
		oneInput:   false,
	},
	{
		camel:      "Drot",
		sig:        "n int, x []float64, incX int, y []float64, incY int, c, s float64",
		call:       "n, x, incX, y, incY, c, s",
		extraSetup: "c := 0.89725836967\ns:= 0.44150585279",
		oneInput:   false,
	},
	{
		camel:      "Drotm",
		sig:        "n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams",
		call:       "n, x, incX, y, incY, p",
		extraSetup: "p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375,0}}",
		oneInput:   false,
		extraName:  "OffDia",
	},
	{
		camel:      "Drotm",
		sig:        "n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams",
		call:       "n, x, incX, y, incY, p",
		extraSetup: "p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}",
		oneInput:   false,
		extraName:  "Dia",
	},
	{
		camel:      "Drotm",
		sig:        "n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams",
		call:       "n, x, incX, y, incY, p",
		extraSetup: "p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}",
		oneInput:   false,
		extraName:  "Resc",
	},
	{
		camel:      "Dscal",
		sig:        "n int, alpha float64, x []float64, incX int",
		call:       "n, alpha, x, incX",
		extraSetup: "alpha := 2.4",
		oneInput:   true,
	},
}

func init() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		panic("gopath not set")
	}
}

func main() {
	pkgs := []string{"gonum", "netlib"}
	for _, pkg := range pkgs {
		blasPath := filepath.Join(gopath, "src", "gonum.org", "v1", pkg, "blas", pkg)
		err := level1(blasPath, pkg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = exec.Command("goimports", "-w", blasPath).Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func printHeader(f errFile, name string) {
	f.Write(autogen)
	f.WriteString("\n\n")
	f.Write(copyrightnotice)
	f.WriteString("\n\n")
	f.WriteString("package " + name)
	f.WriteString("\n\n")
	f.Write(imports)
	f.WriteString("\n\n")
}

// Generate the benchmark scripts for level1
func level1(benchPath string, pkgname string) error {
	// Generate level 1 benchmarks
	level1Filepath := filepath.Join(benchPath, "level1float64_bench_test.go")
	var f errFile
	f.file, f.err = os.Create(level1Filepath)
	if f.err != nil {
		return f.err
	}
	defer f.file.Close()

	printHeader(f, pkgname)

	// Print all of the constants
	f.WriteString("const (\n")
	f.WriteString("\tposInc1 = " + strconv.Itoa(posInc1) + "\n")
	f.WriteString("\tposInc2 = " + strconv.Itoa(posInc2) + "\n")
	f.WriteString("\tnegInc1 = " + strconv.Itoa(negInc1) + "\n")
	f.WriteString("\tnegInc2 = " + strconv.Itoa(negInc2) + "\n")
	for _, con := range level1Sizes {
		f.WriteString("\t" + con.upper + " = " + strconv.Itoa(con.size) + "\n")
	}
	f.WriteString(")\n")
	f.WriteString("\n")

	// Write the randomSlice function
	f.Write(randomSliceFunction)
	f.WriteString("\n\n")

	// Start writing the benchmarks
	for _, fun := range level1Functions {
		writeLevel1Benchmark(fun, f)
		f.WriteString("\n/* ------------------ */ \n")
	}

	return f.err
}

func writeLevel1Benchmark(fun level1functionStruct, f errFile) {
	// First, write the base benchmark file
	f.WriteString("func benchmark" + fun.camel + fun.extraName + "(b *testing.B, ")
	f.WriteString(fun.sig)
	f.WriteString(") {\n")

	f.WriteString("b.ResetTimer()\n")
	f.WriteString("for i := 0; i < b.N; i++{\n")
	f.WriteString("\timpl." + fun.camel + "(")

	f.WriteString(fun.call)
	f.WriteString(")\n}\n}\n")
	f.WriteString("\n")

	// Write all of the benchmarks to call it
	for _, sz := range level1Sizes {
		lambda := func(incX, incY, name string, twoInput bool) {
			f.WriteString("func Benchmark" + fun.camel + fun.extraName + sz.camel + name + "(b *testing.B){\n")
			f.WriteString("n := " + sz.upper + "\n")
			f.WriteString("incX := " + incX + "\n")
			f.WriteString("x := randomSlice(n, incX)\n")
			if twoInput {
				f.WriteString("incY := " + incY + "\n")
				f.WriteString("y := randomSlice(n, incY)\n")
			}
			f.WriteString(fun.extraSetup + "\n")
			f.WriteString("benchmark" + fun.camel + fun.extraName + "(b, " + fun.call + ")\n")
			f.WriteString("}\n\n")
		}
		if fun.oneInput {
			lambda("1", "", "UnitaryInc", false)
			lambda("posInc1", "", "PosInc", false)
		} else {
			lambda("1", "1", "BothUnitary", true)
			lambda("posInc1", "1", "IncUni", true)
			lambda("1", "negInc1", "UniInc", true)
			lambda("posInc1", "negInc1", "BothInc", true)
		}
	}
}

type errFile struct {
	file *os.File
	err  error
}

func (f errFile) Write(b []byte) {
	if f.err != nil {
		return
	}
	_, f.err = f.file.Write(b)
}

func (f errFile) WriteString(s string) {
	if f.err != nil {
		return
	}
	_, f.err = f.file.WriteString(s)
}
