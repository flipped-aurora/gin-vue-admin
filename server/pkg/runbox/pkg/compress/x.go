package compress

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"github.com/mholt/archiver/v3"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// DeCompressx 使用 archiver 解压缩文件
func DeCompressx(compressFilePath, dest string) error {
	// 创建输出目录
	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	// 使用 archiver 解压缩
	if err := archiver.Unarchive(compressFilePath, dest); err != nil {
		return fmt.Errorf("error unarchiving file: %w", err)
	}

	return nil
}

// unzip 解压 ZIP 文件
func unzip(r io.Reader, dest string) (string, error) {
	zipReader, err := zip.NewReader(r.(io.ReaderAt), r.(os.FileInfo).Size())
	if err != nil {
		return "", err
	}

	for _, f := range zipReader.File {
		fPath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fPath, 0755)
			continue
		}

		outFile, err := os.Create(fPath)
		if err != nil {
			return "", err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return "", err
		}
	}
	return dest, nil
}

// ungzip 解压 Gzip 文件
func ungzip(r io.Reader, compressFilePath, dest string) (string, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	// 从压缩文件名中提取输出文件名
	outputFileName := strings.TrimSuffix(filepath.Base(compressFilePath), ".gz")
	outFile, err := os.Create(filepath.Join(dest, outputFileName))
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, gzr)
	if err != nil {
		return "", err
	}

	return dest, nil
}

// untar 解压 Tar 文件
func untar(r io.Reader, dest string) (string, error) {
	tr := tar.NewReader(r)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		fPath := filepath.Join(dest, header.Name)

		if header.Typeflag == tar.TypeDir {
			os.MkdirAll(fPath, 0755)
			continue
		}

		outFile, err := os.Create(fPath)
		if err != nil {
			return "", err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, tr)
		if err != nil {
			return "", err
		}
	}

	return dest, nil
}
