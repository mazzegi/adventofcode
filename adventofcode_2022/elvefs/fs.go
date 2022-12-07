package elvefs

import (
	"fmt"
	"strconv"
	"strings"
)

type Dir struct {
	Name    string
	Parent  *Dir
	SubDirs []*Dir
	Files   []*File
}

func (d *Dir) subDir(name string) (*Dir, bool) {
	for _, sd := range d.SubDirs {
		if sd.Name == name {
			return sd, true
		}
	}
	return nil, false
}

func (d *Dir) file(name string) (*File, bool) {
	for _, f := range d.Files {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
}

func (d *Dir) ensureSubDir(name string) {
	if _, ok := d.subDir(name); ok {
		return
	}
	d.SubDirs = append(d.SubDirs, &Dir{
		Name:   name,
		Parent: d,
	})
}

func (d *Dir) ensureFile(name string, size int) {
	if _, ok := d.file(name); ok {
		return
	}
	d.Files = append(d.Files, &File{
		Name: name,
		Size: size,
	})
}

func (d *Dir) TotalSize() int {
	var tsz int
	for _, f := range d.Files {
		tsz += f.Size
	}
	for _, sd := range d.SubDirs {
		tsz += sd.TotalSize()
	}
	return tsz
}

//

type File struct {
	Name string
	Size int
}

func New() *Fs {
	root := &Dir{
		Name: "/",
	}
	return &Fs{
		Root: root,
		curr: root,
	}
}

type Fs struct {
	Root *Dir
	curr *Dir
}

type io struct {
	in  string
	out []string
}

func (io io) isZero() bool {
	return io.in == ""
}

func (fs *Fs) Process(lines []string) error {
	var ios []io
	var curr io
	flush := func() {
		if !curr.isZero() {
			ios = append(ios, curr)
		}
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			//flush current io
			flush()
			curr = io{
				in: strings.TrimSpace(strings.TrimPrefix(line, "$")),
			}
			continue
		}
		if curr.isZero() {
			return fmt.Errorf("output without input %q", line)
		}
		curr.out = append(curr.out, strings.TrimSpace(line))
	}
	flush()
	for _, io := range ios {
		err := fs.processIO(io)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fs *Fs) processIO(io io) error {
	switch {
	case strings.HasPrefix(io.in, "cd"):
		if len(io.out) > 0 {
			return fmt.Errorf("cannot process cd with non-empty output")
		}
		return fs.processCD(strings.TrimSpace(strings.TrimPrefix(io.in, "cd")))
	case io.in == "ls":
		return fs.processLS(io.out)
	default:
		return fmt.Errorf("unknown cmomand %q", io.in)
	}
}

func (fs *Fs) processCD(dir string) error {
	if dir == "/" {
		fs.curr = fs.Root
		return nil
	}
	if dir == ".." {
		if fs.curr.Parent == nil {
			return fmt.Errorf("cannot move out of %q", fs.curr.Name)
		}
		fs.curr = fs.curr.Parent
		return nil
	}
	sub, ok := fs.curr.subDir(dir)
	if !ok {
		return fmt.Errorf("dir %q has no subdir named %q", fs.curr.Name, dir)
	}
	fs.curr = sub
	return nil
}

func (fs *Fs) processLS(out []string) error {
	for _, item := range out {
		v, name, ok := strings.Cut(item, " ")
		if !ok {
			return fmt.Errorf("invalid ls item %q", item)
		}
		v = strings.TrimSpace(v)
		name = strings.TrimSpace(name)
		if v == "" || name == "" {
			return fmt.Errorf("invalid output %q", item)
		}

		if v == "dir" {
			fs.curr.ensureSubDir(name)
			continue
		}
		// a file
		sz, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("parse filesize %q: %w", v, err)
		}
		fs.curr.ensureFile(name, sz)
	}
	return nil
}

type FileType string

const (
	TypeDir  FileType = "dir"
	TypeFile FileType = "file"
)

type FileInfo struct {
	Path      string
	Type      FileType
	TotalSize int
}

func (fs *Fs) Used() int {
	return fs.Root.TotalSize()
}

func (fs *Fs) Walk(cb func(fi FileInfo)) {
	cb(FileInfo{
		Path:      "/",
		Type:      TypeDir,
		TotalSize: fs.Root.TotalSize(),
	})
	fs.walkDir(fs.Root, "/", cb)
}

func (fs *Fs) walkDir(dir *Dir, parentPath string, cb func(fi FileInfo)) {
	for _, f := range dir.Files {
		cb(FileInfo{
			Path:      parentPath + f.Name,
			Type:      TypeFile,
			TotalSize: f.Size,
		})
	}
	for _, sd := range dir.SubDirs {
		cb(FileInfo{
			Path:      parentPath + sd.Name,
			Type:      TypeDir,
			TotalSize: sd.TotalSize(),
		})
		fs.walkDir(sd, parentPath+sd.Name+"/", cb)
	}
}
