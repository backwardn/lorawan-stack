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

import React, { useEffect } from 'react'
import { defineMessages } from 'react-intl'

import PropTypes from '../../../lib/prop-types'

import Field from '../../../components/form/field'
import connect from './connect'
import JoinEUIPrefixesInput from './join-eui-prefixes-input'

const m = defineMessages({
  prefixesFetchingFailure: 'Could not retrieve the list of available prefixes',
})

const JoinEUIPrefixesField = function ({
  name,
  title,
  required,
  horizontal,
  autoFocus,
  disabled,
  prefixes,
  error,
  fetching,
  getPrefixes,
  select,
}) {

  useEffect(function () {
    getPrefixes()
  }, [])

  return (
    <Field
      name={name}
      title={title}
      required={required}
      autoFocus={autoFocus}
      horizontal={horizontal}
      component={JoinEUIPrefixesInput}
      disabled={disabled}
      fetching={fetching}
      warning={Boolean(error) ? m.prefixesFetchingFailure : undefined}
      prefixes={prefixes}
      select={select}
    />
  )
}

JoinEUIPrefixesField.propTypes = {
  name: PropTypes.string.isRequired,
  title: PropTypes.message.isRequired,
  required: PropTypes.bool,
  horizontal: PropTypes.bool,
  autoFocus: PropTypes.bool,
  disabled: PropTypes.bool,
  prefixes: PropTypes.arrayOf(PropTypes.shape({
    prefix: PropTypes.string,
    length: PropTypes.number,
  })),
  fetching: PropTypes.bool.isRequired,
  error: PropTypes.error,
  getPrefixes: PropTypes.func.isRequired,
  select: PropTypes.bool,
}

JoinEUIPrefixesField.defaultProps = {
  required: false,
  horizontal: false,
  autoFocus: false,
  disabled: false,
  prefixes: [],
  select: true,
}

export default connect(JoinEUIPrefixesField)
