package main

import "fmt"

// define Reader interface
type Reader interface {
	read(b []byte) (int, error)
}

// define Writer interface
type Writer interface {
	write(b []byte) (int, error)
}

// here we have a composite interface that descibes the behaviour of an entity that can both read and write
type ReaderWriter interface {
	Reader
	Writer
}

// concrete data types:

type file struct {
	name string
}

type pipe struct {
	name string
}

type system struct {
	Host string
}

// we need to define read methods for file and pipe to satisfy the Reader interface
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// system implements all three interfaces: Reader, Writer and ReaderWriter
func (s *system) read(b []byte) (int, error) {
	var i int = 100
	var err error
	return 0, err
}

func (s *system) write(b []byte) (int, error) {
	var i int = 100
	var err error
	return 0, err
}

// We can now write a retrieve function which is capable of taking a file or a pipe
// we no longer need to write two functions, we have one reusable piece of code
// Polymorphism in action
func retrieve(r Reader) error {
	data := make([]byte, 0)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

func main() {
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	retrieve(f)
	retrieve(p)

	// interface value for ReaderWriter
	var rw ReaderWriter
	// interface value for Reader
	var r Reader
	// interface value for Writer
	var w Writer

	s := system{"127.0.0.0"}
	rw = &s

	// we can do these next two assignments because the compiler knows that rw also satisfies the reqs for r and w
	r = rw
	w = rw

}
