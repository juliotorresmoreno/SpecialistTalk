import React, { useEffect, useState } from 'react'
import Loading from '../components/Loading'
import config from '../config'
import authSlice from '../features/auth'
import { ISession } from '../models/session'
import { IUser } from '../models/user'
import { useAppDispatch, useAppSelector } from '../store/hooks'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const url = config.baseUrl + '/auth/session'

const withSession = function <T = any>(
  WrappedComponent: React.ComponentType<any>
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const _session = useAppSelector((state) => state.auth.session)
    const [session, setSession] = useState<ISession | null>(null)
    const [isloading, setIsloading] = useState(false)
    const dispatch = useAppDispatch()

    useEffect(() => {
      if (!_session) return
      if (isloading) return
      setIsloading(true)
      const token = _session?.token ?? ''

      const opts: RequestInit = {
        headers: {
          'X-API-Key': token,
        },
      }

      fetch(url, opts)
        .then(async (response) => {
          if (response.ok) {
            const iuser: IUser = await response.json()
            const session: ISession = { token, user: iuser }
            setSession(session)
            dispatch(authSlice.actions.setSession(session))
            setIsloading(false)
            return
          }
          throw new Error('Unauthorized')
        })
        .catch(() => {
          dispatch(authSlice.actions.logout())
        })
    }, [_session])

    if (isloading) return <Loading />

    return <WrappedComponent {...props} />
  }
  return result
}

export default withSession
