import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router'
import Loading from '../components/Loading'
import ErrorPage from '../pages/ErrorPage'
import { useGetData } from '../services/api'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const withDataById = function <T = any>(
  WrappedComponent: React.ComponentType<any>,
  url: string
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const [dataId, setDataId] = useState<any>(null)
    const [data, setData] = useState<any>(null)
    const id = useParams<{ id: string }>().id as string
    const _url = url + '/' + id
    const { get, error, isLoading } = useGetData(_url)

    useEffect(() => {
      let _id = id
      if (isLoading) return
      if (dataId === _id) return
      get().then((data) => {
        setData(data)
        setDataId(_id)
      })
    }, [isLoading, id])

    if (error) return <ErrorPage error={error} />
    if (!data || isLoading) return <Loading />

    return <WrappedComponent chat={data} {...props} />
  }
  return result
}

export default withDataById
