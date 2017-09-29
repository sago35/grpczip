package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/sago35/grpczip"
	"google.golang.org/grpc"
)

func main() {
	filenames := kingpin.Arg("file", "Set filename to zip").Strings()
	kingpin.Parse()

	conn, err := grpc.Dial("localhost:8000", grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(0xFFFFFFFF)), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpczip.NewGrpczipClient(conn)
	files, err := f(*filenames)
	if err != nil {
		log.Fatal(err)
	}

	req := &grpczip.Request{
		Id:    `client`,
		Files: files,
	}

	resp, err := client.Grpczip(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Id)

	z := resp.GetZipfile()
	w, err := os.Create(z.GetFilename())
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	r := bytes.NewReader(z.GetData())
	size, err := io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (%d bytes)\n", z.GetFilename(), size)
}

func f(files []string) ([]*grpczip.File, error) {
	ret := []*grpczip.File{}

	for _, file := range files {
		r, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer r.Close()

		b, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &grpczip.File{
			Filename: file,
			Data:     b,
		})
	}
	return ret, nil
}
