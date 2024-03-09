package service

import (
	"log"
	"os"
	"strconv"

	"github.com/Metadiv-Atomic-Engine/atomic-util/static/config"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var ContentFileService = new(contentFileService)

type contentFileService struct{}

func (c *contentFileService) Save(staticUUID string, content []byte, workspaceID uint) (ok bool) {
	if !c.createStaticFolderIfNotExists() {
		return false
	}
	if !c.createWorkspaceFolderIfNotExists(workspaceID) {
		return false
	}
	err := os.WriteFile(c.formFilePath(workspaceID, staticUUID), content, os.ModePerm)
	return err == nil
}

func (c *contentFileService) Get(workspaceID uint, staticUUID string) []byte {
	content, err := os.ReadFile(c.formFilePath(workspaceID, staticUUID))
	if err != nil {
		log.Println(err)
		return nil
	}
	return EncryptService.Decrypt(content)
}

func (c *contentFileService) Delete(workspaceID uint, staticUUID string) (ok bool) {
	err := os.Remove(c.formFilePath(workspaceID, staticUUID))
	return err == nil
}

func (c *contentFileService) formFilePath(workspaceID uint, staticUUID string) (path string) {
	return atomic.Engine.EnvString(config.STATIC_FOLDER_NAME) + "/" + strconv.Itoa(int(workspaceID)) + "/" + staticUUID + ".bin"
}

func (c *contentFileService) createStaticFolderIfNotExists() (ok bool) {
	_, err := os.Stat(atomic.Engine.EnvString(config.STATIC_FOLDER_NAME))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(atomic.Engine.EnvString(config.STATIC_FOLDER_NAME), os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

func (c *contentFileService) createWorkspaceFolderIfNotExists(workspaceID uint) (ok bool) {
	_, err := os.Stat(atomic.Engine.EnvString(config.STATIC_FOLDER_NAME) + "/" + strconv.Itoa(int(workspaceID)))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(atomic.Engine.EnvString(config.STATIC_FOLDER_NAME)+"/"+strconv.Itoa(int(workspaceID)), os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}
