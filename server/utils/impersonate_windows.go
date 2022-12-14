// A demonstration example for http://stackoverflow.com/a/26124494
// It runs a goroutine locked to an OS thread on Windows
// then impersonates that thread as another user using its name
// and plaintext password, then reverts to the default security
// context before detaching from its OS thread.

// +build windows

package utils

import (
	"log"
	"runtime"
	"sync"
	"syscall"
	"unsafe"
)

var (
	advapi32        = syscall.NewLazyDLL("advapi32.dll")
	logonProc       = advapi32.NewProc("LogonUserW")
	impersonateProc = advapi32.NewProc("ImpersonateLoggedOnUser")
	revertSelfProc  = advapi32.NewProc("RevertToSelf")
)

func WindowsLogin(user string, pass string) bool {
	log.SetFlags(0)

	var wg sync.WaitGroup
	wg.Add(1)

	var x bool
	go func(user, pass string) {
		defer wg.Done()
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		log.Println("In a goroutine")
		err := impersonate(user, pass)
		if err != nil {
			x = false
			return
		}
		defer mustRevertToSelf()
		// Here, we're impersonated as $user identified by $pass
		log.Println("Impersonated")
		x = true
	}(user, pass)
	wg.Wait()
	log.Println("OK")

	return x
}

func impersonate(user, pass string) error {
	token, err := logonUser(user, pass)
	if err != nil {
		return err
	}
	defer mustCloseHandle(token)
	return impersonateUser(token)
}

func logonUser(user, pass string) (token syscall.Handle, err error) {
	const (
		// Taken from WinBase.h (SDK 7.1):
		LOGON32_LOGON_NETWORK    = uintptr(3)
		LOGON32_PROVIDER_DEFAULT = uintptr(0)
	)

	// ".\0" meaning "this computer:
	domain,_ := syscall.UTF16FromString("LPU")

	var pu, pp []uint16
	pu, err = syscall.UTF16FromString(user)
	if err != nil {
		return
	}
	pp, err = syscall.UTF16FromString(pass)
	if err != nil {
		return
	}

	rc, _, ec := syscall.Syscall6(logonProc.Addr(), 6,
		uintptr(unsafe.Pointer(&pu[0])),
		uintptr(unsafe.Pointer(&domain[0])),
		uintptr(unsafe.Pointer(&pp[0])),
		LOGON32_LOGON_NETWORK,
		LOGON32_PROVIDER_DEFAULT,
		uintptr(unsafe.Pointer(&token)))
	if rc == 0 {
		err = error(ec)
	}
	return
}

func impersonateUser(token syscall.Handle) error {
	rc, _, ec := syscall.Syscall(impersonateProc.Addr(), 1, uintptr(token), 0, 0)
	if rc == 0 {
		return error(ec)
	}
	return nil
}

func revertToSelf() error {
	rc, _, ec := syscall.Syscall(revertSelfProc.Addr(), 0, 0, 0, 0)
	if rc == 0 {
		return error(ec)
	}
	return nil
}

func mustRevertToSelf() {
	err := revertToSelf()
	if err != nil {
		panic(err)
	}
}

func mustCloseHandle(handle syscall.Handle) {
	err := syscall.CloseHandle(handle)
	if err != nil {
		panic(err)
	}
}