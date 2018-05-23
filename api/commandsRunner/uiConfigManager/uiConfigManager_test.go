/*
###############################################################################
# Licensed Materials - Property of IBM Copyright IBM Corporation 2017, 2018. All Rights Reserved.
# U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP
# Schedule Contract with IBM Corp.
#
# Contributors:
#  IBM Corporation - initial API and implementation
###############################################################################
*/
package uiConfigManager

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.ibm.com/IBMPrivateCloud/cfp-commands-runner/api/commandsRunner/extensionManager"
	"github.ibm.com/IBMPrivateCloud/cfp-commands-runner/api/commandsRunner/global"
)

const COPYRIGHT_TEST string = `###############################################################################
# Licensed Materials - Property of IBM Copyright IBM Corporation 2017, 2018. All Rights Reserved.
# U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP
# Schedule Contract with IBM Corp.
#
# Contributors:
#  IBM Corporation - initial API and implementation
###############################################################################`

/*
func TestGetUIConfig(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	//	extensionManager.SetExtensionEmbeddedFile("../../test/resource/extensions/ibm-test-extensions.txt")
	//extensionManager.SetExtensionPath("../../test/resource/")
	data, err := GetUIConfig("ui-cf-deploy-vmware")
	if err != nil {
		t.Error(err.Error())
	}
	expected, _ := ioutil.ReadFile("../../resource/ui-cf-deploy-vmware.json")
	if string(data) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(data), string(expected))
	}

}
*/
func TestGetUIConfigExtentionTest(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	global.SetExtensionResourcePath("../../test/resource/extensions")
	extensionManager.SetExtensionEmbeddedFile("../../test/resource/extensions/ibm-test-extensions.txt")
	extensionManager.SetExtensionPath("../../test/resource/extensions/")
	_, err := GetUIConfig("cfp-ext-template")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUIConfigError(t *testing.T) {
	_, err := GetUIConfig("does-not-exist")
	if err == nil {
		t.Error("An error should be raised as this file doesn't exists")
	}
}
