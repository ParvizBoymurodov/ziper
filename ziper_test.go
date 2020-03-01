package main

import "testing"

func Benchmark_conArchive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conArchivator([]string{
			"par",
			"bar",
			"par",
			"bar",
			"par",
			"bar",
			"par",
			"bar",
			"par",
			"bar",
		})
	}
}

func Benchmark_Archive(b *testing.B){
	for i := 0; i < b.N; i++{
		archivator([]string{
			"par",
			"bar",
			"par",
			"bar",
			"par",
			"bar",
			"par",
			"bar",
			"par",
			"bar",
		})
	}
}