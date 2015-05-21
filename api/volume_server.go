package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nu7hatch/gouuid"
	"github.com/pivotal-golang/lager"
)

type VolumeResponse struct {
	GUID string `json:"guid"`
	Path string `json:"path"`
}

type VolumesResponse []VolumeResponse

type VolumeServer struct {
	volumeDir string

	logger lager.Logger
}

var ErrCreateVolumeFailed = errors.New("failed to create volume")
var ErrListVolumesFailed = errors.New("failed to list volumes")

func NewVolumeServer(logger lager.Logger, volumeDir string) *VolumeServer {
	return &VolumeServer{
		volumeDir: volumeDir,
		logger:    logger,
	}
}

func (vs *VolumeServer) CreateVolume(w http.ResponseWriter, req *http.Request) {
	guid, err := uuid.NewV4()
	if err != nil {
		vs.logger.Error("failed-to-generate-guid", err)
		respondWithError(w, ErrCreateVolumeFailed)
		return
	}

	createdVolume := filepath.Join(vs.volumeDir, guid.String())

	err = os.MkdirAll(createdVolume, 0755)
	if err != nil {
		vs.logger.Error("failed-to-make-dir", err, lager.Data{
			"volume-path": createdVolume,
		})

		respondWithError(w, ErrCreateVolumeFailed)
		return
	}

	createVolumeResponse := VolumeResponse{
		Path: createdVolume,
		GUID: guid.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createVolumeResponse); err != nil {
		vs.logger.Error("failed-to-encode", err, lager.Data{
			"volume-path": createdVolume,
		})
	}
}

func (vs *VolumeServer) GetVolumes(w http.ResponseWriter, req *http.Request) {
	volumes, err := ioutil.ReadDir(vs.volumeDir)
	if err != nil {
		vs.logger.Error("failed-to-list-dirs", err, lager.Data{
			"volume-dir": vs.volumeDir,
		})
		respondWithError(w, ErrListVolumesFailed)
		return
	}

	response := make(VolumesResponse, 0, len(volumes))
	for _, volume := range volumes {
		response = append(response, VolumeResponse{
			GUID: volume.Name(),
			Path: filepath.Join(vs.volumeDir, volume.Name()),
		})
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		vs.logger.Error("failed-to-encode", err, lager.Data{
			"volume-dir": vs.volumeDir,
		})
	}
}