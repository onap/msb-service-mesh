/**
 * Copyright (c) 2018 ZTE Corporation.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * and the Apache License 2.0 which both accompany this distribution,
 * and are available at http://www.eclipse.org/legal/epl-v10.html
 * and http://www.apache.org/licenses/LICENSE-2.0
 *
 * Contributors:
 *     ZTE - initial Project
 */
package pilot

var (
	cachedServices []*models.MsbService
)

func SyncMsbData(newServices []*models.MsbService) {
	log.Log.Debug("sync msb rewrite rule to pilot")
	createServices, updateServices, deleteServices := compareServices(cachedServices, newServices)

	log.Log.Debug("SyncMsbData: ", len(createServices), len(updateServices), len(deleteServices))

	cachedServices = newServices
}

func compareServices(oldServices, newServices []*models.MsbService) (createServices, updateServices, deleteServices []*models.MsbService) {
	oldServiceMap := toServiceMap(oldServices)
	newServiceMap := toServiceMap(newServices)

	for key, newService := range newServiceMap {
		if oldService, exist := oldServiceMap[key]; exist {
			// service exist: check whether need to update
			if oldService.ModifyIndex != newService.ModifyIndex {
				updateServices = append(updateServices, newService)
			}
		} else {
			// service not exist: add
			createServices = append(createServices, newService)
		}

		delete(oldServiceMap, key)
	}

	for _, service := range oldServiceMap {
		deleteServices = append(deleteServices, service)
	}

	return
}

func toServiceMap(services []*models.MsbService) map[string]*models.MsbService {
	serviceMap := make(map[string]*models.MsbService)

	for _, service := range services {
		serviceMap[service.ServiceName] = service
	}

	return serviceMap
}
