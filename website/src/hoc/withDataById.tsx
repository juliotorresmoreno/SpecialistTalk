import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router'
import Loading from '../components/Loading'
import ErrorPage from '../pages/ErrorPage'
import { useGetData } from '../services/api'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

type args<T> = {
  WrappedComponent: React.ComponentType<any>
  url: string
  skipper?: (id: string) => T | null
  callback: (data: T) => void
}

const withDataById = function <T = any, S = any>({
  WrappedComponent,
  url,
  callback,
  skipper,
}: args<S>) {
  const result: React.FC<T & ResultProps> = (props) => {
    const [dataId, setDataId] = useState<string | null>(null)
    const [data, setData] = useState<S | null>(null)
    const id = useParams<{ id: string }>().id as string
    const _url = url + '/' + id
    const { get, error, isLoading } = useGetData(_url)
    const omit = skipper ? skipper(id) : null

    useEffect(() => {
      let _id = id
      if (isLoading) return
      if (dataId === _id) return
      if (omit) {
        setData(omit)
        return
      }

      get().then((data: S) => {
        setData(data)
        setDataId(_id)

        callback(data)
      })
    }, [isLoading, id])

    if (error) return <ErrorPage error={error} />
    if (!data || isLoading) return <Loading />

    return <WrappedComponent data={data} {...props} />
  }
  return result
}

export default withDataById
