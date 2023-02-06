import React, { useState } from 'react'
import { useAppSelector } from '../store/hooks'
import config from '../config'
import { IUser, User } from '../models/user'

type Result = [string, React.ChangeEventHandler<any>, User[]]

export default function useSearch(url = ''): Result {
  const session = useAppSelector((state) => state.auth.session)
  const [value, setValue] = useState('')
  const [result, setResult] = useState<User[]>([])

  const handler: React.ChangeEventHandler<any> = (evt) => {
    //
    const value: string = evt.target.value
    setValue(value)
    const token = session?.token ?? ''

    if (value.length >= 3) {
      const _url = new URL(config.baseUrl + url)
      _url.searchParams.set('q', value)

      fetch(_url.toString(), {
        headers: {
          'X-API-Key': token,
        },
      }).then(async (response) => {
        if (response.ok) {
          const content: IUser[] = await response.json()
          const result = content.map((el) => new User(el))
          setResult(result)
        }
      })
      return
    }
    setResult([])
  }

  return [value, handler, result]
}
