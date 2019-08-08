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
import bind from 'autobind-decorator'
import { defineMessages } from 'react-intl'

import Select from '../../../components/select'
import Input from '../../../components/input'

import PropTypes from '../../../lib/prop-types'
import computePrefix from './compute-prefix'

import style from './join-eui-prefixes-input.styl'

const m = defineMessages({
  empty: 'No prefix',
})

const getOptions = prefixes => prefixes
  .filter(prefix => Boolean(prefix.length))
  .map(function ({ join_eui, length }) {
    const computedPrefix = computePrefix(join_eui, length).toUpperCase()

    return { label: computedPrefix, value: computedPrefix }
  })

const emptyOption = { label: m.empty, value: '' }

@bind
class JoinEUIPrefixesInput extends React.PureComponent {

  constructor (props) {
    super(props)

    this.inputRef = React.createRef()
    this.state = {
      prefix: emptyOption.value,
    }
  }

  async handleChange (value) {
    const { onChange } = this.props
    const { prefix } = this.state

    if (Boolean(prefix) && !value.startsWith(prefix)) {
      await this.setState({ prefix: emptyOption.value })
    }

    onChange(value)
  }

  async handlePrefixChange (prefix) {
    const { onChange } = this.props

    await this.setState({ prefix })
    onChange(prefix)
    if (this.inputRef.current) {
      const instance = this.inputRef.current.getWrappedInstance()

      instance.focus()
    }
  }

  handleBlur (event) {
    const { name, onBlur } = this.props

    if (name === event.target.name) {
      onBlur(event)
    }
  }

  render () {
    const {
      className,
      name,
      disabled,
      value,
      error,
      prefixes,
      fetching,
      select,
    } = this.props
    const { prefix } = this.state

    let selectComponent = null
    if (select) {
      const selectOptions = getOptions(prefixes)
      selectOptions.unshift(emptyOption)

      selectComponent = (
        <Select
          className={style.select}
          name={`${name}-prefix`}
          disabled={disabled}
          options={selectOptions}
          onChange={this.handlePrefixChange}
          error={error}
          isLoading={fetching}
          value={prefix}
        />
      )
    }

    return (
      <div className={classnames(className, style.container)}>
        {selectComponent}
        <Input
          ref={this.inputRef}
          className={style.byte}
          value={value}
          defaultValue={prefix}
          name={name}
          disabled={disabled}
          min={8}
          max={8}
          type="byte"
          onChange={this.handleChange}
          error={error}
        />
      </div>
    )
  }
}

JoinEUIPrefixesInput.propTypes = {
  className: PropTypes.string,
  prefixes: PropTypes.arrayOf(PropTypes.shape({
    join_eui: PropTypes.string,
    length: PropTypes.number,
  })),
  name: PropTypes.string.isRequired,
  disabled: PropTypes.bool,
  onChange: PropTypes.func.isRequired,
  onBlur: PropTypes.func,
  value: PropTypes.string,
  fetching: PropTypes.bool,
  select: PropTypes.bool,
}

JoinEUIPrefixesInput.defaultProps = {
  disabled: false,
  onBlur: () => null,
  fetching: false,
  prefixes: [],
  select: true,
}

export default JoinEUIPrefixesInput
