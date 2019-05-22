package bolt

import portainer "github.com/portainer/portainer/api"

// Init creates the default data set.
func (store *Store) Init() error {
	groups, err := store.EndpointGroupService.EndpointGroups()
	if err != nil {
		return err
	}

	if len(groups) == 0 {
		unassignedGroup := &portainer.EndpointGroup{
			Name:               "Unassigned",
			Description:        "Unassigned endpoints",
			Labels:             []portainer.Pair{},
			UserAccessPolicies: portainer.UserAccessPolicies{},
			TeamAccessPolicies: portainer.TeamAccessPolicies{},
			Tags:               []string{},
		}

		err = store.EndpointGroupService.CreateEndpointGroup(unassignedGroup)
		if err != nil {
			return err
		}
	}

	roles, err := store.RoleService.Roles()
	if err != nil {
		return err
	}

	if len(roles) == 0 {
		environmentAdministratorRole := &portainer.Role{
			Name:        "Endpoint administrator",
			Description: "Full control on all the resources inside an endpoint",
			Authorizations: map[portainer.Authorization]bool{
				portainer.OperationDockerContainerArchiveInfo:         true,
				portainer.OperationDockerContainerList:                true,
				portainer.OperationDockerContainerExport:              true,
				portainer.OperationDockerContainerChanges:             true,
				portainer.OperationDockerContainerInspect:             true,
				portainer.OperationDockerContainerTop:                 true,
				portainer.OperationDockerContainerLogs:                true,
				portainer.OperationDockerContainerStats:               true,
				portainer.OperationDockerContainerAttachWebsocket:     true,
				portainer.OperationDockerContainerArchive:             true,
				portainer.OperationDockerContainerCreate:              true,
				portainer.OperationDockerContainerPrune:               true,
				portainer.OperationDockerContainerKill:                true,
				portainer.OperationDockerContainerPause:               true,
				portainer.OperationDockerContainerUnpause:             true,
				portainer.OperationDockerContainerRestart:             true,
				portainer.OperationDockerContainerStart:               true,
				portainer.OperationDockerContainerStop:                true,
				portainer.OperationDockerContainerWait:                true,
				portainer.OperationDockerContainerResize:              true,
				portainer.OperationDockerContainerAttach:              true,
				portainer.OperationDockerContainerExec:                true,
				portainer.OperationDockerContainerRename:              true,
				portainer.OperationDockerContainerUpdate:              true,
				portainer.OperationDockerContainerPutContainerArchive: true,
				portainer.OperationDockerContainerDelete:              true,
				portainer.OperationDockerImageList:                    true,
				portainer.OperationDockerImageSearch:                  true,
				portainer.OperationDockerImageGetAll:                  true,
				portainer.OperationDockerImageGet:                     true,
				portainer.OperationDockerImageHistory:                 true,
				portainer.OperationDockerImageInspect:                 true,
				portainer.OperationDockerImageLoad:                    true,
				portainer.OperationDockerImageCreate:                  true,
				portainer.OperationDockerImagePrune:                   true,
				portainer.OperationDockerImagePush:                    true,
				portainer.OperationDockerImageTag:                     true,
				portainer.OperationDockerImageDelete:                  true,
				portainer.OperationDockerImageCommit:                  true,
				portainer.OperationDockerImageBuild:                   true,
				portainer.OperationDockerNetworkList:                  true,
				portainer.OperationDockerNetworkInspect:               true,
				portainer.OperationDockerNetworkCreate:                true,
				portainer.OperationDockerNetworkConnect:               true,
				portainer.OperationDockerNetworkDisconnect:            true,
				portainer.OperationDockerNetworkPrune:                 true,
				portainer.OperationDockerNetworkDelete:                true,
				portainer.OperationDockerVolumeList:                   true,
				portainer.OperationDockerVolumeInspect:                true,
				portainer.OperationDockerVolumeCreate:                 true,
				portainer.OperationDockerVolumePrune:                  true,
				portainer.OperationDockerVolumeDelete:                 true,
				portainer.OperationDockerExecInspect:                  true,
				portainer.OperationDockerExecStart:                    true,
				portainer.OperationDockerExecResize:                   true,
				portainer.OperationDockerSwarmInspect:                 true,
				portainer.OperationDockerSwarmUnlockKey:               true,
				portainer.OperationDockerSwarmInit:                    true,
				portainer.OperationDockerSwarmJoin:                    true,
				portainer.OperationDockerSwarmLeave:                   true,
				portainer.OperationDockerSwarmUpdate:                  true,
				portainer.OperationDockerSwarmUnlock:                  true,
				portainer.OperationDockerNodeList:                     true,
				portainer.OperationDockerNodeInspect:                  true,
				portainer.OperationDockerNodeUpdate:                   true,
				portainer.OperationDockerNodeDelete:                   true,
				portainer.OperationDockerServiceList:                  true,
				portainer.OperationDockerServiceInspect:               true,
				portainer.OperationDockerServiceLogs:                  true,
				portainer.OperationDockerServiceCreate:                true,
				portainer.OperationDockerServiceUpdate:                true,
				portainer.OperationDockerServiceDelete:                true,
				portainer.OperationDockerSecretList:                   true,
				portainer.OperationDockerSecretInspect:                true,
				portainer.OperationDockerSecretCreate:                 true,
				portainer.OperationDockerSecretUpdate:                 true,
				portainer.OperationDockerSecretDelete:                 true,
				portainer.OperationDockerConfigList:                   true,
				portainer.OperationDockerConfigInspect:                true,
				portainer.OperationDockerConfigCreate:                 true,
				portainer.OperationDockerConfigUpdate:                 true,
				portainer.OperationDockerConfigDelete:                 true,
				portainer.OperationDockerTaskList:                     true,
				portainer.OperationDockerTaskInspect:                  true,
				portainer.OperationDockerTaskLogs:                     true,
				portainer.OperationDockerPluginList:                   true,
				portainer.OperationDockerPluginPrivileges:             true,
				portainer.OperationDockerPluginInspect:                true,
				portainer.OperationDockerPluginPull:                   true,
				portainer.OperationDockerPluginCreate:                 true,
				portainer.OperationDockerPluginEnable:                 true,
				portainer.OperationDockerPluginDisable:                true,
				portainer.OperationDockerPluginPush:                   true,
				portainer.OperationDockerPluginUpgrade:                true,
				portainer.OperationDockerPluginSet:                    true,
				portainer.OperationDockerPluginDelete:                 true,
				portainer.OperationDockerSessionStart:                 true,
				portainer.OperationDockerDistributionInspect:          true,
				portainer.OperationDockerBuildPrune:                   true,
				portainer.OperationDockerBuildCancel:                  true,
				portainer.OperationDockerPing:                         true,
				portainer.OperationDockerInfo:                         true,
				portainer.OperationDockerVersion:                      true,
				portainer.OperationDockerEvents:                       true,
				portainer.OperationDockerSystem:                       true,
				portainer.OperationPortainerResourceControlCreate:     true,
				portainer.OperationPortainerResourceControlUpdate:     true,
				portainer.OperationPortainerResourceControlDelete:     true,
				portainer.OperationPortainerStackList:                 true,
				portainer.OperationPortainerStackInspect:              true,
				portainer.OperationPortainerStackFile:                 true,
				portainer.OperationPortainerStackCreate:               true,
				portainer.OperationPortainerStackMigrate:              true,
				portainer.OperationPortainerStackUpdate:               true,
				portainer.OperationPortainerStackDelete:               true,
				portainer.OperationPortainerWebsocketExec:             true,
				portainer.OperationPortainerWebhookList:               true,
				portainer.OperationPortainerWebhookCreate:             true,
				portainer.OperationPortainerWebhookDelete:             true,
				portainer.EndpointResourcesAccess:                     true,
			},
		}

		err = store.RoleService.CreateRole(environmentAdministratorRole)
		if err != nil {
			return err
		}

		environmentReadOnlyUserRole := &portainer.Role{
			Name:        "Helpdesk",
			Description: "Read-only authorizations on all the resources inside an endpoint",
			Authorizations: map[portainer.Authorization]bool{
				portainer.OperationDockerContainerArchiveInfo: true,
				portainer.OperationDockerContainerList:        true,
				portainer.OperationDockerContainerChanges:     true,
				portainer.OperationDockerContainerInspect:     true,
				portainer.OperationDockerContainerTop:         true,
				portainer.OperationDockerContainerLogs:        true,
				portainer.OperationDockerContainerStats:       true,
				portainer.OperationDockerImageList:            true,
				portainer.OperationDockerImageSearch:          true,
				portainer.OperationDockerImageGetAll:          true,
				portainer.OperationDockerImageGet:             true,
				portainer.OperationDockerImageHistory:         true,
				portainer.OperationDockerImageInspect:         true,
				portainer.OperationDockerNetworkList:          true,
				portainer.OperationDockerNetworkInspect:       true,
				portainer.OperationDockerVolumeList:           true,
				portainer.OperationDockerVolumeInspect:        true,
				portainer.OperationDockerSwarmInspect:         true,
				portainer.OperationDockerNodeList:             true,
				portainer.OperationDockerNodeInspect:          true,
				portainer.OperationDockerServiceList:          true,
				portainer.OperationDockerServiceInspect:       true,
				portainer.OperationDockerServiceLogs:          true,
				portainer.OperationDockerSecretList:           true,
				portainer.OperationDockerSecretInspect:        true,
				portainer.OperationDockerConfigList:           true,
				portainer.OperationDockerConfigInspect:        true,
				portainer.OperationDockerTaskList:             true,
				portainer.OperationDockerTaskInspect:          true,
				portainer.OperationDockerTaskLogs:             true,
				portainer.OperationDockerPluginList:           true,
				portainer.OperationDockerDistributionInspect:  true,
				portainer.OperationDockerPing:                 true,
				portainer.OperationDockerInfo:                 true,
				portainer.OperationDockerVersion:              true,
				portainer.OperationDockerEvents:               true,
				portainer.OperationDockerSystem:               true,
				portainer.OperationPortainerStackList:         true,
				portainer.OperationPortainerStackInspect:      true,
				portainer.OperationPortainerStackFile:         true,
				portainer.OperationPortainerWebhookList:       true,
				portainer.EndpointResourcesAccess:             true,
			},
		}

		err = store.RoleService.CreateRole(environmentReadOnlyUserRole)
		if err != nil {
			return err
		}

		standardUserRole := &portainer.Role{
			Name:        "Standard user",
			Description: "Regular user account restricted to access authorized resources",
			Authorizations: map[portainer.Authorization]bool{
				portainer.OperationDockerContainerArchiveInfo:         true,
				portainer.OperationDockerContainerList:                true,
				portainer.OperationDockerContainerExport:              true,
				portainer.OperationDockerContainerChanges:             true,
				portainer.OperationDockerContainerInspect:             true,
				portainer.OperationDockerContainerTop:                 true,
				portainer.OperationDockerContainerLogs:                true,
				portainer.OperationDockerContainerStats:               true,
				portainer.OperationDockerContainerAttachWebsocket:     true,
				portainer.OperationDockerContainerArchive:             true,
				portainer.OperationDockerContainerCreate:              true,
				portainer.OperationDockerContainerKill:                true,
				portainer.OperationDockerContainerPause:               true,
				portainer.OperationDockerContainerUnpause:             true,
				portainer.OperationDockerContainerRestart:             true,
				portainer.OperationDockerContainerStart:               true,
				portainer.OperationDockerContainerStop:                true,
				portainer.OperationDockerContainerWait:                true,
				portainer.OperationDockerContainerResize:              true,
				portainer.OperationDockerContainerAttach:              true,
				portainer.OperationDockerContainerExec:                true,
				portainer.OperationDockerContainerRename:              true,
				portainer.OperationDockerContainerUpdate:              true,
				portainer.OperationDockerContainerPutContainerArchive: true,
				portainer.OperationDockerContainerDelete:              true,
				portainer.OperationDockerImageList:                    true,
				portainer.OperationDockerImageSearch:                  true,
				portainer.OperationDockerImageGetAll:                  true,
				portainer.OperationDockerImageGet:                     true,
				portainer.OperationDockerImageHistory:                 true,
				portainer.OperationDockerImageInspect:                 true,
				portainer.OperationDockerImageLoad:                    true,
				portainer.OperationDockerImageCreate:                  true,
				portainer.OperationDockerImagePush:                    true,
				portainer.OperationDockerImageTag:                     true,
				portainer.OperationDockerImageDelete:                  true,
				portainer.OperationDockerImageCommit:                  true,
				portainer.OperationDockerImageBuild:                   true,
				portainer.OperationDockerNetworkList:                  true,
				portainer.OperationDockerNetworkInspect:               true,
				portainer.OperationDockerNetworkCreate:                true,
				portainer.OperationDockerNetworkConnect:               true,
				portainer.OperationDockerNetworkDisconnect:            true,
				portainer.OperationDockerNetworkDelete:                true,
				portainer.OperationDockerVolumeList:                   true,
				portainer.OperationDockerVolumeInspect:                true,
				portainer.OperationDockerVolumeCreate:                 true,
				portainer.OperationDockerVolumeDelete:                 true,
				portainer.OperationDockerExecInspect:                  true,
				portainer.OperationDockerExecStart:                    true,
				portainer.OperationDockerExecResize:                   true,
				portainer.OperationDockerSwarmInspect:                 true,
				portainer.OperationDockerSwarmUnlockKey:               true,
				portainer.OperationDockerSwarmInit:                    true,
				portainer.OperationDockerSwarmJoin:                    true,
				portainer.OperationDockerSwarmLeave:                   true,
				portainer.OperationDockerSwarmUpdate:                  true,
				portainer.OperationDockerSwarmUnlock:                  true,
				portainer.OperationDockerNodeList:                     true,
				portainer.OperationDockerNodeInspect:                  true,
				portainer.OperationDockerNodeUpdate:                   true,
				portainer.OperationDockerNodeDelete:                   true,
				portainer.OperationDockerServiceList:                  true,
				portainer.OperationDockerServiceInspect:               true,
				portainer.OperationDockerServiceLogs:                  true,
				portainer.OperationDockerServiceCreate:                true,
				portainer.OperationDockerServiceUpdate:                true,
				portainer.OperationDockerServiceDelete:                true,
				portainer.OperationDockerSecretList:                   true,
				portainer.OperationDockerSecretInspect:                true,
				portainer.OperationDockerSecretCreate:                 true,
				portainer.OperationDockerSecretUpdate:                 true,
				portainer.OperationDockerSecretDelete:                 true,
				portainer.OperationDockerConfigList:                   true,
				portainer.OperationDockerConfigInspect:                true,
				portainer.OperationDockerConfigCreate:                 true,
				portainer.OperationDockerConfigUpdate:                 true,
				portainer.OperationDockerConfigDelete:                 true,
				portainer.OperationDockerTaskList:                     true,
				portainer.OperationDockerTaskInspect:                  true,
				portainer.OperationDockerTaskLogs:                     true,
				portainer.OperationDockerPluginList:                   true,
				portainer.OperationDockerPluginPrivileges:             true,
				portainer.OperationDockerPluginInspect:                true,
				portainer.OperationDockerPluginPull:                   true,
				portainer.OperationDockerPluginCreate:                 true,
				portainer.OperationDockerPluginEnable:                 true,
				portainer.OperationDockerPluginDisable:                true,
				portainer.OperationDockerPluginPush:                   true,
				portainer.OperationDockerPluginUpgrade:                true,
				portainer.OperationDockerPluginSet:                    true,
				portainer.OperationDockerPluginDelete:                 true,
				portainer.OperationDockerSessionStart:                 true,
				portainer.OperationDockerDistributionInspect:          true,
				portainer.OperationDockerBuildPrune:                   true,
				portainer.OperationDockerBuildCancel:                  true,
				portainer.OperationDockerPing:                         true,
				portainer.OperationDockerInfo:                         true,
				portainer.OperationDockerVersion:                      true,
				portainer.OperationDockerEvents:                       true,
				portainer.OperationDockerSystem:                       true,
				portainer.OperationDockerUndefined:                    true,
				portainer.OperationPortainerResourceControlCreate:     true,
				portainer.OperationPortainerResourceControlUpdate:     true,
				portainer.OperationPortainerResourceControlDelete:     true,
				portainer.OperationPortainerStackList:                 true,
				portainer.OperationPortainerStackInspect:              true,
				portainer.OperationPortainerStackFile:                 true,
				portainer.OperationPortainerStackCreate:               true,
				portainer.OperationPortainerStackMigrate:              true,
				portainer.OperationPortainerStackUpdate:               true,
				portainer.OperationPortainerStackDelete:               true,
				portainer.OperationPortainerWebsocketExec:             true,
				portainer.OperationPortainerWebhookList:               true,
				portainer.OperationPortainerWebhookCreate:             true,
			},
		}

		err = store.RoleService.CreateRole(standardUserRole)
		if err != nil {
			return err
		}

		readOnlyUserRole := &portainer.Role{
			Name:        "Read-only user",
			Description: "Read-only user account restricted to access authorized resources",
			Authorizations: map[portainer.Authorization]bool{
				portainer.OperationDockerContainerArchiveInfo: true,
				portainer.OperationDockerContainerList:        true,
				portainer.OperationDockerContainerChanges:     true,
				portainer.OperationDockerContainerInspect:     true,
				portainer.OperationDockerContainerTop:         true,
				portainer.OperationDockerContainerLogs:        true,
				portainer.OperationDockerContainerStats:       true,
				portainer.OperationDockerImageList:            true,
				portainer.OperationDockerImageSearch:          true,
				portainer.OperationDockerImageGetAll:          true,
				portainer.OperationDockerImageGet:             true,
				portainer.OperationDockerImageHistory:         true,
				portainer.OperationDockerImageInspect:         true,
				portainer.OperationDockerNetworkList:          true,
				portainer.OperationDockerNetworkInspect:       true,
				portainer.OperationDockerVolumeList:           true,
				portainer.OperationDockerVolumeInspect:        true,
				portainer.OperationDockerSwarmInspect:         true,
				portainer.OperationDockerNodeList:             true,
				portainer.OperationDockerNodeInspect:          true,
				portainer.OperationDockerServiceList:          true,
				portainer.OperationDockerServiceInspect:       true,
				portainer.OperationDockerServiceLogs:          true,
				portainer.OperationDockerSecretList:           true,
				portainer.OperationDockerSecretInspect:        true,
				portainer.OperationDockerConfigList:           true,
				portainer.OperationDockerConfigInspect:        true,
				portainer.OperationDockerTaskList:             true,
				portainer.OperationDockerTaskInspect:          true,
				portainer.OperationDockerTaskLogs:             true,
				portainer.OperationDockerPluginList:           true,
				portainer.OperationDockerDistributionInspect:  true,
				portainer.OperationDockerPing:                 true,
				portainer.OperationDockerInfo:                 true,
				portainer.OperationDockerVersion:              true,
				portainer.OperationDockerEvents:               true,
				portainer.OperationDockerSystem:               true,
				portainer.OperationPortainerStackList:         true,
				portainer.OperationPortainerStackInspect:      true,
				portainer.OperationPortainerStackFile:         true,
				portainer.OperationPortainerWebhookList:       true,
			},
		}

		err = store.RoleService.CreateRole(readOnlyUserRole)
		if err != nil {
			return err
		}
	}

	return nil
}
