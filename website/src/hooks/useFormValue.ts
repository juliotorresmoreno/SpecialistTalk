import React, { useState } from 'react'

export default function useFormValue<T = string | number>(defaultValue: T) {
  const [value, setValue] = useState(defaultValue)

  const handler: React.ChangeEventHandler<any> = (evt) => {
    setValue(evt.target.value)
  }
  type ResultType = [typeof value, typeof handler, typeof setValue]
  const result: ResultType = [value, handler, setValue]

  return result
}
