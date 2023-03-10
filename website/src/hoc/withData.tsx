import React, { useEffect, useState } from 'react'
import Loading from '../components/Loading'
import ErrorPage from '../pages/ErrorPage'
import { useGetData } from '../services/api'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const withData = function <T = any>(
  WrappedComponent: React.ComponentType<any>,
  url: string
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const [data, setData] = useState<any>(null)
    const { get, error, isLoading } = useGetData(url)

    useEffect(() => {
      if (isLoading) return
      if (data) return
      get().then(setData)
    }, [isLoading])

    if (error) return <ErrorPage error={error} />
    if (!data || isLoading) return <Loading />
    return <WrappedComponent payload={data} {...props} />
  }
  return result
}

export default withData
