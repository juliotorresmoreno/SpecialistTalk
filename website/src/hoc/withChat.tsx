import React, { useEffect, useState } from 'react'
import { useLocation, useParams } from 'react-router'
import config from '../config'
import { User } from '../models/user'
import ErrorPage from '../pages/ErrorPage'
import { useAppSelector } from '../store/hooks'
import { HTTPError } from '../types/http'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const url = config.baseUrl + '/chats/'

const withChat = function <T = any>(
  WrappedComponent: React.ComponentType<any>
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const session = useAppSelector((state) => state.auth.session)
    const [isload, setIsload] = useState(false)
    const [chat, setChat] = useState<User | null>(null)
    const [error, setError] = useState<HTTPError | null>(null)
    const { id } = useParams()

    useEffect(() => {
      if (isload) return
      setIsload(true)

      fetch(url + id, {
        headers: {
          'X-API-Key': session?.token ?? '',
        },
      })
        .then(async (response) => {
          const content = await response.json()
          if (!response.ok) {
            setError(content)
            return
          }
          setChat(content)
        })
        .catch((err: Error) => {
          setError({ message: err.message })
        })
    }, [session, isload])

    if (error) return <ErrorPage error={error} />

    if (!chat) return <>Loading</>

    return <WrappedComponent {...props} />
  }
  return result
}

export default withChat
