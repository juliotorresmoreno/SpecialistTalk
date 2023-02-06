import React, { useState } from 'react'

type Result = [string, React.ChangeEventHandler<any>]

export default function useFormValue(defaultValue = ''): Result {
  const [value, setValue] = useState(defaultValue)

  const handler: React.ChangeEventHandler<any> = (evt) => {
    setValue(evt.target.value)
  }

  return [value, handler]
}
