package linkedlist

// import (
// 	"bytes"
// 	"io"
// 	"os"
// 	"testing"
// )

// // Helper to build a preset list.
// func sampleList() *LinkedList {
// 	l := &LinkedList{}
// 	l.AddLast(10)
// 	l.AddLast(20)
// 	l.AddLast(30)
// 	return l
// }

// func TestAddFirst(t *testing.T) {
// 	l := LinkedList{}
// 	l.AddFirst(5)
// 	if l.head.data != 5 {
// 		t.Fatalf("Expected head to be 5, got %d", l.head.data)
// 	}
// 	if l.size != 1 {
// 		t.Fatalf("Expected size 1, got %d", l.size)
// 	}
// }

// func TestAddLast(t *testing.T) {
// 	l := sampleList()
// 	l.AddLast(40)
// 	if l.size != 4 {
// 		t.Fatalf("Expected size 4, got %d", l.size)
// 	}

// 	if l.head.next.next.next.data != 40 {
// 		t.Fatalf("Expected last element 40")
// 	}
// }

// func TestAddAt(t *testing.T) {
// 	l := sampleList()
// 	l.AddAt(1, 99)

// 	if l.head.next.data != 99 {
// 		t.Fatalf("Expected index 1 to be 99, got %d", l.head.next.data)
// 	}

// 	if l.size != 4 {
// 		t.Fatalf("Expected size 4, got %d", l.size)
// 	}
// }

// func TestRemoveFirst(t *testing.T) {
// 	l := sampleList()
// 	l.RemoveFirst()

// 	if l.head.data != 20 {
// 		t.Fatalf("Expected head to become 20, got %d", l.head.data)
// 	}

// 	if l.size != 2 {
// 		t.Fatalf("Expected size 2, got %d", l.size)
// 	}
// }

// func TestRemoveLast(t *testing.T) {
// 	l := sampleList()
// 	l.RemoveLast()

// 	if l.size != 2 {
// 		t.Fatalf("Expected size 2, got %d", l.size)
// 	}
// 	if l.head.next.data != 20 {
// 		t.Fatalf("Expected last to be 20, got %d", l.head.next.data)
// 	}
// }

// func TestRemoveAt(t *testing.T) {
// 	l := sampleList()
// 	l.RemoveAt(1)

// 	if l.head.next.data != 30 {
// 		t.Fatalf("Expected index 1 to become 30, got %d", l.head.next.data)
// 	}

// 	if l.size != 2 {
// 		t.Fatalf("Expected size 2, got %d", l.size)
// 	}
// }

// func TestFind(t *testing.T) {
// 	l := sampleList()
// 	idx := l.Find(20)
// 	if idx != 1 {
// 		t.Fatalf("Expected index 1, got %d", idx)
// 	}

// 	idx2 := l.Find(99)
// 	if idx2 != -1 {
// 		t.Fatalf("Expected -1 for missing value, got %d", idx2)
// 	}
// }

// func TestLength(t *testing.T) {
// 	l := sampleList()
// 	if l.Length() != 3 {
// 		t.Fatalf("Expected length 3, got %d", l.Length())
// 	}
// }

// // Smoke test for Print().
// func TestPrint(t *testing.T) {
// 	l := sampleList()

// 	// Capture stdout
// 	var buffer bytes.Buffer
// 	orig := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	l.Print()

// 	// Restore stdout
// 	w.Close()
// 	os.Stdout = orig

// 	// Read printed data
// 	output, _ := io.ReadAll(r)
// 	buffer.Write(output)

// 	if buffer.Len() == 0 {
// 		t.Fatalf("Print output is empty")
// 	}
// }
