package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPublishPreservesUnownedDirectory(t *testing.T) {
	root := t.TempDir()
	stage, destination := filepath.Join(root, "stage"), filepath.Join(root, "arctex-client-go")
	if err := os.Mkdir(stage, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(destination, 0755); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(destination, "user-file")
	if err := os.WriteFile(file, []byte("keep"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := publish(stage, destination); err == nil {
		t.Fatal("replaced unowned directory")
	}
	if raw, err := os.ReadFile(file); err != nil || string(raw) != "keep" {
		t.Fatal("user contents changed")
	}
}

func TestPublishReplacesOnlyOwnedGeneratedDirectory(t *testing.T) {
	root := t.TempDir()
	stage, destination := filepath.Join(root, "stage"), filepath.Join(root, "arctex-client-go")
	for _, directory := range []string{stage, destination} {
		if err := os.Mkdir(directory, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(directory, marker), []byte(module+"\n"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(destination, "stale"), []byte("stale"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stage, "fresh"), []byte("fresh"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := publish(stage, destination); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "stale")); !os.IsNotExist(err) {
		t.Fatal("retained stale generated output")
	}
	if _, err := os.Stat(filepath.Join(destination, "fresh")); err != nil {
		t.Fatal(err)
	}
}
