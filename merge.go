package main

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// MergeWAVBytes 修复后的合并函数
func MergeWAVBytes(wavs [][]byte) ([]byte, error) {
	if len(wavs) == 0 {
		return nil, nil
	}
	if len(wavs) == 1 {
		return wavs[0], nil
	}

	// 使用第一个文件的格式参数
	var (
		sampleRate     int
		bitDepth       int
		numChans       int
		wavAudioFormat int
		combinedData   []int
	)

	baseDec, baseBuf, err := decodeWAV(wavs[0])
	if err != nil {
		return nil, err
	}

	sampleRate = int(baseDec.SampleRate)
	bitDepth = int(baseDec.BitDepth)
	numChans = int(baseDec.NumChans)
	wavAudioFormat = int(baseDec.WavAudioFormat)

	for _, data := range wavs {
		_, buf, err := decodeWAV(data)
		if err != nil {
			return nil, err
		}

		combinedData = append(combinedData, buf.Data...)
	}

	// 合并数据（代码同前）
	mergedBuf := &audio.IntBuffer{
		Format:         baseBuf.Format,
		SourceBitDepth: baseBuf.SourceBitDepth,
		Data:           combinedData,
	}

	// 创建内存WriteSeeker
	out := &byteWriteSeeker{}

	// 初始化编码器
	enc := wav.NewEncoder(out, sampleRate, bitDepth, numChans, wavAudioFormat)

	// 写入数据
	if err := enc.Write(mergedBuf); err != nil {
		return nil, fmt.Errorf("编码失败: %v", err)
	}

	// 关闭编码器（自动更新头部）
	if err := enc.Close(); err != nil {
		return nil, fmt.Errorf("编码关闭失败: %v", err)
	}

	return out.data, nil
}

// 解码函数保持不变
func decodeWAV(data []byte) (*wav.Decoder, *audio.IntBuffer, error) {
	dec := wav.NewDecoder(bytes.NewReader(data))
	if !dec.IsValidFile() {
		return nil, nil, errors.New("无效WAV文件")
	}

	buf, err := dec.FullPCMBuffer()
	if err != nil {
		return nil, nil, err
	}

	return dec, buf, nil
}
