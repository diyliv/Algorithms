package main

import (
	"unsafe"

	"github.com/elliotchance/c2go/noarch"
)

type size_t uint32
type wchar_t int32

const P_ALL int32 = 0
const P_PID int32 = 1
const P_PGID int32 = 2

type idtype_t int32
type _Float32 float32
type _Float64 float64
type _Float32x float64
type _Float64x float64
type div_t struct {
	quot int32
	rem  int32
}
type ldiv_t struct {
	quot int32
	rem  int32
}
type lldiv_t struct {
	quot int64
	rem  int64
}
type __locale_t int32
type locale_t __locale_t
type __u_char uint8
type __u_short uint16
type __u_int uint32
type __u_long uint32
type __int8_t int8
type __uint8_t uint8
type __int16_t int16
type __uint16_t uint16
type __int32_t int32
type __uint32_t uint32
type __int64_t int32
type __uint64_t uint32
type __quad_t int32
type __u_quad_t uint32
type __intmax_t int32
type __uintmax_t uint32
type __dev_t uint32
type __uid_t uint32
type __gid_t uint32
type __ino_t uint32
type __ino64_t uint32
type __mode_t uint32
type __nlink_t uint32
type __off_t int32
type __off64_t int32
type __pid_t int32
type __fsid_t struct {
	__val [2]int32
}
type __clock_t int32
type __rlim_t uint32
type __rlim64_t uint32
type __id_t uint32
type __time_t int32
type __useconds_t uint32
type __suseconds_t int32
type __daddr_t int32
type __key_t int32
type __clockid_t int32
type __timer_t unsafe.Pointer
type __blksize_t int32
type __blkcnt_t int32
type __blkcnt64_t int32
type __fsblkcnt_t uint32
type __fsblkcnt64_t uint32
type __fsfilcnt_t uint32
type __fsfilcnt64_t uint32
type __fsword_t int32
type __ssize_t int32
type __syscall_slong_t int32
type __syscall_ulong_t uint32
type __loff_t __off64_t
type __caddr_t *byte
type __intptr_t int32
type __socklen_t uint32
type __sig_atomic_t int32
type u_char __u_char
type u_short __u_short
type u_int __u_int
type u_long __u_long
type quad_t __quad_t
type u_quad_t __u_quad_t
type fsid_t __fsid_t
type loff_t __loff_t
type ino_t __ino_t
type ino64_t __ino64_t
type dev_t __dev_t
type gid_t __gid_t
type mode_t __mode_t
type nlink_t __nlink_t
type uid_t __uid_t
type off_t __off_t
type off64_t __off64_t
type pid_t __pid_t
type id_t __id_t
type ssize_t __ssize_t
type daddr_t __daddr_t
type caddr_t __caddr_t
type key_t __key_t
type clock_t __clock_t
type clockid_t __clockid_t
type time_t __time_t
type timer_t __timer_t
type useconds_t __useconds_t
type suseconds_t __suseconds_t
type ulong uint32
type ushort uint16
type uint uint32
type int8_t int8
type int16_t int16
type int32_t int32
type int64_t int64
type u_int8_t uint8
type u_int16_t uint16
type u_int32_t uint32
type u_int64_t uint32
type register_t int32

// __uint16_identity - transpiled function from  /usr/include/x86_64-linux-gnu/bits/uintn-identity.h:32
func __uint16_identity(__x uint16) uint16 {
	return uint16(__x)
}

// __uint32_identity - transpiled function from  /usr/include/x86_64-linux-gnu/bits/uintn-identity.h:38
func __uint32_identity(__x uint32) uint32 {
	return uint32(__x)
}

// __uint64_identity - transpiled function from  /usr/include/x86_64-linux-gnu/bits/uintn-identity.h:44
func __uint64_identity(__x uint64) uint64 {
	return uint64(__x)
}

type __sigset_t struct {
	__val [16]uint32
}
type sigset_t __sigset_t
type timeval struct {
	tv_sec  __time_t
	tv_usec __suseconds_t
}
type timespec struct {
	tv_sec  __time_t
	tv_nsec __syscall_slong_t
}
type __fd_mask int32
type fd_set struct {
	fds_bits [16]__fd_mask
}
type fd_mask __fd_mask
type blksize_t __blksize_t
type blkcnt_t __blkcnt_t
type fsblkcnt_t __fsblkcnt_t
type fsfilcnt_t __fsfilcnt_t
type blkcnt64_t __blkcnt64_t
type fsblkcnt64_t __fsblkcnt64_t
type fsfilcnt64_t __fsfilcnt64_t
type __pthread_rwlock_arch_t struct {
	__readers       uint32
	__writers       uint32
	__wrphase_futex uint32
	__writers_futex uint32
	__pad3          uint32
	__pad4          uint32
	__cur_writer    int32
	__shared        int32
	__rwelision     int8
	__pad1          [7]uint8
	__pad2          uint32
	__flags         uint32
}
type __pthread_internal_list struct {
	__prev *__pthread_internal_list
	__next *__pthread_internal_list
}
type __pthread_list_t __pthread_internal_list
type __pthread_mutex_s struct {
	__lock    int32
	__count   uint32
	__owner   int32
	__nusers  uint32
	__kind    int32
	__spins   int16
	__elision int16
	__list    __pthread_list_t
}
type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD176D5E struct {
	__low  uint32
	__high uint32
}
type __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E struct{ memory unsafe.Pointer }

func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E) copy() __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E {
	var buffer [16]byte
	for i := range buffer {
		buffer[i] = (*((*[16]byte)(unionVar.memory)))[i]
	}
	var newUnion __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E) __wseq() *uint64 {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*uint64)(unionVar.memory)
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E) __wseq32() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD176D5E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD176D5E)(unionVar.memory)
}

type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD185D5E struct {
	__low  uint32
	__high uint32
}
type __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E struct{ memory unsafe.Pointer }

func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E) copy() __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E {
	var buffer [16]byte
	for i := range buffer {
		buffer[i] = (*((*[16]byte)(unionVar.memory)))[i]
	}
	var newUnion __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E) __g1_start() *uint64 {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*uint64)(unionVar.memory)
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E) __g1_start32() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD185D5E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD185D5E)(unionVar.memory)
}

type __pthread_cond_s struct {
	__g_refs       [2]uint32
	__g_size       [2]uint32
	__g1_orig_size uint32
	__wrefs        uint32
	__g_signals    [2]uint32
}
type pthread_t uint32
type pthread_mutexattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_mutexattr_t) copy() pthread_mutexattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_mutexattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_mutexattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_mutexattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_condattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_condattr_t) copy() pthread_condattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_condattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_condattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_condattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_key_t uint32
type pthread_once_t int32
type pthread_attr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_attr_t) copy() pthread_attr_t {
	var buffer [56]byte
	for i := range buffer {
		buffer[i] = (*((*[56]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_attr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_attr_t) __size() *[56]byte {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[56]byte)(unionVar.memory)
}
func (unionVar *pthread_attr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_mutex_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_mutex_t) copy() pthread_mutex_t {
	var buffer [40]byte
	for i := range buffer {
		buffer[i] = (*((*[40]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_mutex_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_mutex_t) __data() *__pthread_mutex_s {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__pthread_mutex_s)(unionVar.memory)
}
func (unionVar *pthread_mutex_t) __size() *[40]byte {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[40]byte)(unionVar.memory)
}
func (unionVar *pthread_mutex_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_rwlock_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_rwlock_t) copy() pthread_rwlock_t {
	var buffer [56]byte
	for i := range buffer {
		buffer[i] = (*((*[56]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_rwlock_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_rwlock_t) __data() *__pthread_rwlock_arch_t {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__pthread_rwlock_arch_t)(unionVar.memory)
}
func (unionVar *pthread_rwlock_t) __size() *[56]byte {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[56]byte)(unionVar.memory)
}
func (unionVar *pthread_rwlock_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_rwlockattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_rwlockattr_t) copy() pthread_rwlockattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_rwlockattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_rwlockattr_t) __size() *[8]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[8]byte)(unionVar.memory)
}
func (unionVar *pthread_rwlockattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_spinlock_t int32
type pthread_barrier_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_barrier_t) copy() pthread_barrier_t {
	var buffer [32]byte
	for i := range buffer {
		buffer[i] = (*((*[32]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_barrier_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_barrier_t) __size() *[32]byte {
	if unionVar.memory == nil {
		var buffer [32]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[32]byte)(unionVar.memory)
}
func (unionVar *pthread_barrier_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [32]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_barrierattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_barrierattr_t) copy() pthread_barrierattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_barrierattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_barrierattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_barrierattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type random_data struct {
	fptr      *int32
	rptr      *int32
	state     *int32
	rand_type int32
	rand_deg  int32
	rand_sep  int32
	end_ptr   *int32
}
type drand48_data struct {
	__x     [3]uint16
	__old_x [3]uint16
	__c     uint16
	__init  uint16
	__a     uint64
}
type __compar_fn_t func(unsafe.Pointer, unsafe.Pointer) int32
type comparison_fn_t __compar_fn_t
type __compar_d_fn_t func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32
type ptrdiff_t int32
type max_align_t struct {
	__clang_max_align_nonce1 int64
	__clang_max_align_nonce2 float64
}
type Data unsafe.Pointer
type FFree func(unsafe.Pointer)
type Array struct {
}

// array_create - transpiled function from  /home/lawliet/go/src/github.com/diyliv/lab/array.c:46
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:173 : Error : name of FieldDecl is empty
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:175 : could not parse &{93839361008136 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 175 0 42 0 } col:42 true __wseq unsigned long long [0xc0001be180 0xc0001be200]}
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:180 : could not parse &{93839361008216 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 180 0 7 0 } col:7 true __wseq32 struct (anonymous struct at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:176:5)':'struct __pthread_cond_s::(anonymous at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:176:5) [0xc000026580 0xc000026600]}
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:182 : Error : name of FieldDecl is empty
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:184 : could not parse &{93839361009240 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 184 0 42 0 } col:42 true __g1_start unsigned long long [0xc0005ae200 0xc0005ae280]}
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:189 : could not parse &{93839361009320 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 189 0 7 0 } col:7 true __g1_start32 struct (anonymous struct at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:185:5)':'struct __pthread_cond_s::(anonymous at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:185:5) [0xc0005ae300 0xc0005ae380]}
// Warning (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/pthreadtypes.h:75 : could not determine the size of type `union pthread_cond_t` for that reason: Cannot determine sizeof : |union pthread_cond_t|. err = Cannot determine sizeof : |struct __pthread_cond_s|. err = Cannot determine sizeof : |__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E|. err = error in array size
// Error (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/pthreadtypes.h:75 : Cannot determine sizeof : |union pthread_cond_t|. err = Cannot determine sizeof : |struct __pthread_cond_s|. err = Cannot determine sizeof : |__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E|. err = error in array size
// Non-resizeable array
// Stores pointer to custom user data
// Custom function to free user pointers on delete
// create array
// delete array, free memory
// returns specified array element
// sets the specified array element to the value
// returns array size
// create array
//
func array_create(size uint32, f FFree) *Array {
	return &Array{}
}

// array_delete - transpiled function from  /home/lawliet/go/src/github.com/diyliv/lab/array.c:52
// delete array, free memory
//
func array_delete(arr *Array) {
	noarch.Free(unsafe.Pointer(arr))
}

// array_get - transpiled function from  /home/lawliet/go/src/github.com/diyliv/lab/array.c:58
// returns specified array element
//
func array_get(arr *Array, index uint32) Data {
	return Data((nil))
}

// array_set - transpiled function from  /home/lawliet/go/src/github.com/diyliv/lab/array.c:64
// sets the specified array element to the value
//
func array_set(arr *Array, index uint32, value Data) {
}

// array_size - transpiled function from  /home/lawliet/go/src/github.com/diyliv/lab/array.c:69
// returns array size
//
func array_size(arr *Array) uint32 {
	return uint32(int32(0))
}
