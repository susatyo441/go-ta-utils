package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

// DeleteImage menghapus file gambar berdasarkan folder dan nama file
func DeleteImage(folderName, fileName string) error {
	// Dapatkan path absolut folder tempat gambar disimpan
	absPath, err := filepath.Abs(filepath.Join("../acts-files", folderName, fileName))
	if err != nil {
		return fmt.Errorf("error mendapatkan path file: %w", err)
	}

	// Periksa apakah file ada sebelum dihapus
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("file tidak ditemukan: %s", absPath)
	}

	// Hapus file
	if err := os.Remove(absPath); err != nil {
		return fmt.Errorf("gagal menghapus file: %w", err)
	}

	return nil
}
