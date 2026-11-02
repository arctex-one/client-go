// prepare materializes the pinned public module in an ignored consumer build
// directory and regenerates its API from source before every consumer build.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const marker = ".arctex-client-generated"
const module = "github.com/arctex-one/client-go"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: prepare path/to/ignored/arctex-client-go")
		os.Exit(2)
	}
	if err := prepare(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func prepare(destination string) error {
	destination, err := filepath.Abs(destination)
	if err != nil {
		return err
	}
	if filepath.Base(destination) != "arctex-client-go" {
		return errors.New("generated module destination must be named arctex-client-go")
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		return err
	}
	lock := destination + ".lock"
	if err := os.Mkdir(lock, 0700); err != nil {
		return fmt.Errorf("acquire client generation lock: %w", err)
	}
	defer os.Remove(lock)
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("cannot locate pinned client source")
	}
	source := filepath.Clean(filepath.Join(filepath.Dir(file), "../.."))
	if relative, err := filepath.Rel(source, destination); err != nil || relative == "." || !strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return errors.New("client output must be outside its source directory")
	}
	mod, err := os.ReadFile(filepath.Join(source, "go.mod"))
	if err != nil || !strings.HasPrefix(string(mod), "module "+module+"\n") {
		return errors.New("cannot verify pinned client source module")
	}
	stage, err := os.MkdirTemp(filepath.Dir(destination), ".arctex-client-build-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(stage)
	err = filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if relative == "." {
			return nil
		}
		if entry.IsDir() && (entry.Name() == ".git" || entry.Name() == "generated") {
			return filepath.SkipDir
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return errors.New("symlink in client source")
		}
		target := filepath.Join(stage, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		if !entry.Type().IsRegular() {
			return errors.New("non-regular client source")
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, raw, 0644)
	})
	if err != nil {
		return err
	}
	command := exec.Command("go", "run", "./tools/generate", "./contract/openapi.json", "./generated")
	command.Dir = stage
	command.Env = append(os.Environ(), "GOWORK=off")
	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	if err := command.Run(); err != nil {
		return fmt.Errorf("generate public client: %w", err)
	}
	if err := os.WriteFile(filepath.Join(stage, marker), []byte(module+"\n"), 0644); err != nil {
		return err
	}
	return publish(stage, destination)
}

func publish(stage, destination string) error {
	previous := stage + ".previous"
	if info, err := os.Lstat(destination); err == nil {
		if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			return errors.New("refusing to replace non-directory client destination")
		}
		owned, err := os.ReadFile(filepath.Join(destination, marker))
		if err != nil || string(owned) != module+"\n" {
			return errors.New("refusing to replace unowned client directory")
		}
		if err := os.Rename(destination, previous); err != nil {
			return err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.Rename(stage, destination); err != nil {
		if _, statErr := os.Stat(previous); statErr == nil {
			if restoreErr := os.Rename(previous, destination); restoreErr != nil {
				return errors.Join(err, restoreErr)
			}
		}
		return err
	}
	return os.RemoveAll(previous)
}
