package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func TestMarshalHashAccount(t *testing.T) {
	v := Account{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashAccount(b *testing.B) {
	v := Account{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgAccount(b *testing.B) {
	v := Account{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashSQLChainProfile(t *testing.T) {
	v := SQLChainProfile{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashSQLChainProfile(b *testing.B) {
	v := SQLChainProfile{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgSQLChainProfile(b *testing.B) {
	v := SQLChainProfile{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashSQLChainUser(t *testing.T) {
	v := SQLChainUser{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashSQLChainUser(b *testing.B) {
	v := SQLChainUser{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgSQLChainUser(b *testing.B) {
	v := SQLChainUser{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashTokenList(t *testing.T) {
	v := TokenList{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashTokenList(b *testing.B) {
	v := TokenList{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgTokenList(b *testing.B) {
	v := TokenList{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}
