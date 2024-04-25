package component

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/*
* @Author <bypanghu> (bypanghu@163.com)
* @Date 2023/12/27 17:16
**/

var MinioPkg = new(minioPkg)

type minioPkg struct {
	conf *config.Minio
}

func (m *minioPkg) Init(minio *config.Minio) {
	defer func() {
		if r := recover(); r != nil {
			println("minio init error")
		}
	}()
	m.conf = minio
}

func (m *minioPkg) NewClient() (*minio.Client, error) {
	if m.conf == nil {
		panic("please make minio init first")
	}
	client, err := minio.New(m.conf.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.conf.AccessKey, m.conf.SecretKey, m.conf.Token),
		Secure: m.conf.UseSSl,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[oss][minio]获取minio.Client对象失败!")
	}
	return client, nil
}

// Filepath 访问全路径
func (m *minioPkg) Filepath(key string, bucket string) string {
	if m.conf == nil {
		panic("please make minio init first")
	}
	return m.conf.Domain + path.Join(bucket, key)
}

// Filename 格式化文件名
func (m *minioPkg) Filename(filename string) string {
	if m.conf == nil {
		panic("please make minio init first")
	}
	names := strings.Split(filename, string(os.PathSeparator))
	filename = names[len(names)-1:][0]
	if m.conf.Prefix == "" {
		return fmt.Sprintf("%s", filename)
	}
	return fmt.Sprintf("%s_%s", m.conf.Prefix, filename)
}

// FileKey 文件key
func (m *minioPkg) FileKey(filename string) string {
	if m.conf == nil {
		panic("please make minio init first")
	}
	return filepath.Join(m.conf.Path, time.Now().Format("2006-01-02"), filename)
}

func (m *minioPkg) FileKeyWithTime(filename string) string {
	if m.conf == nil {
		panic("please make minio init first")
	}
	return filepath.Join(time.Now().Format("2006-01-02"), m.Filename(filename))
}

func (m *minioPkg) UploadByFile(ctx context.Context, file *os.File, bucket string) (filepath string, filename string, err error) {
	if m.conf == nil {
		panic("please make minio init first")
	}
	var client *minio.Client
	client, err = m.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	err = m.MakeBucket(ctx, bucket)
	if err != nil {
		return filepath, filename, err
	}
	filename = m.Filename(file.Name())
	key := m.FileKey(filename)
	filepath = m.Filepath(key, bucket)
	info, _ := file.Stat()
	options := minio.PutObjectOptions{}
	_, err = client.PutObject(ctx, bucket, key, file, info.Size(), options)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]文件上传失败!")
	}
	return filepath, filename, nil
}

func (m *minioPkg) UploadByHeader(ctx context.Context, header *multipart.FileHeader, bucket string) (filepath string, filename string, err error) {
	if m.conf == nil {
		panic("please make minio init first")
	}
	var client *minio.Client
	client, err = m.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	err = m.MakeBucket(ctx, bucket)
	if err != nil {
		return filepath, filename, err
	}
	filename = m.Filename(header.Filename)

	filename = fmt.Sprintf("%d-%s", time.Now().Unix(), filename)

	key := m.FileKey(filename)
	filepath = m.Filepath(key, bucket)
	// 生成时间戳
	time.Now().Unix()

	options := minio.PutObjectOptions{ContentType: header.Header.Get("content-type")}

	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "通过header获取io.Reader对象失败!")
	}
	_, err = client.PutObject(ctx, bucket, key, file, header.Size, options)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]文件上传失败!")
	}
	return filepath, filename, nil
}

func (m *minioPkg) UploadByFilePath(ctx context.Context, file string, bucket string) (filepath string, filename string, err error) {
	if m.conf == nil {
		panic("please make minio init first")
	}
	var client *minio.Client
	client, err = m.NewClient()
	if err != nil {
		return filepath, filename, err
	}
	err = m.MakeBucket(ctx, bucket)
	if err != nil {
		return filepath, filename, err
	}
	filename = m.Filename(file)
	key := m.FileKey(filename)
	filepath = m.Filepath(key, bucket)
	// 获取文件
	fileFile, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]文件上传失败!")
	}
	info, _ := fileFile.Stat()
	options := minio.PutObjectOptions{}
	_, err = client.PutObject(ctx, bucket, key, fileFile, info.Size(), options)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "[oss][minio]文件上传失败!")
	}
	return filepath, filename, nil
}

func (m *minioPkg) MakeBucket(ctx context.Context, bucket string) error {
	if m.conf == nil {
		panic("please make minio init first")
	}
	var client *minio.Client
	client, err := m.NewClient()
	if err != nil {
		return err
	}
	var exists bool
	exists, err = client.BucketExists(ctx, bucket)
	if !exists && err == nil {
		options := minio.MakeBucketOptions{}
		err = client.MakeBucket(ctx, bucket, options)
		if err != nil {
			return errors.Wrapf(err, "[oss][minio][bucket:%s]存储桶创建失败!", bucket)
		}
		policy := fmt.Sprintf("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:GetBucketLocation\",\"s3:ListBucket\",\"s3:ListBucketMultipartUploads\"],\"Resource\":[\"arn:aws:s3:::%s\"]},{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:GetObject\",\"s3:ListMultipartUploadParts\",\"s3:PutObject\",\"s3:AbortMultipartUpload\",\"s3:DeleteObject\"],\"Resource\":[\"arn:aws:s3:::%s/*\"]}]}", bucket, bucket)
		err = client.SetBucketPolicy(ctx, bucket, policy)
		if err != nil {
			return errors.Wrapf(err, "[oss][minio][bucket:%s]设置存储桶权限失败!", bucket)
		}
	}
	return nil
}
