import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router'
import Loading from '../components/Loading'
import authSlice from '../features/auth'
import ErrorPage from '../pages/ErrorPage'
import { useGetData } from '../services/api'
import { useAppDispatch, useAppSelector } from '../store/hooks'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

type args<T> = {
  WrappedComponent: React.ComponentType<any>
  FallbackComponent?: React.ComponentType<any>
  url: string
  skipper?: (id: string) => T | null
  callback: (data: T) => void
  withAuth: boolean
}

const withDataById = function <T = any, S = any>({
  WrappedComponent,
  FallbackComponent,
  url,
  callback,
  skipper,
  withAuth,
}: args<S>) {
  const result: React.FC<T & ResultProps> = (props) => {
    const [dataId, setDataId] = useState<string | null>(null)
    const [data, setData] = useState<S | null>(null)
    const id = useParams<{ id: string }>().id as string
    const _url = url + '/' + id
    const { get, error, isLoading } = useGetData(_url)
    const omit = skipper ? skipper(id) : null
    const session = useAppSelector((state) => state.auth.session)
    const dispatch = useAppDispatch()

    useEffect(() => {
      let _id = id
      if (withAuth && !session) return
      if (isLoading) return
      if (dataId === _id) return
      if (omit) {
        setData(omit)
        return
      }

      get()
        .then((data: S) => {
          setData(data)
          setDataId(_id)

          callback(data)
        })
        .catch((err) => {
          if (err.message === 'unauthorized')
            dispatch(authSlice.actions.logout())
        })
    }, [isLoading, session, id])

    const isFallback = !data || !session
    if (error) return <ErrorPage error={error} />
    if (isLoading) return <Loading />
    if (isFallback && FallbackComponent) return <FallbackComponent />
    if (isFallback && !FallbackComponent) return <Loading />

    return <WrappedComponent data={data} {...props} />
  }
  return result
}

export default withDataById
