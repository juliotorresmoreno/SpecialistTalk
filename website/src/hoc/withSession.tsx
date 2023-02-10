import React, { useEffect, useState } from 'react'
import Loading from '../components/Loading'
import { useGetSession } from '../services/auth'
import { useAppSelector } from '../store/hooks'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const withSession = function <T = any>(
  WrappedComponent: React.ComponentType<any>
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const session = useAppSelector((state) => state.auth.session)
    const [data, setData] = useState<any>(null)
    const { isLoading, get } = useGetSession()

    useEffect(() => {
      if (!session) return
      if (isLoading) return
      if (data) return
      get().then(setData)
    }, [isLoading])

    if (isLoading) return <Loading />

    return <WrappedComponent {...props} />
  }

  return result
}

export default withSession
