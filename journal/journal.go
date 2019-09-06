package journal

import (
	"encoding/json"
	"infoedge/journalapp/helpers"
	"time"
)

type Journal struct {
	Session string
}

type JournalData struct {
	Time      time.Time
	ToRemeber string
}

type File struct {
	Pass  string
	Data  []JournalData
}

func GetInstace(sessionId string) (*Journal, error) {

	return &Journal{
		Session: sessionId,
	}, nil
}

func (this *Journal) ListEntries() ([]JournalData, error) {

	data, err := helpers.GetFileData(this.Session)
	if err != nil {
		return nil, err
	}
	var fData File
	err = json.Unmarshal(data, &fData)
	if err != nil {
		return nil, err
	}
	return fData.Data, nil
}

func (this *Journal) InputEntry(entry string) error {

	data, err := helpers.GetFileData(this.Session)
	if err != nil {
		return err
	}
	start := time.Now()
	var fData File
	err = json.Unmarshal(data, &fData)
	if err != nil {
		return err
	}
	if len(fData.Data) >= 10 {
		fData.Data = append(fData.Data[:1], fData.Data[1+1:]...)
	}
	fData.Data = append(fData.Data, JournalData{
		Time:      start,
		ToRemeber: entry,
	})
	fBytes, err := json.Marshal(fData)

	if err != nil {
		return err
	}
	err = helpers.Write2File(fBytes, this.Session)
	if err != nil {
		return err
	}

	return nil
}
