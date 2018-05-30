// this file was generated by gomacro command: import _b "archive/zip"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"archive/zip"
)

// reflection: allow interpreted code to import "archive/zip"
func init() {
	Packages["archive/zip"] = Package{
	Binds: map[string]Value{
		"Deflate":	ValueOf(zip.Deflate),
		"ErrAlgorithm":	ValueOf(&zip.ErrAlgorithm).Elem(),
		"ErrChecksum":	ValueOf(&zip.ErrChecksum).Elem(),
		"ErrFormat":	ValueOf(&zip.ErrFormat).Elem(),
		"FileInfoHeader":	ValueOf(zip.FileInfoHeader),
		"NewReader":	ValueOf(zip.NewReader),
		"NewWriter":	ValueOf(zip.NewWriter),
		"OpenReader":	ValueOf(zip.OpenReader),
		"RegisterCompressor":	ValueOf(zip.RegisterCompressor),
		"RegisterDecompressor":	ValueOf(zip.RegisterDecompressor),
		"Store":	ValueOf(zip.Store),
	}, Types: map[string]Type{
		"Compressor":	TypeOf((*zip.Compressor)(nil)).Elem(),
		"Decompressor":	TypeOf((*zip.Decompressor)(nil)).Elem(),
		"File":	TypeOf((*zip.File)(nil)).Elem(),
		"FileHeader":	TypeOf((*zip.FileHeader)(nil)).Elem(),
		"ReadCloser":	TypeOf((*zip.ReadCloser)(nil)).Elem(),
		"Reader":	TypeOf((*zip.Reader)(nil)).Elem(),
		"Writer":	TypeOf((*zip.Writer)(nil)).Elem(),
	}, Wrappers: map[string][]string{
		"File":	[]string{"FileInfo","ModTime","Mode","SetModTime","SetMode",},
		"ReadCloser":	[]string{"RegisterDecompressor",},
	}, 
	}
}
