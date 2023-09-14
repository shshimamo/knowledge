package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	pb "github.com/shshimamo/knowledge/protobufs/example/image_uploader/gen/pb_go"
)

type ImageUploaderHandler struct {
	sync.Mutex // 未使用
	files      map[string][]byte
	pb.UnimplementedImageUploadServiceServer
}

func NewImageUploadHandler() *ImageUploaderHandler {
	return &ImageUploaderHandler{
		files: make(map[string][]byte),
	}
}

func (h *ImageUploaderHandler) Upload(stream pb.ImageUploadService_UploadServer) error {
	// Recv関数で1つ目のリクエストを取得
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	// 初回は必ずメタ情報が送られる
	meta := req.GetFileMeta()
	filename := meta.Filename

	// UUID 生成
	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := u.String()

	// 画像データ格納用 Buffer
	buf := &bytes.Buffer{}

	// 2つ目以降のリクエストをループで取得
	for {
		r, err := stream.Recv()

		// すべてのリクエストを受け取ると、io.EOF が返ってくる
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		// バイナリをバッファに書き込む
		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)

	h.files[filename] = data

	err = h.saveImage()
	if err != nil {
		return err
	}

	// SendAndCloseはレスポンスを送信してストリームを閉じる
	err = stream.SendAndClose(&pb.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    filename,
		ContentType: mimeType,
	})

	return err
}

func (h *ImageUploaderHandler) saveImage() error {
	savePath := "./"
	filename := h.fileNames()[0]

	fullPath := filepath.Join(savePath, filename)

	outFile, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer func() { _ = outFile.Close() }()

	data, exists := h.files[filename]
	if !exists {
		return fmt.Errorf("file data for %s not found", filename)
	}
	_, err = outFile.Write(data)
	if err != nil {
		return err
	}

	return err
}

func (h *ImageUploaderHandler) fileNames() []string {
	filenames := make([]string, 0, len(h.files))

	for filename := range h.files {
		filenames = append(filenames, filename)
	}
	return filenames
}
