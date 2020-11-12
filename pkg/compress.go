package pkg

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func CompressFile(path, outPath string) error {
	byt, err := CompressBytes(path)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outPath, byt, 0700)
}

func CompressBytes(path string) ([]byte, error) {
	reader, err := CompressReader(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

func CompressReader(dir string) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	w := zip.NewWriter(buf)
	defer w.Close()
	f, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}

	if !f.IsDir() {
		if err := write(w, dir, path.Base(dir)); err != nil {
			return nil, err
		}
	} else {
		err := travelDir(w, dir, "")
		if err != nil {
			return nil, err
		}
	}
	return buf, err
}

func travelDir(w *zip.Writer, basePath, baseInZip string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			err := travelDir(w, path.Join(basePath, file.Name()), path.Join(baseInZip, file.Name()))
			if err != nil {
				return err
			}
		} else {
			err := write(w, path.Join(basePath, file.Name()), path.Join(baseInZip, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func write(w *zip.Writer, filePath, zipFilePath string) error {
	f, err := w.Create(zipFilePath)
	if err != nil {
		return err
	}
	byt, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if _, err = f.Write(byt); err != nil {
		return err
	}
	return nil
}
