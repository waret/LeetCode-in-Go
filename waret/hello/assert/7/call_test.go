package main

import (
	"fmt"
	"testing"
)

type identifier interface {
	idInline() int32
	idNoInline() int32
}
type id32 struct{ id int32 }

func (id *id32) idInline() int32 { return id.id }

//go:noinline
func (id *id32) idNoInline() int32 { return id.id }

var escapeMePlease *id32

//主要作用是强制变量内存在 heap 上分配
//go:noinline
func escapeToHeap(id *id32) identifier {
	escapeMePlease = id
	return escapeMePlease
}

func BenchmarkMethodCall_direct(b *testing.B) {
	var myID int32

	//直接调用
	b.Run("single/noinline", func(b *testing.B) {
		m := escapeToHeap(&id32{id: 6754}).(*id32)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			//CALL "".(*id32).idNoinline(SB)
			//MOVL 8(SP), AX
			//MOVQ "".&myID+40(SP), CX
			//MOVL AX, (CX)
			myID = m.idNoInline()
		}
	})

	b.Run("single/inline", func(b *testing.B) {
		m := escapeToHeap(&id32{id: 6754}).(*id32)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			//MOVL (DX), SI
			//MOVL SI, (CX)
			myID = m.idInline()
		}
	})

	//接口调用
	b.Run("single/noinline", func(b *testing.B) {
		m := escapeToHeap(&id32{id: 6754})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// MOVQ 32(AX), CX
			// MOVQ "".m.data+40(SP), DX
			// MOVQ DX, (SP)
			// CALL CX
			// MOVL 8(SP), AX
			// MOVQ "".&myID+48(SP), CX
			// MOVL AX, (CX)
			myID = m.idNoInline()
		}
	})

	b.Run("single/inline", func(b *testing.B) {
		m := escapeToHeap(&id32{id: 6754})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			//MOVQ 24(AX), CX
			//MOVQ "".m.data+40(SP), DX
			//MOVQ DX, (SP)
			//CALL CX
			//MOVL 8(SP), AX
			//MOVQ "". &myID+48(SP), ex
			//MOVL AX, (CX)
			myID = m.idInline()
		}
	})
	fmt.Println(myID)
}

func main() {}
