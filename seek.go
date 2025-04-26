package main

import (
	"errors"
	"io"
)

// 实现内存中的WriteSeeker
type byteWriteSeeker struct {
	data []byte
	pos  int
}

func (b *byteWriteSeeker) Write(p []byte) (n int, err error) {
	// 自动扩展切片容量
	needLen := b.pos + len(p)
	if needLen > cap(b.data) {
		newCap := needLen * 2
		newData := make([]byte, len(b.data), newCap)
		copy(newData, b.data)
		b.data = newData
	}

	// 扩展切片长度
	if needLen > len(b.data) {
		b.data = b.data[:needLen]
	}

	copy(b.data[b.pos:], p)
	b.pos += len(p)
	return len(p), nil
}


func (b *byteWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		b.pos = int(offset)
	case io.SeekCurrent:
		b.pos += int(offset)
	case io.SeekEnd:
		b.pos = len(b.data) + int(offset)
	default:
		return 0, errors.New("invalid whence")
	}

	// 防止越界
	if b.pos < 0 {
		b.pos = 0
	}
	return int64(b.pos), nil
}