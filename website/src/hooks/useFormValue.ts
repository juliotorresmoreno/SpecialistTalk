import React, { useState } from 'react'

export default function useFormValue(
  defaultValue = ''
): [string, React.ChangeEventHandler<any>] {
  const [value, setValue] = useState(defaultValue)

  const handler: React.ChangeEventHandler<any> = (evt) => {
    setValue(evt.target.value)
  }

  return [value, handler]
}
