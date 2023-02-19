import React, { useEffect, useState } from 'react'
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
  skipper?: () => T | null
  callback?: (data: T) => void
  withAuth: boolean
}

const withData = function <T = any, S = any>({
  WrappedComponent,
  FallbackComponent,
  url,
  callback,
  skipper,
  withAuth,
}: args<S>) {
  const result: React.FC<T & ResultProps> = (props) => {
    const session = useAppSelector((state) => state.auth.session)
    const [data, setData] = useState<S | null>(null)
    const { get, error, isLoading } = useGetData(url)
    const dispatch = useAppDispatch()
    const omit = skipper ? skipper() : null

    useEffect(() => {
      if (omit) {
        setData(omit)
        return
      }
      if (withAuth && !session) return
      if (isLoading) return
      if (data) return
      get()
        .then((data: S) => {
          setData(data)

          callback && callback(data)
        })
        .catch((err) => {
          if (err.message === 'unauthorized')
            dispatch(authSlice.actions.logout())
        })
    }, [isLoading, session])

    const isFallback = !data || !session
    if (error) return <ErrorPage error={error} />
    if (isLoading) return <Loading />
    if (isFallback && FallbackComponent) return <FallbackComponent />
    if (isFallback && !FallbackComponent) return <Loading />
    return <WrappedComponent payload={data} {...props} />
  }
  return result
}

export default withData
