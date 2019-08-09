// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { defineMessages } from 'react-intl'

import Message from '../../../lib/components/message'
import { getApplicationId, getDeviceId } from '../../../lib/selectors/id'
import Link from '../../../components/link'
import Map from '../../../components/map'

import style from './device-map.styl'

const m = defineMessages({
  location: 'Location',
  changeLocation: 'Change Location',
})

class DeviceMapWidget extends React.Component {
  render () {
    const {
      devIds,
      markers,
    } = this.props

    const devId = getDeviceId(devIds)
    const appId = getApplicationId(devIds)
    const leafletConfig = {
      zoomControl: false,
    }

    return (
      <aside className={style.wrapper}>
        <div className={style.header}>
          <Message
            className={style.titleMessage}
            content={m.location}
          />
          <Link to={`/applications/${appId}/devices/${devId}/location`}>
            <Message
              className={style.changeLocationMessage}
              content={m.changeLocation}
            />
            →
          </Link>
        </div>
        <Map id="device-map-widget" markers={markers} leafletConfig={leafletConfig} widget />
      </aside>
    )
  }
}

export default DeviceMapWidget
