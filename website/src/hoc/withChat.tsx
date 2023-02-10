import React from 'react'
import { useParams } from 'react-router'
import Loading from '../components/Loading'
import ErrorPage from '../pages/ErrorPage'
import { useGetChat } from '../services/chats'

type ResultProps = {
  [x: string | number | symbol]: any
} & React.PropsWithChildren

const withChat = function <T = any>(
  WrappedComponent: React.ComponentType<any>
) {
  const result: React.FC<T & ResultProps> = (props) => {
    const id = useParams<{ id: string }>().id as string
    const { data, error, isLoading } = useGetChat(id)

    if (error) return <ErrorPage error={error} />
    if (!data || isLoading) return <Loading />

    return <WrappedComponent chat={data} {...props} />
  }
  return result
}

export default withChat
