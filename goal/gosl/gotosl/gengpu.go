// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gotosl

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

// genSysName is the name to use for system in generating code.
// if only one system, the name is empty
func (st *State) genSysName(sy *System) string {
	if len(st.Systems) == 1 {
		return ""
	}
	return sy.Name
}

// genSysVar is the name to use for system in generating code.
// if only one system, the name is empty
func (st *State) genSysVar(sy *System) string {
	return fmt.Sprintf("GPU%sSystem", st.genSysName(sy))
}

// GenGPU generates and writes the Go GPU helper code
func (st *State) GenGPU() {
	var b strings.Builder

	header := `// Code generated by "gosl"; DO NOT EDIT

package %s

import (
	"embed"
	"unsafe"
	"cogentcore.org/core/gpu"
)

//go:embed %s/*.wgsl
var shaders embed.FS

// ComputeGPU is the compute gpu device
var ComputeGPU *gpu.GPU

// UseGPU indicates whether to use GPU vs. CPU.
var UseGPU bool

`

	b.WriteString(fmt.Sprintf(header, st.Package, st.Config.Output))

	sys := maps.Keys(st.Systems)
	slices.Sort(sys)

	for _, synm := range sys {
		sy := st.Systems[synm]
		b.WriteString(fmt.Sprintf("// %s is a GPU compute System with kernels operating on the\n// same set of data variables.\n", st.genSysVar(sy)))
		b.WriteString(fmt.Sprintf("var %s *gpu.ComputeSystem\n", st.genSysVar(sy)))
	}

	venum := `
// GPUVars is an enum for GPU variables, for specifying what to sync.
type GPUVars int32 //enums:enum

const (
`

	b.WriteString(venum)

	vidx := 0
	for _, synm := range sys {
		sy := st.Systems[synm]
		for _, gp := range sy.Groups {
			for _, vr := range gp.Vars {
				b.WriteString(fmt.Sprintf("\t%sVar GPUVars = %d\n", vr.Name, vidx))
				vidx++
			}
		}
	}
	b.WriteString(")\n")

	initf := `
// GPUInit initializes the GPU compute system,
// configuring system(s), variables and kernels.
// It is safe to call multiple times: detects if already run.
func GPUInit() {
	if ComputeGPU != nil {
		return
	}
	gp := gpu.NewComputeGPU()
	ComputeGPU = gp
`

	b.WriteString(initf)

	for _, synm := range sys {
		sy := st.Systems[synm]
		b.WriteString(st.GenGPUSystemInit(sy))
	}
	b.WriteString("}\n\n")

	release := `// GPURelease releases the GPU compute system resources.
// Call this at program exit.
func GPURelease() {
`

	b.WriteString(release)

	sysRelease := `	if %[1]s != nil {
		%[1]s.Release()
		%[1]s = nil
	}
`

	for _, synm := range sys {
		sy := st.Systems[synm]
		b.WriteString(fmt.Sprintf(sysRelease, st.genSysVar(sy)))
	}

	gpuRelease := `
	if ComputeGPU != nil {
		ComputeGPU.Release()
		ComputeGPU = nil
	}
}

`
	b.WriteString(gpuRelease)

	for _, synm := range sys {
		sy := st.Systems[synm]
		b.WriteString(st.GenGPUSystemOps(sy))
	}

	gs := b.String()
	fn := "gosl.go"
	os.WriteFile(fn, []byte(gs), 0644)
}

// GenGPUSystemInit generates GPU Init code for given system.
func (st *State) GenGPUSystemInit(sy *System) string {
	var b strings.Builder

	syvar := st.genSysVar(sy)

	b.WriteString("\t{\n")
	b.WriteString(fmt.Sprintf("\t\tsy := gpu.NewComputeSystem(gp, %q)\n", sy.Name))
	b.WriteString(fmt.Sprintf("\t\t%s = sy\n", syvar))

	kns := maps.Keys(sy.Kernels)
	slices.Sort(kns)
	for _, knm := range kns {
		kn := sy.Kernels[knm]
		b.WriteString(fmt.Sprintf("\t\tgpu.NewComputePipelineShaderFS(shaders, %q, sy)\n", kn.Filename))
	}
	b.WriteString("\t\tvars := sy.Vars()\n")
	for _, gp := range sy.Groups {
		b.WriteString("\t\t{\n")
		gtyp := "gpu.Storage"
		if gp.Uniform {
			gtyp = "gpu.Uniform"
		}
		b.WriteString(fmt.Sprintf("\t\t\tsgp := vars.AddGroup(%s)\n", gtyp))
		b.WriteString("\t\t\tvar vr *gpu.Var\n\t\t\t_ = vr\n")
		for _, vr := range gp.Vars {
			if vr.Tensor {
				typ := strings.TrimPrefix(vr.Type, "tensor.")
				b.WriteString(fmt.Sprintf("\t\t\tvr = sgp.Add(%q, gpu.%s, 1, gpu.ComputeShader)\n", vr.Name, typ))
			} else {
				b.WriteString(fmt.Sprintf("\t\t\tvr = sgp.AddStruct(%q, int(unsafe.Sizeof(%s{})), 1, gpu.ComputeShader)\n", vr.Name, vr.SLType()))
			}
			if vr.ReadOnly {
				b.WriteString("\t\t\tvr.ReadOnly = true\n")
			}
		}
		b.WriteString("\t\t\tsgp.SetNValues(1)\n")
		b.WriteString("\t\t}\n")
	}
	b.WriteString("\t\tsy.Config()\n")
	b.WriteString("\t}\n")
	return b.String()
}

// GenGPUSystemOps generates GPU helper functions for given system.
func (st *State) GenGPUSystemOps(sy *System) string {
	var b strings.Builder

	syvar := st.genSysVar(sy)
	synm := st.genSysName(sy)

	// 1 = kernel, 2 = system var, 3 = sysname (blank for 1 default)
	run := `// Run%[1]s runs the %[1]s kernel with given number of elements,
// on either the CPU or GPU depending on the UseGPU variable.
// Can call multiple Run* kernels in a row, which are then all launched
// in the same command submission on the GPU, which is by far the most efficient.
// MUST call RunDone (with optional vars to sync) after all Run calls.
// Alternatively, a single-shot RunOne%[1]s call does Run and Done for a
// single run-and-sync case.
func Run%[1]s(n int) {
	if UseGPU {
		Run%[1]sGPU(n)
	} else {
		Run%[1]sCPU(n)
	}
}

// Run%[1]sGPU runs the %[1]s kernel on the GPU. See [Run%[1]s] for more info.
func Run%[1]sGPU(n int) {
	sy := %[2]s
	pl := sy.ComputePipelines[%[1]q]
	ce, _ := sy.BeginComputePass()
	pl.Dispatch1D(ce, n, 64)
}

// Run%[1]sCPU runs the %[1]s kernel on the CPU.
func Run%[1]sCPU(n int) {
	gpu.VectorizeFunc(0, n, %[1]s)
}

// RunOne%[1]s runs the %[1]s kernel with given number of elements,
// on either the CPU or GPU depending on the UseGPU variable.
// This version then calls RunDone with the given variables to sync
// after the Run, for a single-shot Run-and-Done call. If multiple kernels
// can be run in sequence, it is much more efficient to do multiple Run*
// calls followed by a RunDone call.
func RunOne%[1]s(n int, syncVars ...GPUVars) {
	if UseGPU {
		Run%[1]sGPU(n)
		RunDone%[3]s(syncVars...)
	} else {
		Run%[1]sCPU(n)
	}
}
`
	// 1 = sysname (blank for 1 default), 2 = system var
	runDone := `// RunDone%[1]s must be called after Run* calls to start compute kernels.
// This actually submits the kernel jobs to the GPU, and adds commands
// to synchronize the given variables back from the GPU to the CPU.
// After this function completes, the GPU results will be available in 
// the specified variables.
func RunDone%[1]s(syncVars ...GPUVars) {
	if !UseGPU {
		return
	}
	sy := %[2]s
	sy.ComputeEncoder.End()
	%[1]sReadFromGPU(syncVars...)
	sy.EndComputePass()
	%[1]sSyncFromGPU(syncVars...)
}

// %[1]sToGPU copies given variables to the GPU for the system.
func %[1]sToGPU(vars ...GPUVars) {
	if !UseGPU {
		return
	}
	sy := %[2]s
	syVars := sy.Vars()
	for _, vr := range vars {
		switch vr {
`

	kns := maps.Keys(sy.Kernels)
	slices.Sort(kns)
	for _, knm := range kns {
		kn := sy.Kernels[knm]
		b.WriteString(fmt.Sprintf(run, kn.Name, syvar, synm))
	}
	b.WriteString(fmt.Sprintf(runDone, synm, syvar))

	for gi, gp := range sy.Groups {
		for _, vr := range gp.Vars {
			b.WriteString(fmt.Sprintf("\t\tcase %sVar:\n", vr.Name))
			b.WriteString(fmt.Sprintf("\t\t\tv, _ := syVars.ValueByIndex(%d, %q, 0)\n", gi, vr.Name))
			vv := vr.Name
			if vr.Tensor {
				vv += ".Values"
			}
			b.WriteString(fmt.Sprintf("\t\t\tgpu.SetValueFrom(v, %s)\n", vv))
		}
	}
	b.WriteString("\t\t}\n\t}\n}\n")

	fmGPU := `
// %[1]sReadFromGPU starts the process of copying vars to the GPU.
func %[1]sReadFromGPU(vars ...GPUVars) {
	sy := %[2]s
	syVars := sy.Vars()
	for _, vr := range vars {
		switch vr {
`

	b.WriteString(fmt.Sprintf(fmGPU, synm, syvar))

	for gi, gp := range sy.Groups {
		for _, vr := range gp.Vars {
			b.WriteString(fmt.Sprintf("\t\tcase %sVar:\n", vr.Name))
			b.WriteString(fmt.Sprintf("\t\t\tv, _ := syVars.ValueByIndex(%d, %q, 0)\n", gi, vr.Name))
			b.WriteString("\t\t\tv.GPUToRead(sy.CommandEncoder)\n")
		}
	}
	b.WriteString("\t\t}\n\t}\n}\n")

	syncGPU := `
// %[1]sSyncFromGPU synchronizes vars from the GPU to the actual variable.
func %[1]sSyncFromGPU(vars ...GPUVars) {
	sy := %[2]s
	syVars := sy.Vars()
	for _, vr := range vars {
		switch vr {
`

	b.WriteString(fmt.Sprintf(syncGPU, synm, syvar))

	for gi, gp := range sy.Groups {
		for _, vr := range gp.Vars {
			b.WriteString(fmt.Sprintf("\t\tcase %sVar:\n", vr.Name))
			b.WriteString(fmt.Sprintf("\t\t\tv, _ := syVars.ValueByIndex(%d, %q, 0)\n", gi, vr.Name))
			b.WriteString(fmt.Sprintf("\t\t\tv.ReadSync()\n"))
			vv := vr.Name
			if vr.Tensor {
				vv += ".Values"
			}
			b.WriteString(fmt.Sprintf("\t\t\tgpu.ReadToBytes(v, %s)\n", vv))
		}
	}
	b.WriteString("\t\t}\n\t}\n}\n")

	getFun := `
// Get%[1]s returns a pointer to the given global variable: 
// [%[1]s] []%[2]s at given index.
// To ensure that values are updated on the GPU, you must call [Set%[1]s].
// after all changes have been made.
func Get%[1]s(idx uint32) *%[2]s {
	return &%[1]s[idx]
}
`
	for _, gp := range sy.Groups {
		for _, vr := range gp.Vars {
			if vr.Tensor {
				continue
			}
			b.WriteString(fmt.Sprintf(getFun, vr.Name, vr.SLType()))
		}
	}

	return b.String()
}
