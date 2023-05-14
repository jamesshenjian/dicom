## Important notes

:eyes: 
This project is forked from suyashkumar's original project at https://github.com/suyashkumar/dicom
Major changes are made to dataset struct, from a slice to a map type, thus allowing easier get/set of tag values

This is a library and command-line tool to read, write, and generally work with DICOM medical image files in native Go. The goal is to build a full-featured, high-performance, and readable DICOM parser for the Go community.

## CLI Tool
A CLI tool that uses this package to parse imagery and metadata out of DICOMs is provided in the `cmd/dicomutil` package. This tool can take in a DICOM, and dump out all the elements to STDOUT, in addition to writing out any imagery to the current working directory either as PNGs or JPEG (note, it does not perform any automatic color rescaling by default).

### Usage
```
dicomutil -path myfile.dcm
```
Note: for some DICOMs (with native pixel data) no automatic intensity scaling is applied yet (this is coming). You can apply this in your image viewer if needed (in Preview on mac, go to Tools->Adjust Color). 


### Build manually
To build manually, ensure you have `make` and `go` installed. Clone (or `go get`) this repo into your `$GOPATH` and then simply run:
```sh
make
```
Which will build the dicomutil binary and include it in a `build/` folder in your current working directory. 

You can also built it using Go directly:

```sh
go build -o dicomutil ./cmd/dicomutil
```