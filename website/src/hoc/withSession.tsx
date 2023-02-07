import React, { useEffect, useState } from 'react'
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
    const [session, setSession] = useState(_session)
    const [isload, setIsload] = useState(false)
    const dispatch = useAppDispatch()

    useEffect(() => {
      if (!_session) return
      if (isload) return
      setIsload(true)
      const token = session?.token ?? ''

      fetch(url, {
        headers: {
          'X-API-Key': token,
        },
      }).then(async (response) => {
        if (response.ok) {
          const iuser: IUser = await response.json()
          const session: ISession = { token, user: iuser }
          setSession(session)
          dispatch(authSlice.actions.setSession(session))
        }
      })
    }, [session])

    return <WrappedComponent {...props} />
  }
  return result
}

export default withSession
