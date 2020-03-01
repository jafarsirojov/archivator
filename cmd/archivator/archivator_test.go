package main

import (
	"log"
	"os"
	"testing"
)

func Benchmark_conArchive(b *testing.B) {
	err := os.Chdir("../../")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		data := []string{
			"a.exe",
			"a.txt",
			"b.exe",
			"b.txt",
			"c.exe",
			"c.txt",
			"d.exe",
			"d.txt",
			"f.exe",
			"f.txt",
			"ChromeSetup.exe",
			"go.pdf",
		}
		conArchive(data)
	}
}

func Benchmark_seqArchive(b *testing.B) {
	err := os.Chdir("../../")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		data := []string{
			"a.exe",
			"a.txt",
			"b.exe",
			"b.txt",
			"c.exe",
			"c.txt",
			"d.exe",
			"d.txt",
			"f.exe",
			"f.txt",
			"ChromeSetup.exe",
			"go.pdf",
		}
		seqArchive(data)
	}
}
