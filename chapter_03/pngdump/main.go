package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	// 長さを表す32bit(4byte)を読み込む
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)

	// チャンク名を表す4byte分だけ読み込む
	buffer := make([]byte, 4)
	chunk.Read(buffer)

	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	// シグネチャ部分を飛ばす
	file.Seek(8, 0)
	// オフセットは初期値は8
	var offset int64 = 8

	for {
		// 32bit(4byte)が今のチャンクの長さを表すので4byteだけ読み込む
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		fmt.Println(length)
		if err == io.EOF {
			fmt.Println("break")
			break
		}

		// offsetは今読み込もうとしているチャンクの先頭までずらす値
		// lengthはデータ部の長さであって、1つのチャンクの中にはデータ部以外に長さ部分(4byte)、チャンク名部分(4byte)、CRC部分(4byte)が含まれている
		// 一つ分のSectionReaderをchunksへ追加する
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// 今カーソルがあるところは長さ部分だけを読み終わった位置になっている（上記のNewSectionReaderはReaderを作っただけで読んだわけではない）
		// 今カーソルがある場所を起点にして、前のチャンクの長さ+8（チャンク名とCRC）を飛ばす
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
