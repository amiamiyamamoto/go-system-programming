// go言語本体のコードから一部を抜き出したもの

package filelock

import (
	"os"
	"syscall"
)

type lockType int16

const (
	readLock  lockType = syscall.LOCK_SH
	writeLock lockType = syscall.LOCK_EX
)

func lock(f os.File, lt lockType) (err error) {
	for {
		err = syscall.Flock(int(f.Fd()), int(lt))
		if err != syscall.EINTR {
			break
		}
	}
	return err
}

func unlock(f os.File) error {
	return lock(f, syscall.LOCK_UN)
}
