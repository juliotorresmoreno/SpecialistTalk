import React from 'react'
import Loading from '../components/Loading'
import { useGetSession } from '../services/auth'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const withSession = function <T = any>(
  WrappedComponent: React.ComponentType<any>
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const { isLoading } = useGetSession()

    if (isLoading) return <Loading />

    return <WrappedComponent {...props} />
  }

  return result
}

export default withSession
