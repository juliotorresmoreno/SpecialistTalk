import React, { useState } from 'react'

export default function useToggle(defaultValue = false): [boolean, () => void] {
  const [value, setValue] = useState(defaultValue)

  const toggle = () => {
    setValue(!value)
  }

  return [value, toggle]
}
