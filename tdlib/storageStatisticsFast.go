// AUTOGENERATED - DO NOT EDIT

package tdlib

// StorageStatisticsFast Contains approximate storage usage statistics, excluding files of unknown file type
type StorageStatisticsFast struct {
	tdCommon
	FilesSize                int64 `json:"files_size"`                  // Approximate total size of files, in bytes
	FileCount                int32 `json:"file_count"`                  // Approximate number of files
	DatabaseSize             int64 `json:"database_size"`               // Size of the database
	LanguagePackDatabaseSize int64 `json:"language_pack_database_size"` // Size of the language pack database
	LogSize                  int64 `json:"log_size"`                    // Size of the TDLib internal log
}

// MessageType return the string telegram-type of StorageStatisticsFast
func (storageStatisticsFast *StorageStatisticsFast) MessageType() string {
	return "storageStatisticsFast"
}

// NewStorageStatisticsFast creates a new StorageStatisticsFast
//
// @param filesSize Approximate total size of files, in bytes
// @param fileCount Approximate number of files
// @param databaseSize Size of the database
// @param languagePackDatabaseSize Size of the language pack database
// @param logSize Size of the TDLib internal log
func NewStorageStatisticsFast(filesSize int64, fileCount int32, databaseSize int64, languagePackDatabaseSize int64, logSize int64) *StorageStatisticsFast {
	storageStatisticsFastTemp := StorageStatisticsFast{
		tdCommon:                 tdCommon{Type: "storageStatisticsFast"},
		FilesSize:                filesSize,
		FileCount:                fileCount,
		DatabaseSize:             databaseSize,
		LanguagePackDatabaseSize: languagePackDatabaseSize,
		LogSize:                  logSize,
	}

	return &storageStatisticsFastTemp
}
