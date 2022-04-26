// Copyright (c) 2022, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is initially adapted from https://github.com/vulkan-go/asche
// Copyright © 2017 Maxim Kupriianov <max@kc.vc>, under the MIT License

package vgpu

import (
	"log"
	"unsafe"

	vk "github.com/vulkan-go/vulkan"
)

// MemSizeAlign returns the size aligned according to align byte increments
// e.g., if align = 16 and size = 12, it returns 16
func MemSizeAlign(size, align int) int {
	if size%align == 0 {
		return size
	}
	nb := size / align
	return (nb + 1) * align
}

// MemReg is a region of memory
type MemReg struct {
	Offset int
	Size   int
}

// Memory manages memory for the GPU, using separate buffers for
// Images (Textures) vs. other values.
type Memory struct {
	GPU         *GPU
	Device      Device          `desc:"logical device that this memory is managed for: a Surface or GPU itself"`
	CmdPool     CmdPool         `desc:"command pool for memory transfers"`
	Vals        Vals            `desc:"values of Vars, each with a unique name -- can be any number of different values per same Var (e.g., different meshes with vertex data) -- up to user code to bind each Var prior to pipeline execution.  Each of these Vals is mapped into GPU memory  This is only for non-Image objects."`
	Images      Vals            `desc:"Image-type values"`
	BuffSize    int             `desc:"allocated buffer size"`
	BuffHost    vk.Buffer       `view:"-" desc:"logical descriptor for host CPU-visible memory, for staging"`
	BuffHostMem vk.DeviceMemory `view:"-" desc:"host CPU-visible memory, for staging"`
	BuffDev     vk.Buffer       `view:"-" desc:"logical descriptor for device GPU-local memory, for computation"`
	BuffDevMem  vk.DeviceMemory `view:"-" desc:"device GPU-local memory, for computation"`

	Active bool `inactive:"+" desc:"device memory is allocated and tranferred -- ready for use"`
}

// Init configures the Memory for use with given gpu, device, and associated queueindex
func (mm *Memory) Init(gp *GPU, device *Device) {
	mm.GPU = gp
	mm.Device = *device
	mm.CmdPool.Init(device, vk.CommandPoolCreateTransientBit)
}

func (mm *Memory) Destroy(dev vk.Device) {
	mm.Free()
	mm.CmdPool.Destroy(dev)
	mm.GPU = nil
}

// Config should be called after all Vals have been configured
// and are ready to go with their initial data.
// Does: Alloc(), AllocDev(), CopyToStaging(), TransferAllToGPU()
func (mm *Memory) Config() {
	mm.Alloc()
	mm.AllocDev()
	mm.TransferAllToGPU()
	mm.Active = true
}

// Alloc allocates memory for all Vars and Images
func (mm *Memory) Alloc() {
	bsz := mm.Vals.MemSize()
	if bsz != mm.BuffSize {

		usage := vk.BufferUsageVertexBufferBit | vk.BufferUsageIndexBufferBit | vk.BufferUsageUniformBufferBit | vk.BufferUsageStorageBufferBit | vk.BufferUsageUniformTexelBufferBit | vk.BufferUsageStorageTexelBufferBit

		xfer := vk.BufferUsageTransferSrcBit | vk.BufferUsageTransferDstBit

		mm.BuffHost = mm.MakeBuffer(bsz, xfer)
		mm.BuffDev = mm.MakeBuffer(bsz, xfer|usage)
		mm.BuffHostMem = mm.AllocMem(mm.BuffHost, vk.MemoryPropertyHostVisibleBit|vk.MemoryPropertyHostCoherentBit)
		mm.BuffSize = bsz
	}
	var buffPtr unsafe.Pointer
	ret := vk.MapMemory(mm.Device.Device, mm.BuffHostMem, 0, vk.DeviceSize(mm.BuffSize), 0, &buffPtr)
	if IsError(ret) {
		log.Printf("vulkan Memory:CopyBuffs warning: failed to map device memory for data (len=%d)", mm.BuffSize)
		return
	}
	align := int(mm.GPU.GpuProps.Limits.MinUniformBufferOffsetAlignment)
	mm.Vals.Alloc(buffPtr, 0, align)
}

// AllocDev allocates memory on the device
func (mm *Memory) AllocDev() {
	mm.BuffDevMem = mm.AllocMem(mm.BuffDev, vk.MemoryPropertyDeviceLocalBit)
}

// MakeBuffer makes a buffer of given size, usage
func (mm *Memory) MakeBuffer(size int, usage vk.BufferUsageFlagBits) vk.Buffer {
	var buffer vk.Buffer
	ret := vk.CreateBuffer(mm.Device.Device, &vk.BufferCreateInfo{
		SType: vk.StructureTypeBufferCreateInfo,
		Usage: vk.BufferUsageFlags(usage),
		Size:  vk.DeviceSize(size),
	}, nil, &buffer)
	IfPanic(NewError(ret))
	return buffer
}

// AllocMem allocates memory for given buffer, with given properties
func (mm *Memory) AllocMem(buffer vk.Buffer, props vk.MemoryPropertyFlagBits) vk.DeviceMemory {
	// Ask device about its memory requirements.
	var memReqs vk.MemoryRequirements
	vk.GetBufferMemoryRequirements(mm.Device.Device, buffer, &memReqs)
	memReqs.Deref()

	memProps := mm.GPU.MemoryProps
	memType, ok := FindRequiredMemoryType(memProps, vk.MemoryPropertyFlagBits(memReqs.MemoryTypeBits), props)
	if !ok {
		log.Println("vulkan warning: failed to find required memory type")
	}

	var memory vk.DeviceMemory
	// Allocate device memory and bind to the buffer.
	ret := vk.AllocateMemory(mm.Device.Device, &vk.MemoryAllocateInfo{
		SType:           vk.StructureTypeMemoryAllocateInfo,
		AllocationSize:  memReqs.Size,
		MemoryTypeIndex: memType,
	}, nil, &memory)
	IfPanic(NewError(ret))
	vk.BindBufferMemory(mm.Device.Device, buffer, memory, 0)
	return memory
}

// FreeBuffMem frees given device memory to nil
func (mm *Memory) FreeBuffMem(memory *vk.DeviceMemory) {
	if *memory == nil {
		return
	}
	vk.FreeMemory(mm.Device.Device, *memory, nil)
	*memory = nil
}

// Free frees any allocated memory -- returns true if freed
func (mm *Memory) Free() bool {
	if mm.BuffSize == 0 {
		return false
	}
	vk.UnmapMemory(mm.Device.Device, mm.BuffHostMem)
	mm.Vals.Free()
	mm.FreeBuffMem(&mm.BuffDevMem)
	vk.DestroyBuffer(mm.Device.Device, mm.BuffDev, nil)
	mm.FreeBuffMem(&mm.BuffHostMem)
	vk.DestroyBuffer(mm.Device.Device, mm.BuffHost, nil)
	mm.BuffSize = 0
	mm.BuffHost = nil
	mm.BuffDev = nil
	mm.Active = false
	return true
}

// Deactivate deactivates device memory
func (mm *Memory) Deactivate() {
	mm.FreeBuffMem(&mm.BuffDevMem)
	mm.Active = false
}

// Activate ensures device memory is ready to use
// assumes the staging memory is configured.
// Call Sync after this if needed.
func (mm *Memory) Activate() {
	if mm.Active {
		return
	}
	if mm.BuffDevMem == nil {
		mm.AllocDev()
		mm.TransferAllToGPU()
	}
	mm.Active = true
}

// SyncAllToGPU syncs all modified Val regions from CPU to GPU device memory
func (mm *Memory) SyncAllToGPU() {
	mods := mm.Vals.ModRegs()
	if len(mods) == 0 {
		return
	}
	mm.TransferBuffToGPU(mods)
}

// SyncVarsFmGPU syncs given variables from GPU device memory to CPU
func (mm *Memory) SyncVarsFmGPU(vals ...string) {
	nv := len(vals)
	mods := make([]MemReg, nv)
	for i, vnm := range vals {
		vl, err := mm.Vals.ValByNameTry(vnm)
		if err != nil {
			log.Println(err)
			continue
		}
		mods[i] = vl.MemReg()
	}
	mm.TransferBuffFmGPU(mods)
}

// TransferAllToGPU transfers all staging to GPU
func (mm *Memory) TransferAllToGPU() {
	mm.TransferBuffAllToGPU()
}

// TransferBuffAllToGPU transfers entire staging buffer of memory from CPU to GPU
func (mm *Memory) TransferBuffAllToGPU() {
	if mm.BuffSize == 0 || mm.BuffDevMem == nil {
		return
	}
	mm.TransferBuffToGPU([]MemReg{{Offset: 0, Size: mm.BuffSize}})
}

// TransferBuffToGPU transfers buff memory from CPU to GPU for given regs
func (mm *Memory) TransferBuffToGPU(regs []MemReg) {
	if mm.BuffSize == 0 || mm.BuffDevMem == nil {
		return
	}

	cmdBuff := mm.CmdPool.MakeBuff(&mm.Device)
	mm.CmdPool.BeginCmdOneTime()

	rg := make([]vk.BufferCopy, len(regs))
	for i, mr := range regs {
		rg[i] = vk.BufferCopy{SrcOffset: vk.DeviceSize(mr.Offset), DstOffset: vk.DeviceSize(mr.Offset), Size: vk.DeviceSize(mr.Size)}
	}

	vk.CmdCopyBuffer(cmdBuff, mm.BuffHost, mm.BuffDev, uint32(len(rg)), rg)

	mm.CmdPool.SubmitWaitFree(&mm.Device)
}

// TransferBuffFmGPU transfers buff memory from GPU to CPU for given regs
func (mm *Memory) TransferBuffFmGPU(regs []MemReg) {
	if mm.BuffSize == 0 || mm.BuffDevMem == nil {
		return
	}

	cmdBuff := mm.CmdPool.MakeBuff(&mm.Device)
	mm.CmdPool.BeginCmdOneTime()

	rg := make([]vk.BufferCopy, len(regs))
	for i, mr := range regs {
		rg[i] = vk.BufferCopy{SrcOffset: vk.DeviceSize(mr.Offset), DstOffset: vk.DeviceSize(mr.Offset), Size: vk.DeviceSize(mr.Size)}
	}

	vk.CmdCopyBuffer(cmdBuff, mm.BuffDev, mm.BuffHost, uint32(len(rg)), rg)

	mm.CmdPool.SubmitWaitFree(&mm.Device)
}

func FindRequiredMemoryType(props vk.PhysicalDeviceMemoryProperties,
	deviceRequirements, hostRequirements vk.MemoryPropertyFlagBits) (uint32, bool) {

	for i := uint32(0); i < vk.MaxMemoryTypes; i++ {
		if deviceRequirements&(vk.MemoryPropertyFlagBits(1)<<i) != 0 {
			props.MemoryTypes[i].Deref()
			flags := props.MemoryTypes[i].PropertyFlags
			if flags&vk.MemoryPropertyFlags(hostRequirements) != 0 {
				return i, true
			}
		}
	}
	return 0, false
}

func FindRequiredMemoryTypeFallback(props vk.PhysicalDeviceMemoryProperties,
	deviceRequirements, hostRequirements vk.MemoryPropertyFlagBits) (uint32, bool) {

	for i := uint32(0); i < vk.MaxMemoryTypes; i++ {
		if deviceRequirements&(vk.MemoryPropertyFlagBits(1)<<i) != 0 {
			props.MemoryTypes[i].Deref()
			flags := props.MemoryTypes[i].PropertyFlags
			if flags&vk.MemoryPropertyFlags(hostRequirements) != 0 {
				return i, true
			}
		}
	}
	// Fallback to the first one available.
	if hostRequirements != 0 {
		return FindRequiredMemoryType(props, deviceRequirements, 0)
	}
	return 0, false
}
