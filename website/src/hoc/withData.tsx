import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router'
import Loading from '../components/Loading'
import config from '../config'
import { IChat } from '../models/chat'
import ErrorPage from '../pages/ErrorPage'
import { useAppSelector } from '../store/hooks'
import { HTTPError } from '../types/http'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const withData = function <T = any>(
  WrappedComponent: React.ComponentType<any>,
  url: string
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const session = useAppSelector((state) => state.auth.session)
    const [isloading, setIsloading] = useState<boolean>(false)
    const [data, setData] = useState<IChat | null>(null)
    const [error, setError] = useState<HTTPError | null>(null)

    useEffect(() => {
      setIsloading(true)
      setError(null)

      fetch(url, {
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
          setData(content)
        })
        .catch((err: Error) => setError({ message: err.message }))
        .finally(() => setIsloading(false))
    }, [session])

    if (error) return <ErrorPage error={error} />
    if (!data || isloading) return <Loading />
    return <WrappedComponent payload={data} {...props} />
  }
  return result
}

export default withData
