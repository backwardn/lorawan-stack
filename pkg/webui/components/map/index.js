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
import L from 'leaflet'

import style from './map.styl'

// Fix broken links to leaflet images.
delete L.Icon.Default.prototype._getIconUrl
L.Icon.Default.mergeOptions({
  iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
  iconUrl: require('leaflet/dist/images/marker-icon.png'),
  shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
})
class Map extends React.Component {

  getMapCenter (markers) {
    return markers[0]
  }

  componentDidMount () {
    const { markers } = this.props
    const { position } = ( markers.length >= 1 ) ? this.getMapCenter(markers) : markers[0]

    // create map
    this.map = L.map('map', {
      center: [ position.latitude, position.longitude ],
      zoom: 11,
      layers: [
        L.tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
          attribution: '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
        }),
      ],
      zoomControl: false,
    })

    markers.map(maker =>
      L.marker([ maker.position.latitude, maker.position.longitude ])
        .bindPopup( maker.name )
        .addTo(this.map)
    )
  }

  render () {
    return (
      <div className={style.mapContainer}>
        <div className={style.map} id="map" />
      </div>

    )
  }
}

export default Map

