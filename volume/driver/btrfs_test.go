package driver_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/pivotal-golang/lager/lagertest"

	"github.com/concourse/baggageclaim/fs"
	"github.com/concourse/baggageclaim/volume/driver"
)

var _ = Describe("BtrFS", func() {
	if runtime.GOOS != "linux" {
		fmt.Println("\x1b[33m*** skipping btrfs tests because non-linux ***\x1b[0m")
		return
	}

	var (
		tempDir    string
		volumeDir  string
		fsDriver   *driver.BtrFSDriver
		filesystem *fs.BtrfsFilesystem
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "baggageclaim_driver_test")
		Expect(err).NotTo(HaveOccurred())

		logger := lagertest.NewTestLogger("fs")

		imagePath := filepath.Join(tempDir, "image.img")
		volumeDir = filepath.Join(tempDir, "mountpoint")

		filesystem = fs.New(logger, imagePath, volumeDir)
		err = filesystem.Create(100 * 1024 * 1024)
		Expect(err).NotTo(HaveOccurred())

		fsDriver = driver.NewBtrFSDriver(logger)
	})

	AfterEach(func() {
		err := filesystem.Delete()
		Expect(err).NotTo(HaveOccurred())

		err = os.RemoveAll(tempDir)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Lifecycle", func() {
		It("can create and delete a subvolume", func() {
			subvolumePath := filepath.Join(volumeDir, "subvolume")

			err := fsDriver.CreateVolume(subvolumePath)
			Expect(err).NotTo(HaveOccurred())

			Expect(subvolumePath).To(BeADirectory())

			checkSubvolume := exec.Command("btrfs", "subvolume", "show", subvolumePath)
			session, err := gexec.Start(checkSubvolume, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session).Should(gbytes.Say(subvolumePath))
			Eventually(session).Should(gexec.Exit(0))

			err = fsDriver.DestroyVolume(subvolumePath)
			Expect(err).NotTo(HaveOccurred())

			Expect(subvolumePath).NotTo(BeADirectory())
		})

		It("can delete parent volume when it has subvolumes", func() {
			parentVolumePath := filepath.Join(volumeDir, "parent-volume")
			err := fsDriver.CreateVolume(parentVolumePath)
			Expect(err).NotTo(HaveOccurred())

			childVolumePath := filepath.Join(volumeDir, "parent-volume", "child-volume")
			err = fsDriver.CreateVolume(childVolumePath)
			Expect(err).NotTo(HaveOccurred())

			grandchildVolumePath := filepath.Join(volumeDir, "parent-volume", "child-volume", "grandchild-volume")
			err = fsDriver.CreateVolume(grandchildVolumePath)
			Expect(err).NotTo(HaveOccurred())

			err = fsDriver.DestroyVolume(parentVolumePath)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("GetVolumeSize", func() {
		var parentVolumePath string
		var childVolumePath string

		BeforeEach(func() {
			parentVolumePath = filepath.Join(volumeDir, "parent-volume")
			err := fsDriver.CreateVolume(parentVolumePath)
			Expect(err).NotTo(HaveOccurred())

			bs := make([]byte, 4096)
			for i := 0; i < 4096; i++ {
				bs[i] = 'i'
			}
			err = ioutil.WriteFile(filepath.Join(parentVolumePath, "parent-stuff"), bs, os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := fsDriver.DestroyVolume(childVolumePath)
			Expect(err).NotTo(HaveOccurred())

			err = fsDriver.DestroyVolume(parentVolumePath)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the size of the volume at the given path", func() {
			childVolumePath = filepath.Join(volumeDir, "parent-volume", "child-volume")
			err := fsDriver.CreateVolume(childVolumePath)
			Expect(err).NotTo(HaveOccurred())

			bs := make([]byte, 1024)
			for i := 0; i < 1024; i++ {
				bs[i] = 'i'
			}
			err = ioutil.WriteFile(filepath.Join(childVolumePath, "child-stuff"), bs, os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() uint {
				GinkgoRecover()
				newSize, err := fsDriver.GetVolumeSize(childVolumePath)
				Expect(err).NotTo(HaveOccurred())
				return uint(newSize)
			}, 1*time.Minute, 1*time.Second).Should(Equal(uint(4096))) // 4096B page size
		})
	})
})
