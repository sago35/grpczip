package server

import (
	"archive/zip"
	"bytes"
	"path/filepath"

	"github.com/sago35/grpczip"
	"github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

func (s GrpczipServer) Grpczip(ctx context.Context, r *grpczip.Request) (*grpczip.Response, error) {
	logrus.Info("Grpczip()")
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	for _, file := range r.GetFiles() {
		f, err := w.Create(filepath.Base(file.GetFilename()))
		if err != nil {
			return nil, err
		}
		_, err = f.Write(file.GetData())
		if err != nil {
			return nil, err
		}
	}
	err := w.Close()
	if err != nil {
		return nil, err
	}

	res := &grpczip.Response{
		Id: "zipped",
		Zipfile: &grpczip.File{
			Filename: "abc.zip",
			Data:     buf.Bytes(),
		},
	}
	return res, nil
}
