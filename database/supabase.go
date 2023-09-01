package database

import (
	"mime/multipart"
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Supabase struct {
	supabase *supabasestorageuploader.Client
}

func NewSupabaseClient() *Supabase {
	client := supabasestorageuploader.New(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_TOKEN"),
		os.Getenv("SUPABASE_BUCKET"),
	)
	return &Supabase{
		supabase: client,
	}
}

func (s *Supabase) UploadFile(file *multipart.FileHeader) (string, error) {
	fileString, err := s.supabase.Upload(file)
	if err != nil {
		return "", err
	}
	return fileString, nil
}

func (s *Supabase) DeleteFile(link string) error {
	err := s.supabase.Delete(link)
	if err != nil {
		return err
	}
	return nil
}
