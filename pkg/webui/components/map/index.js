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

class Map extends React.Component {

  constructor(props){
    super(props);
  }

  componentDidMount () {
    const { mapData } = this.props 
    // create map
    this.map = L.map('map', {
      center: [ mapData.position.latitude, mapData.position.longitude ],
      zoom: 11,
      layers: [
        L.tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
          attribution: '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
        }),
      ],
      zoomControl:false,
    })

    mapData.markers.map(maker => {
      L.marker([maker.position.latitude, maker.position.longitude ])
        .bindPopup( maker.name )
        .addTo(this.map)
    })
  }

  render () {
    return (
        <div className={style.map} id="map" />
    )
  }
}

export default Map

