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
import classnames from 'classnames'
import Leaflet from 'leaflet'

import style from './map.styl'

// Fix broken links to leaflet images.
delete Leaflet.Icon.Default.prototype._getIconUrl
Leaflet.Icon.Default.mergeOptions({
  iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
  iconUrl: require('leaflet/dist/images/marker-icon.png'),
  shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
})
class Map extends React.Component {

  getMapCenter (markers) {
    // This will calculate zoom and map center long/lang based on all markers provided.
    // Currenlty its just returns the first marker.
    return markers[0]
  }

  createMap (config, id) {
    this.map = Leaflet.map( id, {
      ...config,
    })
  }

  createMarkers (markers) {
    markers.map(marker =>
      Leaflet.marker([ marker.position.latitude, marker.position.longitude ])
        .addTo(this.map)
    )
  }

  componentDidMount () {
    const {
      id,
      markers,
    } = this.props

    const { position } = ( markers.length >= 1 ) ? this.getMapCenter(markers) : markers[0]

    const config = {
      layers: [
        Leaflet.tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
          attribution: '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
        }),
      ],
      center: [ position.latitude, position.longitude ],
      zoom: 11,
      minZoom: 1,
      ...this.props.leafletConfig,
    }

    this.createMap(config, id)
    this.createMarkers(markers, id)

  }

  render () {
    const {
      id,
      widget,
    } = this.props

    return (
      <div className={style.container}>
        <div className={classnames(
          style.map,
          { [style.widget]: widget },
        )} id={id}
        />
      </div>
    )
  }
}

export default Map
