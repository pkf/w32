// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ole32

import (
    "syscall"
    . "w32"
)

var (
    lib uintptr

    procCoInitializeEx uintptr
    procCoInitialize   uintptr
    procCoUninitialize uintptr
)

func init() {
    lib = LoadLib("ole32.dll")

    procCoInitializeEx = GetProcAddr(lib, "CoInitializeEx")
    procCoInitialize = GetProcAddr(lib, "CoInitialize")
    procCoUninitialize = GetProcAddr(lib, "CoUninitialize")
}

func CoInitializeEx(coInit uintptr) HRESULT {
    ret, _, _ := syscall.Syscall(procCoInitializeEx, 2,
        0,
        coInit,
        0)

    switch uint32(ret) {
    case E_INVALIDARG:
        panic("CoInitializeEx failed with E_INVALIDARG")
    case E_OUTOFMEMORY:
        panic("CoInitializeEx failed with E_OUTOFMEMORY")
    case E_UNEXPECTED:
        panic("CoInitializeEx failed with E_UNEXPECTED")
    }

    return HRESULT(ret)
}

func CoInitialize() {
    syscall.Syscall(procCoInitialize, 1,
        0,
        0,
        0)
}

func CoUninitialize() {
    syscall.Syscall(procCoUninitialize, 0,
        0,
        0,
        0)
}
